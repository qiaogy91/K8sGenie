package ioc

import (
	"gitee.com/qiaogy91/K8sGenie/common"
	"google.golang.org/grpc"
)

type ControllerObject interface {
	Name() string
	Init() error
	RegistrySvc(g *grpc.Server)
}

var controllerContainer = map[string]ControllerObject{}

func RegistryController(o ControllerObject) {
	controllerContainer[o.Name()] = o
}

func GetController(n string) ControllerObject {
	o, ok := controllerContainer[n]
	if !ok {
		panic("没有找到目标Controller: " + n)
	}
	return o
}

// InitController 初始化所有的Impl 实现，内部添加数据库连接池
func InitController() {
	for name, o := range controllerContainer {
		// 添加数据库连接池
		if err := o.Init(); err != nil {
			panic("Controller 初始化失败: " + o.Name())
		}
		common.L().Info().Msgf("%s controller object init success", name)
	}
}

func LoadController() map[string]ControllerObject {
	return controllerContainer
}
