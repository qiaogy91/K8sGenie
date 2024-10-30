package impl_test

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/cron"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	_ "gitee.com/qiaogy91/K8sGenie/test"
	"testing"
)

var (
	c   cron.Service
	ctx = context.Background()
)

func init() {
	c = ioc.GetController(cron.AppName).(cron.Service)
}

func TestImpl_Run(t *testing.T) {
	c.Run()
}
