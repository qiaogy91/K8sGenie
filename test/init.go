package test

import (
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
)

func init() {
	conf.LoadFromToml("/Users/qiaogy/GolandProjects/projects/Genies/K8sGenie/etc/conf.toml")
	// ioc 对象初始化，附加数据库连接池、附加业务实现对象
	ioc.InitController()
	ioc.InitHandler()
}
