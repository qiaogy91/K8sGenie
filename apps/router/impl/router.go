package impl

import (
	"context"
	"gitee.com/qiaogy91/K8sGenie/apps/rancher"
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

	// 如果提供了projectID
	if spec.ProjectId != "" {
		// 判断是否有效ID
		if _, err := i.rc.QueryProject(ctx, &rancher.QueryProjectReq{SearchType: rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_ID, KeyWord: spec.ProjectId}); err != nil {
			return nil, err
		}
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

// QueryRoute 查询各项目告警路由
// 如果要查询 “智能财务” 产线的情况，那么内部会调用 getProjectIds 来获取 “智能财务产线” 所有的ProjectID，来构造sql
func (i *Impl) QueryRoute(ctx context.Context, req *router.QueryRouteReq) (*router.RouterSet, error) {
	sql := i.db.WithContext(ctx).Model(&router.Router{})
	ins := &router.RouterSet{}

	switch req.SearchType {
	case router.SEARCH_TYPE_SEARCH_TYPE_PROJECT_ID:
		sql = sql.Where("project_id = ?", req.KeyWord)
	case router.SEARCH_TYPE_SEARCH_TYPE_ROUTER_ID:
		sql = sql.Where("id = ?", req.KeyWord)
	case router.SEARCH_TYPE_SEARCH_TYPE_CLUSTER_NAME:
		sql = sql.Where("cluster_name like ?", "%"+req.KeyWord+"%")
	default:
		ids, err := i.getProjectIds(ctx, req)
		if err != nil {
			return nil, err
		}
		sql = sql.Where("project_id in ?", ids)
	}

	if err := sql.Find(&ins.Items).Error; err != nil {
		return nil, err
	}
	ins.Total = int64(len(ins.Items))

	return ins, nil
}

func (i *Impl) getProjectIds(ctx context.Context, req *router.QueryRouteReq) ([]string, error) {
	var (
		err  error
		ids  []string
		pros = &rancher.ProjectSet{}
	)
	switch req.SearchType {
	case router.SEARCH_TYPE_SEARCH_TYPE_PROJECT_CODE:
		pros, err = i.rc.QueryProject(ctx, &rancher.QueryProjectReq{SearchType: rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_CODE, KeyWord: req.KeyWord})

	case router.SEARCH_TYPE_SEARCH_TYPE_PROJECT_LINE:
		pros, err = i.rc.QueryProject(ctx, &rancher.QueryProjectReq{SearchType: rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_LINE, KeyWord: req.KeyWord})

	case router.SEARCH_TYPE_SEARCH_TYPE_PROJECT_DESC:
		pros, err = i.rc.QueryProject(ctx, &rancher.QueryProjectReq{SearchType: rancher.SEARCH_TYPE_SEARCH_TYPE_PROJECT_DESC, KeyWord: req.KeyWord})
	}

	if err != nil {
		return nil, err
	}

	for _, pro := range pros.Items {
		ids = append(ids, pro.Spec.ProjectId)
	}
	return ids, nil
}
