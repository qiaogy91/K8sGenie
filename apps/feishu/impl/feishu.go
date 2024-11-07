package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/feishu"
	"github.com/go-playground/validator"
	"io"
	"net/http"
)

func (i *Impl) GetAccessToken(ctx context.Context, r *feishu.GetAccessTokenRequest) (*feishu.GetAccessTokenResponse, error) {
	if err := validator.New().Struct(r); err != nil {
		return nil, err
	}

	// 构造请求
	bs, _ := json.Marshal(r)
	req, err := http.NewRequest("POST", feishu.AccessTokenUrl, bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rsp.Body.Close() }()

	// 解码响应
	ins := &feishu.GetAccessTokenResponse{}
	content, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(content, ins); err != nil {
		return nil, err
	}

	if ins.Code == 0 {
		i.Token = ins.TenantAccessToken
	}
	return ins, nil
}

func (i *Impl) FetchUserId(ctx context.Context, r *feishu.FetchUserIdRequest) (*feishu.FetchUserIdResponse, error) {
	if err := validator.New().Struct(r); err != nil {
		return nil, err
	}

	// 构造请求
	bs, _ := json.Marshal(r)
	req, err := http.NewRequest("POST", feishu.UserOpenIdUrl, bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+i.Token)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// 解码响应
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ins := &feishu.FetchUserIdResponse{}
	if err := json.Unmarshal(content, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) SendMessage(ctx context.Context, r *feishu.SendMessageRequest) (*feishu.SendMessageResponse, error) {
	if err := validator.New().Struct(r); err != nil {
		return nil, err
	}

	// 构造请求
	bs, _ := json.Marshal(r)
	req, err := http.NewRequest("POST", feishu.MessageSendUrl, bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+i.Token)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// 解码响应
	content, _ := io.ReadAll(resp.Body)

	// 1. 先解码外层
	ins := &feishu.SendMessageResponse{}
	if err := json.Unmarshal(content, ins); err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *Impl) UrgentCall(ctx context.Context, r *feishu.UrgentCallRequest) (*feishu.UrgentCallResponse, error) {
	if err := validator.New().Struct(r); err != nil {
		return nil, err
	}

	// 构造请求
	bs, _ := json.Marshal(r)

	url := fmt.Sprintf(feishu.UrgentCallUrl, r.MessageId)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+i.Token)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// 解码响应
	content, _ := io.ReadAll(resp.Body)

	// 1. 先解码外层
	ins := &feishu.UrgentCallResponse{}
	if err := json.Unmarshal(content, ins); err != nil {
		return nil, err
	}
	return ins, nil
}
