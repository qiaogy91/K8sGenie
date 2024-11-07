package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/feishu"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"google.golang.org/grpc"
)

var _ feishu.Service = &Impl{}

type Impl struct {
	AppID     string `json:"appID"`
	AppSecret string `json:"appSecret"`
	Token     string `json:"token"` // 获取的临时ID
}

func (i *Impl) RegistrySvc(g *grpc.Server) {}
func (i *Impl) Name() string               { return feishu.AppName }
func (i *Impl) Init() error {
	return nil
}

func init() {
	ioc.RegistryController(&Impl{})
}
