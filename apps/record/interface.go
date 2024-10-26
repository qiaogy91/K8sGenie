package record

import (
	"context"
	"time"
)

const AppName = "record"

type Service interface {
	CreateRecordService
	QueryRecordService
}

/*
------------------------- CreateRecordService ------------------------------------------------------
*/

type CreateRecordService interface {
	CreateTable(ctx context.Context) error
	CreateNamespaceRecord(ctx context.Context, req *CreateNamespaceRecordRequest) error // 创建Record
	CreateProjectRecord(ctx context.Context, req *CreateProjectRecordRequest) error     // 创建Record
	CreateLineRecord(ctx context.Context, req *CreateLineRecordRequest) error           // 创建Record
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

/*
------------------------- QueryRecordService ------------------------------------------------------
*/

type QueryRecordService interface {
	QueryNamespaceRecord(ctx context.Context, req *QueryNamespaceRecordRequest) (*NamespaceRecordSet, error)
	QueryProjectRecord(ctx context.Context, req *QueryProjectRecordRequest) (*ProjectRecordSet, error)
	QueryLineRecord(ctx context.Context, req *QueryLineRecordRequest) (*LineRecordSet, error)
}

type QueryNamespaceRecordRequest struct {
	CreatedTime string `json:"CreatedTime" validator:"required"` // 2014-08-02
	ProjectCode string `json:"projectCode" validator:"required"` // ehs
	ProjectLine string `json:"projectLine" validator:"required"` // 风控
	ClusterName string `json:"clusterName" validator:"required"` // xc-k8s-uat
}
type QueryProjectRecordRequest struct {
	Month       string `json:"month" validator:"required"` // 2014-08
	ProjectLine string `json:"projectLine" validator:"required"`
	ClusterName string `json:"clusterName" validator:"required"` // xc-k8s-uat

}
type QueryLineRecordRequest struct {
	Month       string `json:"month" validator:"required"`       // 2014-08
	ClusterName string `json:"clusterName" validator:"required"` // xc-k8s-uat
}
