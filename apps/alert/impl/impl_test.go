package impl_test

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/alert"
	"gitee.com/qiaogy91/K8sGenie/apps/router"
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

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_HandlerAlert(t *testing.T) {
	alerts := &alert.HandlerAlertReq{
		Alerts: []*alert.Alert{
			//	// 集群级别告警
			//	{
			//		Status:   "firing",
			//		StartsAt: "2024-08-20T07:45:54.632Z",
			//		EndsAt:   "2024-08-20T07:45:54.632Z",
			//		Labels: map[string]string{
			//			"cluster_name": "itcp-k8s-uat",
			//			"alertname":    "alertname test",
			//			"level":        "level test",
			//		},
			//		Annotations: map[string]string{
			//			"description": "test desc",
			//			"record":      "集群级别告警",
			//		},
			//	},
			// 名称空间级别告警
			{
				Status:   "firing",
				StartsAt: "2024-09-20T07:45:55.632Z",
				EndsAt:   "2024-08-20T07:45:59.632Z",
				Labels: map[string]string{
					"cluster_name": "crc-itcp-prd-02",
					"namespace":    "asmp-crcprd",
					"alertname":    "alertname test",
					"level":        "level test",
				},
				Annotations: map[string]string{
					"description": "test t005",
					"record":      "名称空间级项目告警",
				},
			},

			// 名称空间级别告警，但是未归属到项目
			//{
			//	Status:   "firing",
			//	StartsAt: "2024-08-20T07:45:54.632Z",
			//	EndsAt:   "2024-08-20T07:45:54.632Z",
			//	Labels: map[string]string{
			//		"cluster_name": "itcp-k8s-uat",
			//		"namespace":    "default",
			//		"alertname":    "alertname test",
			//		"level":        "level test",
			//	},
			//	Annotations: map[string]string{
			//		"description": "test desc",
			//		"record":     "名称空间级项目告警",
			//	},
			//},
		},
	}
	rsp, err := c.HandlerAlert(ctx, alerts)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", rsp)
}

func TestImpl_UrgentAlertCall(t *testing.T) {
	req := &alert.UrgentAlertCallRequest{
		Users: &router.Users{
			Username: "qiaogy",
			Phone:    "18192106883",
		},
		Alert: nil,
	}
	ins, err := c.UrgentAlertCall(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ins)
}
