package impl_test

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/feishu"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	_ "gitee.com/qiaogy91/K8sGenie/test"
	"testing"
)

var (
	ctx = context.Background()
	c   feishu.Service
)

func init() {
	c = ioc.GetController(feishu.AppName).(feishu.Service)
}

func TestImpl_CreateTable(t *testing.T) {
	if err := c.CreateTable(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestImpl_CallAndAudit(t *testing.T) {
	req := &feishu.CallAndAuditRequest{
		Users: map[string]string{
			"shipengfei":    "18629648489",
			"qiaoguanyu":    "18192106883",
			"zhangyuanzhao": "15191896504",
		},
		Content: "{\"text\":\"电话加急测试001，我是告警信息\"}",
	}

	if err := c.CallAndAudit(ctx, req); err != nil {
		t.Fatal(err)
	}
}
