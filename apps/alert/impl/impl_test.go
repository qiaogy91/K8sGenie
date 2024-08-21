package impl_test

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/alert"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	_ "gitee.com/qiaogy91/K8sGenie/test"
	"testing"
)

var (
	ctx = context.Background()
	c   alert.Service
)

func init() {
	c = ioc.GetController(alert.AppName).(alert.Service)
}

func TestImpl_HandlerAlert(t *testing.T) {
	alerts := &alert.HandlerAlertReq{
		Alerts: []*alert.Alert{
			// 集群级别告警
			{
				Status:   "firing",
				StartsAt: "2024-08-20T07:45:54.632Z",
				EndsAt:   "2024-08-20T07:45:54.632Z",
				Labels: map[string]string{
					"robot_name": "itcp-k8s-uat",
					"alertname":  "alertname test",
					"level":      "level test",
				},
				Annotations: map[string]string{
					"description": "test desc",
					"summary":     "集群级别告警",
				},
			},
			// 名称空间级别告警
			{
				Status:   "firing",
				StartsAt: "2024-08-20T07:45:54.632Z",
				EndsAt:   "2024-08-20T07:45:54.632Z",
				Labels: map[string]string{
					"robot_name": "itcp-k8s-uat",
					"namespace":  "crc-yourh-dev",
					"alertname":  "alertname test",
					"level":      "level test",
				},
				Annotations: map[string]string{
					"description": "test desc",
					"summary":     "名称空间级项目告警",
				},
			},

			// 名称空间级别告警，但是未归属到项目
			{
				Status:   "firing",
				StartsAt: "2024-08-20T07:45:54.632Z",
				EndsAt:   "2024-08-20T07:45:54.632Z",
				Labels: map[string]string{
					"robot_name": "itcp-k8s-uat",
					"namespace":  "default",
					"alertname":  "alertname test",
					"level":      "level test",
				},
				Annotations: map[string]string{
					"description": "test desc",
					"summary":     "名称空间级项目告警",
				},
			},
		},
	}
	rsp, err := c.HandlerAlert(ctx, alerts)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", rsp)
}
