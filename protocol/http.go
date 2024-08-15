package protocol

import (
	"context"
	"errors"
	"gitee.com/qiaogy91/K8sGenie/common"
	"gitee.com/qiaogy91/K8sGenie/conf"
	"gitee.com/qiaogy91/K8sGenie/ioc"
	"net/http"
	"time"
)

type HttpServer struct {
	s *http.Server
}

func (h *HttpServer) Start() error {
	// 路由注册
	for name, o := range ioc.LoadHandler() {
		o.RegistryRoute(common.ProjectName)
		common.L().Info().Msgf("%s handler object's route registry success", name)
	}

	// 辅助信息：打印所有已经注册的URL
	for _, ws := range restful.RegisteredWebServices() {
		for _, r := range ws.Routes() {
			common.L().Info().Msgf("%s %s", r.Method, r.Path)
		}
	}

	// 启动监听
	common.L().Info().Msgf("http server started, at %s", conf.C().HttpAddr())
	if err := h.s.ListenAndServe(); err != nil {
		// 优雅关闭
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}

func (h *HttpServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.s.Shutdown(ctx)
	common.L().Info().Msgf("http server stopped: %v", err)
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		s: &http.Server{
			Addr:              conf.C().HttpAddr(),
			Handler:           restful.DefaultContainer,
			ReadTimeout:       time.Duration(conf.C().Http.ReadTimeout) * time.Second,
			ReadHeaderTimeout: 0,
			WriteTimeout:      time.Duration(conf.C().Http.WriteTimeout) * time.Second,
		},
	}
}
