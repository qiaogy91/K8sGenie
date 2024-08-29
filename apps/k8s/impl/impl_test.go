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

func TestImpl_GetPodRamUsage(t *testing.T) {
	req := &k8s.GetPodRamUsageReq{
		ClusterName: "itcp-k8s-prd",
		NodeName:    "itcp-k8s-prd-worker-02",
	}
	ins, err := c.GetPodRamUsage(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range ins.Items {
		t.Logf("%+v", item)
	}
}

func TestImpl_KillTop1Pod(t *testing.T) {
	req := &k8s.KillTop1PodReq{
		ClusterName:   "itcp-k8s-uat",
		NamespaceName: "itcp-test",
		PodName:       "client-nz9fd",
	}
	if _, err := c.KillTop1Pod(ctx, req); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_QueryK8SWorkload(t *testing.T) {
	req := &k8s.QueryK8SWorkloadReq{
		Type:    k8s.SEARCH_TYPE_SEARCH_TYPE_WORKLOAD_NAME,
		Keyword: "runwork",
	}

	ins, err := c.QueryK8SWorkload(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	for _, pod := range ins.Items {
		t.Logf("%+v", pod)
	}
}
