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
func TestImpl_SyncRancherResource(t *testing.T) {
	if err := c.SyncRancherProject(nil, &MockStream{}); err != nil {
		t.Fatal(err)
	}
}
