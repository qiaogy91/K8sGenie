package k8s

import "context"

const (
	AppName = "k8s"
)

type Service interface {
	RpcServer
	Extra
}

type GetDeploymentSummaryReq struct {
	ClusterName string `json:"clusterName"`
	Namespace   string `json:"namespace"`
}

type Extra interface {
	GetDeploymentSummary(ctx context.Context, req *GetDeploymentSummaryReq) error
}
