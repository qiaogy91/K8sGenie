package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/cron"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	cronv3 "github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"time"
)

var _ cron.Service = &Impl{}

type Impl struct {
	db   *gorm.DB
	cron *cronv3.Cron
}

func (i *Impl) Name() string {
	return cron.AppName
}
func (i *Impl) Init() error {
	i.db = conf.C().GetMysqlPool()
	loc, _ := time.LoadLocation("Asia/Shanghai")

	i.cron = cronv3.New(cronv3.WithLocation(loc), cronv3.WithSeconds())
	i.cron.Start()
	// 0 0 3 */3 * *"
	if _, err := i.cron.AddJob("0 */3 * * * *", i); err != nil {
		panic(err)
	}
	return nil
}

func (i *Impl) RegistrySvc(g *grpc.Server) {}

func init() {
	ioc.RegistryController(&Impl{})
}
