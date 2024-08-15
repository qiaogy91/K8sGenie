package impl

import (
	"gitee.com/qiaogy91/InfraGenie/apps/resourcer"
	resourceClient "gitee.com/qiaogy91/InfraGenie/client"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var _ k8s.Service = &Impl{}

type Impl struct {
	db *gorm.DB
	k8s.UnimplementedRpcServer
	rc resourcer.RpcClient
	kc *kubernetes.Clientset
}

func (i *Impl) Name() string {
	return k8s.AppName
}

func (i *Impl) Init() error {
	// db
	i.db = conf.C().GetMysqlPool().Debug()

	// k8s client
	restConf, err := clientcmd.BuildConfigFromFlags("", conf.C().Rancher.KubeFile)
	if err != nil {
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		panic(err)
	}

	i.kc = clientSet

	// rancher client
	i.rc = resourceClient.NewClient(conf.C().RancherAddr(), conf.C().Rancher.Token)
	return nil
}

func (i *Impl) RegistrySvc(g *grpc.Server) {
	k8s.RegisterRpcServer(g, i)
}

func init() {
	ioc.RegistryController(&Impl{})
}
