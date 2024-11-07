package impl_test

import (
	"context"
	"encoding/json"
	"fmt"
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

func TestImpl_GetAccessToken(t *testing.T) {
	req := &feishu.GetAccessTokenRequest{
		AppId:     "cli_9e063da4943a1008",
		AppSecret: "DTRJz5iWdqbjZEbCE6FB8waYQzYEMexH",
	}

	bs, _ := json.MarshalIndent(req, "", " ")
	t.Logf("%s", bs)

	ins, err := c.GetAccessToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins)
}

func TestImpl_FetchUserId(t *testing.T) {
	TestImpl_GetAccessToken(t)

	c.SetToken(ctx, "t-g101b7h8NDPXR7JNOQX4QYCMVCTOLC7BHOF7LFPJ")

	req := &feishu.FetchUserIdRequest{Mobiles: []string{"18192106883"}}
	ins, err := c.FetchUserId(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", ins.Data.UserList)
}

func TestImpl_SendMessage(t *testing.T) {
	// 获取token
	getTokenReq := &feishu.GetAccessTokenRequest{
		AppId:     "cli_9e063da4943a1008",
		AppSecret: "DTRJz5iWdqbjZEbCE6FB8waYQzYEMexH",
	}
	if _, err := c.GetAccessToken(ctx, getTokenReq); err != nil {
		t.Fatal(err)
	}

	// 发送消息
	sendMessageReq := &feishu.SendMessageRequest{
		ReceiveID: "ou_e0b1c282de786baf239e412e1ea4edbd",
		Content:   fmt.Sprintf(`{"text": "%s"}`, "qiaogy"),
		MsgType:   "text",
	}
	msg, err := c.SendMessage(ctx, sendMessageReq)
	if err != nil {
		t.Fatal(err)
	}

	// 消息加急
	urgentReq := &feishu.UrgentCallRequest{
		UserIdList: []string{"ou_e0b1c282de786baf239e412e1ea4edbd"},
		MessageId:  msg.Data.MessageId,
	}

	out, err := c.UrgentCall(ctx, urgentReq)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", out)
}
