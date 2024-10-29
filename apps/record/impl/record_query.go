package impl

import (
	"context"
	"fmt"
	"gitee.com/qiaogy91/K8sGenie/apps/record"
	"github.com/go-playground/validator"
	"time"
)

// QueryNamespaceRecord 查询项目中各名称空间用量分布
func (i *Impl) QueryNamespaceRecord(ctx context.Context, req *record.QueryNamespaceRecordRequest) (*record.NamespaceRecordSet, error) {
	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	t, err := time.Parse("2006-01-02", req.CreatedAt)
	if err != nil {
		return nil, err
	}

	start := t.Unix()
	end := start + 24*60*60

	ins := &record.NamespaceRecordSet{}

	query := i.db.Debug().WithContext(ctx).Where("cluster_name = ? AND project_line = ? AND project_code = ?", req.ClusterName, req.ProjectLine, req.ProjectCode)
	if req.ProjectCode == "" {
		query = i.db.Debug().WithContext(ctx).Where("cluster_name = ? AND project_line = ?", req.ClusterName, req.ProjectLine)
	}

	sql := query.Where("created_at >= ? AND  created_at <= ?", start, end)

	if err := sql.Find(ins).Error; err != nil {
		return nil, err
	}

	// 获取百分比
	ins.GetPercent()
	return ins, nil
}

// QueryProjectRecord 查询各产线中项目用量分布
func (i *Impl) QueryProjectRecord(ctx context.Context, req *record.QueryProjectRecordRequest) (*record.ProjectRecordSet, error) {
	sql := i.db.WithContext(ctx).Model(&record.ProjectRecord{}).Debug()

	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	fmt.Printf("@@@@@@ %+v\n", req)

	query := sql.Where("month = ? AND cluster_name = ? And project_line = ?", req.Month, req.ClusterName, req.ProjectLine)
	if req.ProjectLine == "" {
		query = sql.Where("month = ? AND cluster_name = ?", req.Month, req.ClusterName)
	}

	ins := &record.ProjectRecordSet{}
	if err := query.Find(ins).Error; err != nil {
		return nil, err
	}

	// 获取百分比
	ins.GetPercent()
	return ins, nil
}

// QueryLineRecord 查询各集群中产线用量分布
func (i *Impl) QueryLineRecord(ctx context.Context, req *record.QueryLineRecordRequest) (*record.LineRecordSet, error) {
	sql := i.db.WithContext(ctx).Model(&record.LineRecord{}).Debug()

	if err := validator.New().Struct(req); err != nil {
		return nil, err
	}

	ins := &record.LineRecordSet{}
	if err := sql.Where("month = ? AND cluster_name = ?", req.Month, req.ClusterName).Find(ins).Error; err != nil {
		return nil, err
	}

	// 计算百分比
	ins.GetPercent()
	return ins, nil
}
