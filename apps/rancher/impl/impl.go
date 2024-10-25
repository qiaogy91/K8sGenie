package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"github.com/rancher/norman/clientbase"
	managementV3 "github.com/rancher/types/client/management/v3"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var _ rancher.Service = &Impl{}

type Impl struct {
	db  *gorm.DB
	adm []string // blacklist 用户在这个名单上的，不进行记录，过滤掉管理员
	rcs map[string]*managementV3.Client
	rancher.UnimplementedRpcServer
}

func (i *Impl) Name() string {
	return rancher.AppName
}

func (i *Impl) Init() error {
	i.adm = []string{"qiaoguanyu", "zhangyuanzhao", "shipengfei", "admin"}

	i.db = conf.C().GetMysqlPool()

	// init rancher client
	client1, err := managementV3.NewClient(&clientbase.ClientOpts{
		URL:        "https://itpub.crdigital.com.cn/v3",
		TokenKey:   "token-wmlpc:bfrtz22ttcfzjvqzhbntn4lwp58v4794hn5nzh5xvhd86dz74k2shb",
		Insecure:   true,
		Timeout:    5 * time.Second,
		HTTPClient: &http.Client{Transport: &http.Transport{}},
	})
	if err != nil {
		panic(err)
	}
	i.rcs["v2.0"] = client1

	client2, err := managementV3.NewClient(&clientbase.ClientOpts{
		URL:        "https://k8spro.crcloud.com/v3",
		TokenKey:   "token-gw6m6:hnxc6jlj2m4ljfmzqsn42j68sqd4zblr4t6fkldrfljj2fs7f5kzwr",
		Insecure:   true,
		Timeout:    5 * time.Second,
		HTTPClient: &http.Client{Transport: &http.Transport{}},
	})
	if err != nil {
		panic(err)
	}
	i.rcs["v1.0"] = client2

	return nil
}

func (i *Impl) RegistrySvc(g *grpc.Server) {
	rancher.RegisterRpcServer(g, i)
}

func init() {

	ioc.RegistryController(&Impl{rcs: make(map[string]*managementV3.Client)})
}
