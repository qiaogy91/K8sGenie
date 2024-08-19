package impl

import (
	"context"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/alert"
	"gitee.com/qiaogy91/K8sGenie/apps/alert/impl/provider"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/common"
)

// HandlerAlert 拿到告警后
// 1. 转换为复合飞书 card 的格式
// 2. 调用router 进行路由查找
// 3. 调用client 进行发送
func (i *Impl) HandlerAlert(ctx context.Context, req *alert.HandlerAlertReq) (*alert.HandlerAlertRsp, error) {
	for _, alter := range req.Alerts {
		// 转化为飞书卡片 json
		bs, err := i.handlerAlert(ctx, alter)
		if err != nil {
			common.L().Error().Msgf("handlerAlert err: %v", err)
			continue
		}

		// 发送卡片告警
		rsp := i.c.SendCard(ctx, "", bs)
		common.L().Info().Msgf("send alert succeed, %s", rsp)
	}
	return nil, nil
}

func (i *Impl) handlerAlert(ctx context.Context, alert *alert.Alert) ([]byte, error) {
	var (
		color   string
		header  string
		content string
		footer  string
	)

	// 告警color、header、footer 处理
	switch alert.Status {
	case "firing":
		color = "red"
		header = "【容器集群】告警通知"
		footer = fmt.Sprintf(`**开始时间：**%s\n`, alert.StartsAt)
	case "resolved":
		color = "green"
		header = "【容器集群】告警恢复"
		footer = fmt.Sprintf(`**开始时间：**%s\n**结束时间：**%s\n`, alert.StartsAt, alert.EndsAt)
	}

	// 告警内容处理
	switch ns, ok := alert.Labels["namespace"]; {
	case ok:
		rsp, err := i.k.DescNamespace(ctx, &k8s.DescNamespaceReq{ClusterName: alert.Labels["robot_name"], NamespaceName: ns})
		if err != nil {
			common.L().Error().Msgf("ns/%s can not trigger alert, because project not found", ns)
			return nil, err
		}

		// 存在namespace
		content = fmt.Sprintf(
			`"**告警名称：**%s\n**告警级别：**%s\n**告警集群：**%s\n**所属产线：**%s\n**项目名称：**%s\n**项目编码：**%s\n**名称空间：**%s\n**故障描述：**%s\n**故障详情：**%s\n`,
			alert.Labels["alertname"],
			alert.Labels["level"],
			alert.Labels["robot_name"],

			rsp.Project.Spec.ProjectLine,
			rsp.Project.Spec.ProjectDesc,
			rsp.Project.Spec.ProjectCode,
			ns,
			alert.Annotations["summary"],
			alert.Annotations["description"],
		)
	default:
		// 不存在namespace
		content = fmt.Sprintf(
			`"**告警名称：**%s\n**告警级别：**%s\n**故障描述：**%s\n**故障详情：**%s\n`,
			alert.Labels["alertname"],
			alert.Labels["level"],
			alert.Annotations["summary"],
			alert.Annotations["description"],
		)
	}

	// 最终模板
	card := fmt.Sprintf(provider.CardTemplate, color, header, content, footer)
	return []byte(card), nil
}
