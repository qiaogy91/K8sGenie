package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/cron"
	"gitee.com/qiaogy91/K8sGenie/ioc"
)

type Handler struct {
	svc cron.Service
}

func (h *Handler) Name() string {
	return cron.AppName
}

func (h *Handler) Init() error {
	h.svc = ioc.GetController(cron.AppName).(cron.Service)
	return nil
}

func (h *Handler) RegistryRoute(p string) {
	//TODO implement me
	panic("implement me")
}

func init() {
	ioc.RegistryHandlerObject(&Handler{})
}
