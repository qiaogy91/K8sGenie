package impl_test

import (
	"context"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	_ "gitee.com/qiaogy91/K8sGenie/test"
	"google.golang.org/grpc"
	"testing"
)

var (
	c   rancher.Service
	ctx = context.Background()
)

type MockStream struct {
	grpc.ServerStream
}

func (s *MockStream) Send(r *rancher.Project) error {
	fmt.Printf("stream: %+v\n", r)
	return nil
}

func init() {
	c = ioc.GetController(rancher.AppName).(rancher.Service)
}

// TestImpl_CreateTable 创建表结构
func TestImpl_CreateTable(t *testing.T) {
	if _, err := c.CreateTable(ctx, nil); err != nil {
		t.Fatal(err)
	}
}

// TestImpl_SyncRancherResource 同步项目信息
func TestImpl_SyncResource(t *testing.T) {
	if err := c.SyncProject(nil, &MockStream{}); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_DescProject(t *testing.T) {
	req := &rancher.DescProjectReq{
		//DescType: rancher.DESC_TYPE_DESC_TYPE_PROJECT_ID,
		//KeyWord:  "c-m-nnljl7c9:p-lv2jz",
		//DescType: rancher.DESC_TYPE_DESC_TYPE_PROJECT_CODE,
		//KeyWord:  "ehs",
		DescType: rancher.DESC_TYPE_DESC_TYPE_PROJECT_CODE,
		KeyWord:  "asmp",
	}
	ins, err := c.DescProject(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}

func TestImpl_QueryProject(t *testing.T) {
	req := &rancher.QueryProjectReq{
		//QueryType: rancher.QUERY_TYPE_QUERY_TYPE_CLUSTER_CODE,
		//KeyWord:   "itcp-k8s-uat",

		QueryType: rancher.QUERY_TYPE_QUERY_TYPE_PROJECT_LINE,
		KeyWord:   "协同效率",
	}
	ins, err := c.QueryProject(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}
