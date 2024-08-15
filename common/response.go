package common

import "github.com/emicklei/go-restful/v3"

type Response struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func SendFailed(rsp *restful.Response, code int, err error) {
	ins := Response{
		Code:  code,
		Error: err.Error(),
		Data:  nil,
	}
	if err := rsp.WriteHeaderAndEntity(code, ins); err != nil {
		L().Error().Msgf("send failed response failed: %v", err)
	}
}

func SendSuccess(rsp *restful.Response, data interface{}) {
	ins := Response{
		Code:  0,
		Error: "",
		Data:  data,
	}
	if err := rsp.WriteHeaderAndEntity(200, ins); err != nil {
		L().Error().Msgf("send success response failed: %v", err)
	}
}
