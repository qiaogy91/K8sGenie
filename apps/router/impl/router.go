package impl

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/router"
	"github.com/go-playground/validator"
	_ "github.com/go-playground/validator"
)

func (i *Impl) CreateTable(ctx context.Context, empty *router.Empty) (*router.Empty, error) {
	return nil, i.db.WithContext(ctx).AutoMigrate(&router.Router{})
}

func (i *Impl) CreateRoute(ctx context.Context, spec *router.Spec) (*router.Router, error) {
	// 校验
	if err := validator.New().Struct(spec); err != nil {
		return nil, err
	}

	// 创建
	ins := &router.Router{Spec: spec}
	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *Impl) DeleteRoute(ctx context.Context, req *router.DeleteRouteReq) (*router.Router, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins := &router.Router{}
	if err := i.db.WithContext(ctx).Where("id = ?", req.Id).First(ins).Error; err != nil {
		return nil, err
	}

	if err := i.db.WithContext(ctx).Where("id = ?", req.Id).Delete(&router.Router{}).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *Impl) DescRoute(ctx context.Context, req *router.QueryRouteReq) (*router.Router, error) {
	ins := &router.Router{}

	if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("type = ? AND identity = ?", req.Type, req.Identity).First(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}
