package impl_test

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/router"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	_ "gitee.com/qiaogy91/K8sGenie/test"
	"testing"
)

var (
	ctx = context.Background()
	c   router.Service
)

func init() {
	c = ioc.GetController(router.AppName).(router.Service)
}

func TestImpl_CreateTable(t *testing.T) {
	if _, err := c.CreateTable(ctx, nil); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_CreateRoute(t *testing.T) {
	// 集群级别告警
	req1 := &router.Spec{
		Identity:     "itcp-k8s-uat",
		WebhookUrl:   "https://open.rwork.crc.com.cn/open-apis/bot/v2/hook/680fc048-24d8-413d-a3c2-4d0297e66048",
		WebhookToken: "DUJbdqoPWa94mBdsD6HSfb",
		Users: []*router.Users{
			{Username: "qiaogy", Phone: "18192106883"},
			{Username: "liangchuanj", Phone: "15302743009"},
		},
	}
	// 项目级别告警
	req2 := &router.Spec{
		Identity:     "c-m-nnljl7c9:p-76q6t",
		WebhookUrl:   "https://open.rwork.crc.com.cn/open-apis/bot/v2/hook/395fb852-7884-45da-880c-e2290ad2698e",
		WebhookToken: "w8mKUHSZi4nEoMRqvVtFCb",
		Users: []*router.Users{
			{Username: "qiaogy", Phone: "18192106883"},
			{Username: "liangchuanj", Phone: "15302743009"},
		},
	}

	ins, err := c.CreateRoute(ctx, req1)
	ins, err = c.CreateRoute(ctx, req2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}

func TestImpl_DeleteRoute(t *testing.T) {
	ins, err := c.DeleteRoute(ctx, &router.DeleteRouteReq{Id: "3"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}

func TestImpl_AlertRoute(t *testing.T) {
	req := &router.AlertRouteReq{
		ClusterName:   "itcp-k8s-uat",
		NamespaceName: "crc-yourh-dev",
	}
	ins, err := c.AlertRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}

func TestImpl_QueryRoute(t *testing.T) {
	req := &router.QueryRouteReq{
		Type:    router.QUERY_TYPE_QUERY_TYPE_BY_ID,
		Keyword: "2",
		//Type:    router.QUERY_TYPE_QUERY_TYPE_BY_PROJECT_LINE,
		//Keyword: "创新学习",
	}

	ins, err := c.QueryRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}
