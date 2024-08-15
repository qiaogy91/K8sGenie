package k8s

import (
	"context"
	"gitee.com/qiaogy91/InfraGenie/apps/resourcer"
)

const (
	AppName = "k8s"
)

type Service interface {
	RpcServer
	RancherResourceSync(context.Context) error
	DescRancherProject(ctx context.Context, pid string) (*resourcer.Rancher, error)
}
