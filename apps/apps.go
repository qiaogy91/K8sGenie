package apps

import (
	_ "gitee.com/qiaogy91/K8sGenie/apps/rancher/api"
	// 01
	_ "gitee.com/qiaogy91/K8sGenie/apps/rancher/impl"
	// 02
	_ "gitee.com/qiaogy91/K8sGenie/apps/k8s/api"
	_ "gitee.com/qiaogy91/K8sGenie/apps/k8s/impl"
	// 03
	_ "gitee.com/qiaogy91/K8sGenie/apps/router/api"
	_ "gitee.com/qiaogy91/K8sGenie/apps/router/impl"
	// 04
	_ "gitee.com/qiaogy91/K8sGenie/apps/alert/api"
	_ "gitee.com/qiaogy91/K8sGenie/apps/alert/impl"
	// 05
	_ "gitee.com/qiaogy91/K8sGenie/apps/cron/impl"
)
