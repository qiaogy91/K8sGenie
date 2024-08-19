package impl_test

import (
	"context"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	_ "gitee.com/qiaogy91/K8sGenie/test"
	"google.golang.org/grpc"
	"testing"
)

var (
	c   k8s.Service
	ctx = context.Background()
)

type Stream struct {
	grpc.ServerStream
}

func (s *Stream) Send(w *k8s.WorkLoad) error {
	fmt.Printf("stream: %+v\n", w)
	return nil
}

func init() {
	c = ioc.GetController(k8s.AppName).(k8s.Service)
}

func TestImpl_TableCreate(t *testing.T) {
	if _, err := c.CreateTable(ctx, nil); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_SyncK8SWorkload(t *testing.T) {
	if err := c.SyncK8SWorkload(nil, &Stream{}); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_DescNamespace(t *testing.T) {
	ins, err := c.DescNamespace(ctx, &k8s.DescNamespaceReq{
		ClusterName:   "itcp-k8s-uat",
		NamespaceName: "djms-dev",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ins)
}
