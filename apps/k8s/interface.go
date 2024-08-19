package k8s

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
}
