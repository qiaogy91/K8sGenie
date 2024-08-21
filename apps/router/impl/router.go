package impl

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
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

func (i *Impl) AlertRoute(ctx context.Context, req *router.AlertRouteReq) (*router.Router, error) {
	ins := &router.Router{}

	// namespace 标签不存在，直接根据robot_name 路由
	if req.NamespaceName == "" {
		if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("identity = ?", req.RobotName).First(ins).Error; err != nil {
			return nil, err
		} else {
			return ins, nil
		}
	}

	// namespace 标签存在，尝试则查找对应的项目
	pro, err := i.kc.DescNamespace(ctx, &k8s.DescNamespaceReq{ClusterName: req.RobotName, NamespaceName: req.NamespaceName})

	// 找不到Project，则进入默认路由
	if err != nil {
		// 默认路由
		if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("identity = ?", req.RobotName).First(ins).Error; err != nil {
			return nil, err
		} else {
			return ins, nil
		}
	}
	// 找到了Project，尝试查找Project 对应的路由，如果找不到，则走默认路由
	if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("identity = ?", pro.Spec.ProjectId).First(ins).Error; err != nil {
		// 默认路由
		if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("identity = ?", req.RobotName).First(ins).Error; err != nil {
			return nil, err
		} else {
			return ins, nil
		}
	}
	return ins, nil
}
