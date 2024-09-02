package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/apps/router"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var _ router.Service = &Impl{}

type Impl struct {
	router.UnimplementedRpcServer
	db *gorm.DB
	rc rancher.Service
	kc k8s.Service
}

func (i *Impl) Name() string {
	return router.AppName
}

func (i *Impl) Init() error {
	i.db = conf.C().GetMysqlPool().Debug()
	i.rc = ioc.GetController(rancher.AppName).(rancher.Service)
	i.kc = ioc.GetController(k8s.AppName).(k8s.Service)
	return nil
}

func (i *Impl) RegistrySvc(g *grpc.Server) {
	router.RegisterRpcServer(g, i)
}

func init() {
	ioc.RegistryController(&Impl{})
}
