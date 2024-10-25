package cron

import (
	"context"
	"time"
)

const (
	AppName = "cron"
)

type QueryRequest struct {
	ClusterName string `json:"clusterName"` // 统计哪个集群
	Month       string `json:"month"`       // 统计哪个月份
}

type Service interface {
	CreateTable(ctx context.Context) error
	CreateNamespaceRecord(ctx context.Context, req *CreateNamespaceRecordRequest) error // 创建Record
	CreateProjectRecord(ctx context.Context, req *CreateProjectRecordRequest) error     // 创建Record
	CreateLineRecord(ctx context.Context, req *CreateLineRecordRequest) error           // 创建Record
	Run()
	// Query 需要写3个查询方法、查询 namespace、project、line 三个级别的汇总数据
	Query(ctx context.Context, req *QueryRequest)
}

type CreateNamespaceRecordRequest struct{}
type CreateProjectRecordRequest struct {
	MonthTime   string `json:"monthTime"`
	ClusterName string `json:"cluster_name"`
}

func (r *CreateProjectRecordRequest) TimeRage() (start int64, end int64, monthStr string, err error) {
	now := time.Now()
	month := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	if r.MonthTime != "" {
		month, err = time.Parse("2006-01", r.MonthTime)
		return 0, 0, "", err

	}
	start = month.Unix()
	end = month.AddDate(0, 1, 0).Unix() - 1
	monthStr = month.Format("2006-01")
	return
}

type CreateLineRecordRequest struct {
	MonthTime   string `json:"monthTime"`
	ClusterName string `json:"cluster_name"`
}

func (r *CreateLineRecordRequest) TimeRage() (start int64, end int64, monthStr string, err error) {
	now := time.Now()
	month := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	if r.MonthTime != "" {
		month, err = time.Parse("2006-01", r.MonthTime)
		return 0, 0, "", err

	}
	start = month.Unix()
	end = month.AddDate(0, 1, 0).Unix() - 1
	monthStr = month.Format("2006-01")
	return
}
