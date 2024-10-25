package impl

import (
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/metrics/pkg/client/clientset/versioned"
)

var _ k8s.Service = &Impl{}

type Impl struct {
	db *gorm.DB
	k8s.UnimplementedRpcServer
	cs map[string]*kubernetes.Clientset
	ms map[string]*versioned.Clientset
	rc rancher.Service
}

func (i *Impl) Name() string {
	return k8s.AppName
}

func (i *Impl) Init() error {
	// db
	i.db = conf.C().GetMysqlPool()

	// rancher svc
	i.rc = ioc.GetController(rancher.AppName).(rancher.Service)

	// kubeConf 文件
	apiConf, err := clientcmd.LoadFromFile(conf.C().Rancher.KubeFile)
	if err != nil {
		panic(err)
	}

	// 初始化map
	i.cs = make(map[string]*kubernetes.Clientset)
	i.ms = make(map[string]*versioned.Clientset)

	// 将各个集群的名称、以及对应的clientSet 做成Map
	for name, _ := range apiConf.Contexts {
		if name == "docker-desktop" {
			continue
		}
		// ******************** 获取restConf ********************
		restConf, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*api.Config, error) {
			apiConf.CurrentContext = name  // apiConf context 切换
			return apiConf.DeepCopy(), nil // 返回切换后的 apiConf 副本
		})
		if err != nil {
			panic(err)
		}

		// ******************** 获取clientSet ********************
		client, err := kubernetes.NewForConfig(restConf)
		if err != nil {
			panic(err)
		}
		i.cs[name] = client

		// ******************** 获取metricClient ********************
		metricsClient, err := versioned.NewForConfig(restConf)
		if err != nil {
			panic(err)
		}
		i.ms[name] = metricsClient
	}

	return nil
}

func (i *Impl) RegistrySvc(g *grpc.Server) {
	k8s.RegisterRpcServer(g, i)
}

func init() {
	ioc.RegistryController(&Impl{})
}
