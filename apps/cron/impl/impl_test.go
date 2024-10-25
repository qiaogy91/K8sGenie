package impl_test

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/cron"
	"gitee.com/qiaogy91/K8sGenie/apps/cron/impl"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	_ "gitee.com/qiaogy91/K8sGenie/test"
	"testing"
)

var (
	ctx = context.Background()
	//c   cron.Service
	c = ioc.GetController(cron.AppName).(*impl.Impl)
)

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_GetCluster(t *testing.T) {
	ins, err := c.GetCluster(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, cluster := range *ins {
		t.Logf("%+v", cluster)
	}
}

func TestImpl_GetProject(t *testing.T) {
	ins, err := c.GetProject(ctx, &cron.GetProjectReq{ClusterName: "itcp-k8s-xc-prd"})
	if err != nil {
		t.Fatal(err)
	}

	for _, cluster := range ins.Items {
		t.Logf("%+v", cluster)
	}
}

func TestImpl_GetSummary(t *testing.T) {
	ins, err := c.GetProject(ctx, &cron.GetProjectReq{ClusterName: "itcp-k8s-xc-prd"})
	if err != nil {
		t.Fatal(err)
	}

	if err := c.HavingNamespace(ctx, ins.Items[1]); err != nil {
		t.Fatal(err)
	}

}
func TestImpl_CreateNamespaceRecord(t *testing.T) {
	if err := c.CreateNamespaceRecord(ctx, &cron.CreateNamespaceRecordRequest{}); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_CreateProjectRecord(t *testing.T) {
	if err := c.CreateProjectRecord(ctx, &cron.CreateProjectRecordRequest{}); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_CreateLineRecord(t *testing.T) {
	if err := c.CreateLineRecord(ctx, &cron.CreateLineRecordRequest{}); err != nil {
		t.Fatal(err)
	}
}
