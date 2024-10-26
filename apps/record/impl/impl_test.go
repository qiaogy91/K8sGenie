package impl_test

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/record"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	_ "gitee.com/qiaogy91/K8sGenie/test"
	"testing"
)

var (
	ctx = context.Background()
	c   = ioc.GetController(record.AppName).(record.Service)
)

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

//	func TestImpl_GetCluster(t *testing.T) {
//		ins, err := c.GetCluster(ctx)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		for _, cluster := range *ins {
//			t.Logf("%+v", cluster)
//		}
//	}
//
//	func TestImpl_GetProject(t *testing.T) {
//		ins, err := c.GetProject(ctx, &record.GetProjectReq{ClusterName: "itcp-k8s-xc-prd"})
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		for _, cluster := range ins.Items {
//			t.Logf("%+v", cluster)
//		}
//	}
//
//	func TestImpl_GetSummary(t *testing.T) {
//		ins, err := c.GetProject(ctx, &record.GetProjectReq{ClusterName: "itcp-k8s-xc-prd"})
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		if err := c.HavingNamespace(ctx, ins.Items[1]); err != nil {
//			t.Fatal(err)
//		}
//
// }
func TestImpl_CreateNamespaceRecord(t *testing.T) {
	if err := c.CreateNamespaceRecord(ctx, &record.CreateNamespaceRecordRequest{}); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_CreateProjectRecord(t *testing.T) {
	if err := c.CreateProjectRecord(ctx, &record.CreateProjectRecordRequest{}); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_CreateLineRecord(t *testing.T) {
	if err := c.CreateLineRecord(ctx, &record.CreateLineRecordRequest{}); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_QueryLineRecord(t *testing.T) {
	req := &record.QueryLineRecordRequest{
		Month:       "2024-10",
		ClusterName: "crc-itcp-test03",
	}

	ins, err := c.QueryLineRecord(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range *ins {
		t.Logf("%+v", item)
	}
}

func TestImpl_QueryProjectRecord(t *testing.T) {
	req := &record.QueryProjectRecordRequest{
		Month:       "2024-10",
		ProjectLine: "智能风控",
		ClusterName: "crc-itcp-test03",
	}
	ins, err := c.QueryProjectRecord(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range *ins {
		t.Logf("%+v", item)
	}
}

func TestImpl_QueryNamespaceRecord(t *testing.T) {
	req := &record.QueryNamespaceRecordRequest{
		CreatedTime: "2024-10-25",
		ProjectCode: "ehs",
		ProjectLine: "智能风控",
		ClusterName: "crc-itcp-test03",
	}
	ins, err := c.QueryNamespaceRecord(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range *ins {
		t.Logf("%+v", item)
	}
}
