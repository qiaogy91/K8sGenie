package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/alert"
	"gitee.com/qiaogy91/K8sGenie/apps/alert/impl/provider"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/router"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"google.golang.org/grpc"
)

var _ alert.Service = &Impl{}

type Impl struct {
	alert.UnimplementedRpcServer
	r router.Service
	c provider.Service
	k k8s.Service
}

func (i *Impl) Name() string {
	return alert.AppName
}

func (i *Impl) Init() error {
	i.r = ioc.GetController(router.AppName).(router.Service)
	i.k = ioc.GetController(k8s.AppName).(k8s.Service)
	i.c = provider.NewHttpClient()
	return nil
}

func (i *Impl) RegistrySvc(g *grpc.Server) {
	alert.RegisterRpcServer(g, i)
}

func init() {
	ioc.RegistryController(&Impl{})
}
