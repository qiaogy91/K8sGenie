package feishu

import (
	"context"
)

type Service interface {
	GetAccessToken(ctx context.Context, req *GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	FetchUserId(ctx context.Context, req *FetchUserIdRequest) (*FetchUserIdResponse, error)
	SendMessage(ctx context.Context, req *SendMessageRequest) (*SendMessageResponse, error)
	UrgentCall(ctx context.Context, req *UrgentCallRequest) (*UrgentCallResponse, error)
}

// UrgentCallRequest UrgentCallResponse 消息加急
type UrgentCallRequest struct {
	UserIdList []string `json:"user_id_list"`
	MessageId  string   `json:"message_id"`
}
type UrgentCallResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// FetchUserIdRequest 获取用户ID
type FetchUserIdRequest struct {
	Mobiles []string `json:"mobiles"`
}

type FetchUserIdResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		UserList []struct {
			Mobile  string `json:"mobile"`
			OpenID  string `json:"open_id"`
			UnionID string `json:"union_id"`
			UserID  string `json:"user_id"`
		} `json:"user_list"`
	} `json:"data"`
}

// GetAccessTokenRequest 获取应用Token
type GetAccessTokenRequest struct {
	AppId     string `json:"app_id" validate:"required"`
	AppSecret string `json:"app_secret" validate:"required"`
}

type GetAccessTokenResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

// SendMessageRequest 发送消息
type SendMessageRequest struct {
	ReceiveID string `json:"receive_id"`
	Content   string `json:"content"`
	MsgType   string `json:"msg_type"`
}

type SendMessageResponse struct {
	Code int `json:"code"`
	Data struct {
		ChartID    string `json:"chat_id"`
		CreateTime string `json:"create_time"`
		MessageId  string `json:"message_id"`
	} `json:"data"`
	Msg string `json:"msg"`
}
