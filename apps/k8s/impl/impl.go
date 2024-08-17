package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

var _ k8s.Service = &Impl{}

type Impl struct {
	db *gorm.DB
	k8s.UnimplementedRpcServer
	cs map[string]*kubernetes.Clientset
}

func (i *Impl) Name() string {
	return k8s.AppName
}

func (i *Impl) Init() error {
	// db
	i.db = conf.C().GetMysqlPool()

	// kubeConf 文件
	apiConf, err := clientcmd.LoadFromFile(conf.C().Rancher.KubeFile)
	if err != nil {
		panic(err)
	}

	// 将各个集群的名称、以及对应的clientSet 做成Map
	clientMap := make(map[string]*kubernetes.Clientset)
	for name, _ := range apiConf.Contexts {
		// 获取restConf
		restConf, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*api.Config, error) {
			apiConf.CurrentContext = name  // apiConf context 切换
			return apiConf.DeepCopy(), nil // 返回切换后的 apiConf 副本
		})

		if err != nil {
			panic(err)
		}

		// 获取clientSet
		client, err := kubernetes.NewForConfig(restConf)
		if err != nil {
			panic(err)
		}
		// 放入map
		clientMap[name] = client
	}

	i.cs = clientMap

	return nil
}

func (i *Impl) RegistrySvc(g *grpc.Server) {
	k8s.RegisterRpcServer(g, i)
}

func init() {
	ioc.RegistryController(&Impl{})
}
