package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/feishu"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var _ feishu.Service = &Impl{}

type Impl struct {
	AppID     string `json:"appID"`
	AppSecret string `json:"appSecret"`
	Token     string `json:"token"` // 获取的临时ID
	db        *gorm.DB
}

func (i *Impl) RegistrySvc(g *grpc.Server) {}
func (i *Impl) Name() string               { return feishu.AppName }
func (i *Impl) Init() error {
	i.AppID = conf.C().QARobot.AppID
	i.AppSecret = conf.C().QARobot.AppSecret
	i.db = conf.C().GetMysqlPool()
	return nil
}

func init() {
	ioc.RegistryController(&Impl{})
}
