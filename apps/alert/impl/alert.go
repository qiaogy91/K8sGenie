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
	"strings"
	"time"
)

// GenSign 飞书提供的签名函数
func (i *Impl) getSign(key string, timestamp int64) string {
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + key

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	h.Write(data)

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}

// sendCard 发送HTTP请求
func (i *Impl) sendCard(ctx context.Context, url string, data []byte) (*alert.Response, error) {
	// req
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	// send
	rsp, err := http.DefaultClient.Do(req)
	defer func() {
		if err := rsp.Body.Close(); err != nil {
			common.L().Error().Msgf("rsp body close err: %v", err)
		}
	}()
	if err != nil {
		fmt.Printf("http default client send req failed: %v", err)
	}

	// rsp
	content, _ := io.ReadAll(rsp.Body)
	ins := &alert.Response{}
	if err := json.Unmarshal(content, ins); err != nil {
		return nil, err
	}
	return ins, nil
}

// HandlerAlert 告警处理入口，AlertManager 会请求这个接口
func (i *Impl) HandlerAlert(ctx context.Context, req *alert.HandlerAlertReq) (*alert.HandlerAlertRsp, error) {
	rsp := &alert.HandlerAlertRsp{}
	for _, alerts := range req.Alerts {
		// ******************** 获取路由信息 ****************************
		route, err := i.r.AlertRoute(ctx, &router.AlertRouteReq{RobotName: alerts.Labels["robot_name"], NamespaceName: alerts.Labels["namespace"]})
		if err != nil {
			return nil, err
		}

		// ******************** 模板渲染 ****************************
		data, err := i.renderTemplate(ctx, alerts, route)
		if err != nil {
			common.L().Error().Msgf("renderTemplate failed: %v", err)
			continue
		}

		// ******************** 发送数据 ****************************
		res, err := i.sendCard(ctx, route.Spec.WebhookUrl, data)
		if err != nil {
			common.L().Error().Msgf("sendCard failed: %v", err)
			continue
		}
		rsp.Rsp = append(rsp.Rsp, res)
		// ******************** 自愈操作链 ****************************
		actions, ok := alerts.Labels["actions"]
		if !ok {
			continue
		}
		for _, action := range strings.Split(actions, ",") {
			if err := i.RecoveryAction(ctx, alerts, action); err != nil {
				common.L().Error().Msgf("RecoveryAction failed: %v", err)
				continue
			}
		}
	}
	return rsp, nil
}

// renderTemplate 模板渲染
func (i *Impl) renderTemplate(ctx context.Context, alerter *alert.Alert, route *router.Router) ([]byte, error) {
	var (
		timestamp = time.Now().Unix()
		sign      = i.getSign(route.Spec.WebhookToken, timestamp)
		color     string
		header    string
		content   string
		footer    string
	)
	// **************************** (1)告警标题处理 ****************************
	switch alerter.Status {
	case "firing":
		color = "red"
		header = "【容器集群】告警通知"
		footer = fmt.Sprintf(`**开始时间：**%s\n`, alerter.GetStartTime())
	case "resolved":
		color = "green"
		header = "【容器集群】告警恢复"
		footer = fmt.Sprintf(`**开始时间：**%s\n**结束时间：**%s\n`, alerter.GetStartTime(), alerter.GetEndTime())
	}

	// **************************** (2)告警内容处理 ****************************
	var (
		alertName        = alerter.Labels["alertname"]
		alertLevel       = alerter.Labels["level"]
		alertSummary     = alerter.Annotations["summary"]
		alertDescription = alerter.Annotations["description"]
		alertCluster     = alerter.Labels["robot_name"]
	)

	switch ns, ok := alerter.Labels["namespace"]; {
	case ok:
		rsp, err := i.k.DescNamespace(ctx, &k8s.DescNamespaceReq{ClusterName: alerter.Labels["robot_name"], NamespaceName: ns})

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
	card := fmt.Sprintf(alert.CardTemplate, timestamp, sign, content, footer, color, header)

	return []byte(card), nil
}

// RecoveryAction 集群自愈操作
func (i *Impl) RecoveryAction(ctx context.Context, req *alert.Alert, action string) error {
	//switch action {
	//case "GetPodRamUsage":
	//	in := &k8s.GetPodRamUsageReq{
	//		ClusterName: "",
	//		NodeName:    "",
	//	}
	//	_, err := i.k.GetPodRamUsage(ctx, in)
	//	if err != nil {
	//		return err
	//	}
	//
	//case "kill_pod":
	//	in := &k8s.KillTop1PodReq{
	//		ClusterName:   "",
	//		NamespaceName: "",
	//		PodName:       "",
	//	}
	//	i.k.KillTop1Pod(ctx, in)
	//}
	return nil
}
