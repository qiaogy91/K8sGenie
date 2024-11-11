package impl

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/alert"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	larkcontact "github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/rs/xid"
	"time"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&alert.UrgentAlert{})
}

func (i *Impl) UrgentAlertCall(ctx context.Context, req *alert.UrgentAlertCallRequest) (*alert.UrgentAlert, error) {
	// 查找用户的id
	uid, err := i.GetUserId(ctx, req.Users.Phone)
	if err != nil {
		return nil, err
	}
	// 消息模板渲染
	content := i.RenderUrgentCard(ctx, req.Alert)
	// 单独发送消息给用户
	mid, err := i.SendUrgentMessage(ctx, uid, content)
	if err != nil {
		return nil, err
	}

	// 对已发送的消息加急
	ins, err := i.SendUrgentCall(ctx, uid, mid)
	if err != nil {
		return nil, err
	}

	// 数据库录入
	ins.Users = req.Users
	ins.Content = content
	if err := i.db.WithContext(ctx).Model(&alert.UrgentAlert{}).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) RenderUrgentCard(ctx context.Context, a *alert.Alert) string {
	var (
		color   string
		header  string
		content string
		footer  string
	)
	// **************************** (1)告警标题处理 ****************************
	switch a.Status {
	case "firing":
		color = "red"
		header = "【容器集群】告警通知"
		footer = fmt.Sprintf(`**开始时间：**%s\n`, a.GetStartTime())
	case "resolved":
		color = "green"
		header = "【容器集群】告警恢复"
		footer = fmt.Sprintf(`**开始时间：**%s\n**结束时间：**%s\n`, a.GetStartTime(), a.GetEndTime())
	}

	// **************************** (2)告警内容处理 ****************************
	var (
		alertName        = a.Labels["alertname"]
		alertLevel       = a.Labels["level"]
		alertSummary     = a.Annotations["record"]
		alertDescription = a.Annotations["description"]
		alertCluster     = a.Labels["cluster_name"]
	)

	switch ns, ok := a.Labels["namespace"]; {
	case ok:
		rsp, err := i.k.DescNamespace(ctx, &k8s.DescNamespaceReq{ClusterName: a.Labels["cluster_name"], NamespaceName: ns})

		// 名称空间告警渲染：能够获取到项目信息的、以及获取不到项目信息的
		if err != nil {
			content = fmt.Sprintf(alert.NamespaceWithoutProject, alertName, alertLevel, alertCluster, ns, alertSummary, alertDescription)
		} else {
			content = fmt.Sprintf(alert.NamespaceContent, alertName, alertLevel, alertCluster, rsp.Spec.ProjectLine, rsp.Spec.ProjectDesc, rsp.Spec.ProjectCode, ns, alertSummary, alertDescription)
		}

	default:
		content = fmt.Sprintf(alert.ClusterContent, alertName, alertLevel, alertSummary, alertDescription)
	}

	// **************************** (3)卡片模板渲染 ****************************
	return fmt.Sprintf(alert.UrgentCardTemplate, content, footer, color, header)
}

func (i *Impl) GetUserId(ctx context.Context, phone string) (string, error) {
	req := larkcontact.NewBatchGetIdUserReqBuilder().
		UserIdType("open_id").
		Body(larkcontact.NewBatchGetIdUserReqBodyBuilder().
			Mobiles([]string{phone}).
			IncludeResigned(false).
			Build()).
		Build()
	rsp, err := i.c.Contact.User.BatchGetId(ctx, req)
	// 处理错误
	if err != nil {
		return "", err
	}
	if !rsp.Success() {
		return "", errors.New(rsp.Msg)
	}

	return *rsp.Data.UserList[0].UserId, nil
}

func (i *Impl) SendUrgentMessage(ctx context.Context, uid, content string) (string, error) {
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType("open_id").
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(uid).
			MsgType("interactive").
			Content(content).
			Uuid(xid.New().String()).
			Build()).
		Build()
	// 发起请求
	rsp, err := i.c.Im.V1.Message.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		return "", err
	}

	// 服务端错误处理
	if !rsp.Success() {
		return "", errors.New(rsp.Msg)
	}
	// 业务处理
	return *rsp.Data.MessageId, nil
}

func (i *Impl) SendUrgentCall(ctx context.Context, uid, mid string) (*alert.UrgentAlert, error) {
	req := larkim.NewUrgentPhoneMessageReqBuilder().
		MessageId(mid).
		UserIdType("open_id").
		UrgentReceivers(larkim.NewUrgentReceiversBuilder().
			UserIdList([]string{uid}).
			Build()).
		Build()
	// 发起请求
	rsp, err := i.c.Im.V1.Message.UrgentPhone(context.Background(), req)

	// 处理错误
	if err != nil {
		return nil, err
	}

	// 服务端错误处理
	if !rsp.Success() {
		return nil, errors.New(rsp.Msg)
	}

	// 业务处理
	ins := &alert.UrgentAlert{
		CreatedAt: time.Now().Unix(),
		Code:      int64(rsp.Code),
		Msg:       rsp.Msg,
	}
	return ins, nil
}
