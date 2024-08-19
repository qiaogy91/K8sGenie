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
	req := &router.Spec{
		ClusterName:  "itcp-k8s-uat",
		ProjectId:    "c-m-nnljl7c9:p-76q6t",
		WebhookUrl:   "www.example.com",
		WebhookToken: "***",
	}
	ins, err := c.CreateRoute(ctx, req)
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

func TestImpl_QueryRoute(t *testing.T) {
	req := &router.QueryRouteReq{
		//SearchType: router.SEARCH_TYPE_SEARCH_TYPE_ROUTER_ID,
		//SearchType: router.SEARCH_TYPE_SEARCH_TYPE_CLUSTER_NAME,
		//SearchType: router.SEARCH_TYPE_SEARCH_TYPE_PROJECT_LINE,
		SearchType: router.SEARCH_TYPE_SEARCH_TYPE_PROJECT_CODE,
		//SearchType: router.SEARCH_TYPE_SEARCH_TYPE_PROJECT_DESC,
		KeyWord: "AJ",
	}

	ins, err := c.QueryRoute(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ins)
}
