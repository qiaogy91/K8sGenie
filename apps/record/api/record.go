package api

import (
	"gitee.com/qiaogy91/K8sGenie/apps/record"
	"gitee.com/qiaogy91/K8sGenie/ioc"
)

type Handler struct {
	svc record.Service
}

func (h *Handler) Name() string {
	return record.AppName
}

func (h *Handler) Init() error {
	h.svc = ioc.GetController(record.AppName).(record.Service)
	return nil
}

func (h *Handler) RegistryRoute(p string) {
	//TODO implement me
	panic("implement me")
}

func init() {
	ioc.RegistryHandlerObject(&Handler{})
}
