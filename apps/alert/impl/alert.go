package impl

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/alert"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/router"
	"gitee.com/qiaogy91/K8sGenie/common"
	"io"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
	"time"
)

// HandlerAlert 拿到告警后
// 1. 转换为复合飞书 card 的格式
// 2. 调用router 进行路由查找
// 3. 调用client 进行发送

// GenSign 辅助函数，飞书机器人消息签名校验，使用timestamp + key 做sha256, 再进行base64 encode
func (i *Impl) getSign(key string, timestamp int64) string {
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + key

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	h.Write(data)

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}

// sendCard 辅助函数，通过http 客户端发送请求
func (i *Impl) sendCard(ctx context.Context, url string, data []byte) []byte {
	// req
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	// send
	rsp, err := http.DefaultClient.Do(req)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		fmt.Printf("http default client send req failed: %v", err)
	}

	// rsp
	content, _ := io.ReadAll(rsp.Body)
	return content
}

// HandlerAlert 告警处理入口，AlertManager 会请求这个接口
func (i *Impl) HandlerAlert(ctx context.Context, req *alert.HandlerAlertReq) (*alert.HandlerAlertRsp, error) {
	rsp := &alert.HandlerAlertRsp{}
	for _, alter := range req.Alerts {
		res, err := i.handler(ctx, alter)
		if err != nil {
			common.L().Error().Msgf("handlerAlert err: %v", err)
			continue
		}

		response := &alert.Response{}
		if err := json.Unmarshal(res, response); err != nil {
			common.L().Error().Msgf("feishu response Unmarshal err: %v", err)
			continue
		}
		rsp.Rsp = append(rsp.Rsp, response)
	}
	// 返回给 AlertManager 的响应
	return rsp, nil
}

// handler 核心
// 1. 从告警中提取感兴趣数据
// 2. 渲染成飞书卡片的json
// 3. 发送给机器人 webhook
func (i *Impl) handler(ctx context.Context, alerter *alert.Alert) ([]byte, error) {

	var (
		webhook   string
		timestamp = time.Now().Unix()
		sign      string
		color     string
		header    string
		content   string
		footer    string
	)
	// **************************** 告警标题处理 ****************************
	switch alerter.Status {
	case "firing":
		color = "red"
		header = "【容器集群】告警通知"
		footer = fmt.Sprintf(`**开始时间：**%s\n`, alerter.StartsAt)
	case "resolved":
		color = "green"
		header = "【容器集群】告警恢复"
		footer = fmt.Sprintf(`**开始时间：**%s\n**结束时间：**%s\n`, alerter.StartsAt, alerter.EndsAt)
	}

	// **************************** 告警内容处理 ****************************
	switch ns, ok := alerter.Labels["namespace"]; {
	// 存在namespace，则根据所在的ProjectID 进行路由
	case ok:
		rsp, err := i.k.DescNamespace(ctx, &k8s.DescNamespaceReq{ClusterName: alerter.Labels["robot_name"], NamespaceName: ns})
		if err != nil {
			common.L().Error().Msgf("ns/%s can not trigger alerter, because project not found", ns)
			return nil, err
		}
		content = fmt.Sprintf(
			`**告警名称：**%s\n**告警级别：**%s\n**告警集群：**%s\n**所属产线：**%s\n**项目名称：**%s\n**项目编码：**%s\n**名称空间：**%s\n**故障描述：**%s\n**故障详情：**%s\n`,
			alerter.Labels["alertname"],
			alerter.Labels["level"],
			alerter.Labels["robot_name"],

			rsp.Project.Spec.ProjectLine,
			rsp.Project.Spec.ProjectDesc,
			rsp.Project.Spec.ProjectCode,
			ns,
			alerter.Annotations["summary"],
			alerter.Annotations["description"],
		)
		// 获取路由信息
		rs, err := i.r.QueryRoute(ctx, &router.QueryRouteReq{SearchType: router.SEARCH_TYPE_SEARCH_TYPE_PROJECT_ID, KeyWord: rsp.Project.Spec.ProjectId})
		if err != nil || len(rs.Items) < 1 {
			common.L().Error().Msgf("ns/%s cquery router failed, %+v", ns, err)
		}
		sign = i.getSign(rs.Items[0].Spec.WebhookToken, timestamp)
		webhook = rs.Items[0].Spec.WebhookUrl

	// 不存在namespace，则根据robotName 进行路由
	default:
		content = fmt.Sprintf(
			`**告警名称：**%s\n**告警级别：**%s\n**故障描述：**%s\n**故障详情：**%s\n`,
			alerter.Labels["alertname"],
			alerter.Labels["level"],
			alerter.Annotations["summary"],
			alerter.Annotations["description"],
		)
		// 获取路由信息
		rs, err := i.r.QueryRoute(ctx, &router.QueryRouteReq{SearchType: router.SEARCH_TYPE_SEARCH_TYPE_CLUSTER_NAME, KeyWord: alerter.Labels["robot_name"]})
		if err != nil || len(rs.Items) < 1 {
			common.L().Error().Msgf("cluster /%s cquery router failed, %+v", alerter.Labels["robot_name"], err)
		}
		sign = i.getSign(rs.Items[0].Spec.WebhookToken, timestamp)
		webhook = rs.Items[0].Spec.WebhookUrl
	}

	// **************************** 卡片模板渲染 ****************************
	card := fmt.Sprintf(alert.CardTemplate, timestamp, sign, content, footer, color, header)
	return i.sendCard(ctx, webhook, []byte(card)), nil
}
