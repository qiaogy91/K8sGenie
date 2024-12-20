package impl

import (
	"context"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/k8s"
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
	"gitee.com/qiaogy91/K8sGenie/apps/router"
	"github.com/go-playground/validator"
	_ "github.com/go-playground/validator"
	"strconv"
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

	// namespace 标签不存在，直接根据ClusterName 路由
	if req.NamespaceName == "" {
		if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("identity = ?", req.ClusterName).First(ins).Error; err != nil {
			return nil, err
		} else {
			return ins, nil
		}
	}

	// namespace 标签存在，尝试则查找对应的项目
	pro, err := i.kc.DescNamespace(ctx, &k8s.DescNamespaceReq{ClusterName: req.ClusterName, NamespaceName: req.NamespaceName})

	// 找不到Project，则进入默认路由
	if err != nil {
		// 默认路由
		if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("identity = ?", req.ClusterName).First(ins).Error; err != nil {
			return nil, err
		} else {
			return ins, nil
		}
	}
	// 找到了Project，尝试查找Project 对应的路由，如果找不到，则走默认路由
	if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("identity = ?", pro.Spec.ProjectId).First(ins).Error; err != nil {
		// 默认路由
		if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("identity = ?", req.ClusterName).First(ins).Error; err != nil {
			return nil, err
		} else {
			return ins, nil
		}
	}
	return ins, nil
}

func (i *Impl) QueryRoute(ctx context.Context, req *router.QueryRouteReq) (*router.RouterSet, error) {

	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	sql := i.db.WithContext(ctx).Model(&router.Router{})
	ins := &router.RouterSet{}

	switch req.Type {
	case router.QUERY_TYPE_QUERY_TYPE_BY_CLUSTER_CODE:

		sql = sql.Where("identity = ?", req.Keyword)

	case router.QUERY_TYPE_QUERY_TYPE_BY_ID:
		sql = sql.Where("id = ?", req.Keyword)

	case router.QUERY_TYPE_QUERY_TYPE_BY_PROJECT_LINE:
		// 查询这个产线下的所有 project
		ps, err := i.rc.QueryProject(ctx, &rancher.QueryProjectReq{QueryType: rancher.QUERY_TYPE_QUERY_TYPE_PROJECT_LINE, KeyWord: req.Keyword})
		if err != nil {
			return nil, err
		}

		sql = sql.Where("identity in ?", ps.ProjectIds())
	case router.QUERY_TYPE_QUERY_TYPE_ALL:
		sql = sql.Where("1=1")
	}

	if err := sql.Find(&ins.Items).Error; err != nil {
		return nil, err
	}

	// 附加Project 信息
	for _, r := range ins.Items {
		pro, err := i.rc.DescProject(ctx, &rancher.DescProjectReq{DescType: rancher.DESC_TYPE_DESC_TYPE_PROJECT_ID, KeyWord: r.Spec.Identity})
		if err != nil {
			continue
		}
		r.Spec.ProjectSpec = pro.Spec
	}

	return ins, nil
}

func (i *Impl) UpdateRoute(ctx context.Context, req *router.UpdateRouteReq) (*router.Router, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}
	// 更新当前用户
	id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, err
	}

	ins := &router.Router{
		Meta: &router.Meta{
			Id: id,
		},
		Spec: req.Spec,
	}
	fmt.Printf("@@@@ %+v\n", req.Spec)

	if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("id = ?", req.Id).Select("*").Updates(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *Impl) UrgentChange(ctx context.Context, req *router.UrgentChangeReq) (*router.Router, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("id = ?", req.Id).
		Update("urgent_call", req.UrgentCall).Error; err != nil {
		return nil, err
	}

	ins := &router.Router{}
	if err := i.db.WithContext(ctx).Model(&router.Router{}).Where("id = ?", req.Id).First(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}
