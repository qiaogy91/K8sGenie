package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/record"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var _ record.Service = &Impl{}

type Impl struct {
	db *gorm.DB
}

func (i *Impl) Name() string {
	return record.AppName
}
func (i *Impl) Init() error {
	i.db = conf.C().GetMysqlPool()
	return nil
}

func (i *Impl) RegistrySvc(g *grpc.Server) {}

func init() {
	ioc.RegistryController(&Impl{})
}
