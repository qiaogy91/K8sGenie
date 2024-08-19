package provider_test

import (
	"context"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/alert/impl/provider"
	"testing"
)

var (
	ctx = context.Background()
	c   provider.Service
)

func init() {
	c = provider.NewHttpClient()
}

func TestHttpClient_SendCard(t *testing.T) {
	req := fmt.Sprintf(provider.CardTemplate)
	rsp := c.SendCard(ctx, "https://open.rwork.crc.com.cn/open-apis/bot/v2/hook/327a89a7-cce9-42f9-8dfe-2deb8b4413b7", []byte(req))

	t.Logf("%s", rsp)
}
