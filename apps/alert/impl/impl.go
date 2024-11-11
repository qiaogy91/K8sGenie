package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/alert"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/router"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"time"
)

var _ alert.Service = &Impl{}

type Impl struct {
	alert.UnimplementedRpcServer
	r  router.Service
	k  k8s.Service
	c  *lark.Client
	db *gorm.DB
}

func (i *Impl) Name() string {
	return alert.AppName
}

func (i *Impl) Init() error {
	i.r = ioc.GetController(router.AppName).(router.Service)
	i.k = ioc.GetController(k8s.AppName).(k8s.Service)
	i.db = conf.C().GetMysqlPool()
	i.c = lark.NewClient(
		"cli_9e063da4943a1008",
		"DTRJz5iWdqbjZEbCE6FB8waYQzYEMexH",
		lark.WithEnableTokenCache(true),    // 自动缓存 AccessToken
		lark.WithLogReqAtDebug(true),       // 开启 Http 请求参数和响应参数的日志打印开
		lark.WithReqTimeout(5*time.Second), // 请求调用的超时时间，0表示永不超时
		lark.WithOpenBaseUrl("https://open.rwork.crc.com.cn"),
	)
	return nil
}

func (i *Impl) RegistrySvc(g *grpc.Server) {
	alert.RegisterRpcServer(g, i)
}

func init() {
	ioc.RegistryController(&Impl{})
}
