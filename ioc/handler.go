package ioc

import (
	"gitee.com/qiaogy91/K8sGenie/common"
)

type HandlerObject interface {
	Name() string
	Init() error
	RegistryRoute(p string)
}

var handlerContainer = map[string]HandlerObject{}

func RegistryHandlerObject(o HandlerObject) {
	handlerContainer[o.Name()] = o
}

func GetHandlerObject(n string) HandlerObject {
	v, ok := handlerContainer[n]
	if !ok {
		panic("未找到Handler: " + n)
	}
	return v
}

// InitHandler 添加对应的svc 实现
func InitHandler() {
	for name, o := range handlerContainer {
		if err := o.Init(); err != nil {
			panic("handler 初始化错误: " + err.Error())
		}
		common.L().Info().Msgf("%s handler object init success", name)
	}
}

func LoadHandler() map[string]HandlerObject {
	return handlerContainer
}
