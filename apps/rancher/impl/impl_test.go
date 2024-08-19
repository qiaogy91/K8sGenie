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

func TestImpl_QueryProject(t *testing.T) {
	req := &rancher.QueryProjectReq{
		// 根据注解查询
		//SearchType: rancher.SEARCH_TYPE_SEARCH_TYPE_ANNOTATION,
		//KeyWord:    "c-m-nnljl7c9:p-76q6t",
		// 根据ID 查询
		SearchType: rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_ID,
		KeyWord:    "c-m-nnljl7c9:p-76q6t",
		// 根据项目名称查询
		//SearchType: rancher.SEARCH_TYPE_SEARCH_TYPE_CLUSTER_NAME,
		//KeyWord:    "itcp-k8s-uat",
		//SearchType: rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_DESC,
		//KeyWord:    "工作",
		//SearchType: rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_LINE,
		//KeyWord:    "协同",
		//SearchType: rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_CODE,
		//KeyWord:    "runwork",
	}
	ins, err := c.QueryProject(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}
