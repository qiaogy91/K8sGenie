package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/apps/record"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"time"
)

const AppName = "cronManager"

type Impl struct {
	rancherSvc rancher.Service
	k8sSvc     k8s.Service
	recordSvc  record.Service
	cron       *cron.Cron
}

func (i *Impl) RegistrySvc(g *grpc.Server) {}
func (i *Impl) Name() string               { return AppName }
func (i *Impl) Init() error {
	i.rancherSvc = ioc.GetController(rancher.AppName).(rancher.Service)
	i.k8sSvc = ioc.GetController(k8s.AppName).(k8s.Service)
	i.recordSvc = ioc.GetController(record.AppName).(record.Service)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	i.cron = cron.New(cron.WithSeconds(), cron.WithLocation(loc))

	i.cron.Start()
	// 0 0 3 */3 * *"
	if _, err := i.cron.AddJob("0 */30 * * * *", i); err != nil {
		panic(err)
	}
	return nil
}

func init() {
	ioc.RegistryController(&Impl{})
}
