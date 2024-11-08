package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/feishu"
	"gitee.com/qiaogy91/K8sGenie/common"
	"github.com/go-playground/validator"
	"io"
	"net/http"
	"time"
)

func (i *Impl) CreateTable(ctx context.Context) error {
	return i.db.WithContext(ctx).AutoMigrate(&feishu.UrgentAuditRequest{})
}

// SetAccessToken 获取一次Token
func (i *Impl) setAccessToken(ctx context.Context) error {
	// 构造请求
	r := &feishu.GetAccessTokenRequest{
		AppId:     i.AppID,
		AppSecret: i.AppSecret,
	}
	bs, _ := json.Marshal(r)
	req, err := http.NewRequest("POST", feishu.AccessTokenUrl, bytes.NewBuffer(bs))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = rsp.Body.Close() }()

	// 解码响应
	ins := &feishu.GetAccessTokenResponse{}
	content, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(content, ins); err != nil {
		return err
	}

	if ins.Code == 0 {
		i.Token = ins.TenantAccessToken
	}
	return nil
}

func (i *Impl) fetchUserId(ctx context.Context, r *feishu.FetchUserIdRequest) (*feishu.FetchUserIdResponse, error) {
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

	if ins.Code != 0 {
		return nil, errors.New(ins.Msg)
	}
	return ins, nil
}

func (i *Impl) sendMessage(ctx context.Context, r *feishu.SendMessageRequest) (*feishu.SendMessageResponse, error) {
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

	if ins.Code != 0 {
		return nil, errors.New(ins.Msg)
	}
	return ins, nil
}

func (i *Impl) urgentMark(ctx context.Context, r *feishu.UrgentMarkRequest) (*feishu.UrgentMarkResponse, error) {
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
	ins := &feishu.UrgentMarkResponse{}
	if err := json.Unmarshal(content, ins); err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *Impl) CallAndAudit(ctx context.Context, req *feishu.CallAndAuditRequest) error {
	// 获取一个Token
	if err := i.setAccessToken(ctx); err != nil {
		return err
	}

	for user, phone := range req.Users {
		// 获取用户在应用内的 openID
		userRsp, err := i.fetchUserId(ctx, &feishu.FetchUserIdRequest{
			Mobiles: []string{phone}},
		)
		if err != nil {
			return err
		}

		//fmt.Printf("@@@@ %+v\n", userRsp)
		// 发送信息
		msgRsp, err := i.sendMessage(ctx, &feishu.SendMessageRequest{
			ReceiveID: userRsp.Data.UserList[0].UserID,
			Content:   req.Content,
			MsgType:   "text",
		})
		if err != nil {
			return err
		}

		// 信息加急
		callRsp, err := i.urgentMark(ctx, &feishu.UrgentMarkRequest{
			UserIdList: []string{userRsp.Data.UserList[0].UserID},
			MessageId:  msgRsp.Data.MessageId,
		})
		if err != nil {
			return err
		}

		// 记录日志
		ins := &feishu.UrgentAuditRequest{
			Meta: &feishu.Meta{CreatedTime: time.Now().Unix()},
			Spec: &feishu.Spec{
				User:    map[string]string{user: phone},
				Content: req.Content,
			},
			Status: &feishu.Status{
				Code: callRsp.Code,
				Msg:  callRsp.Msg,
			},
		}

		if err := i.db.WithContext(ctx).Model(&feishu.UrgentAuditRequest{}).Create(ins).Error; err != nil {
			return err
		}

		common.L().Info().Msgf("发送了一条加急消息: reciver(%s|%s) status(%d)", user, phone, callRsp.Code)
	}
	return nil
}
