package k8s

import (
	"context"
	"gitee.com/qiaogy91/InfraGenie/apps/resourcer"
)

const (
	AppName = "k8s"
)

type DescTypeRancherType int

const (
	FromProjectID DescTypeRancherType = iota + 1
	FromClusterAndNamespace
)

type Service interface {
	RpcServer
	SyncRancherProject(context.Context) error
	DescRancherProject(ctx context.Context, req *DescRancherProjectReq) (*resourcer.Project, error) // 查询项目详细
}

type DescRancherProjectReq struct {
	DescType    DescTypeRancherType
	ProjectID   string
	ClusterName string
	Namespace   string
}
