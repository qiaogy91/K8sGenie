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

func TestImpl_RancherCreateTable(t *testing.T) {
	if _, err := c.TableCreate(ctx, nil); err != nil {
		t.Fatal(err)
	}
}
func TestImpl_RancherResourceSync(t *testing.T) {
	if err := c.RancherResourceSync(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_DescRancherProject(t *testing.T) {
	ins, err := c.DescRancherProject(ctx, "c-m-zcjwcxrr:p-zm2vx")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}

func TestImpl_K8SResourceSync(t *testing.T) {
	if err := c.K8SResourceSync(nil, &Stream{}); err != nil {
		t.Fatal(err)
	}
}
