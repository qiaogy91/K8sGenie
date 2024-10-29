package record

import "math"

/*
------------------------- NamespaceRecord------------------------------------------------------
*/

type NamespaceRecordSet []*NamespaceRecord
type NamespaceRecord struct {
	ID             int64           `json:"id" gorm:"primaryKey"`
	CreatedAt      int64           `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      int64           `json:"updatedAt" gorm:"autoUpdateTime"`
	ClusterName    string          `json:"cluster_name"`
	ProjectLine    string          `json:"project_line"`
	ProjectDesc    string          `json:"project_desc"`
	ProjectCode    string          `json:"project_code"`
	NamespaceWight *NamespaceWight `json:"namespace_wight" gorm:"embedded"`
}

func (s *NamespaceRecordSet) GetPercent() {
	var total float64
	for _, item := range *s {
		total += float64(item.NamespaceWight.Weight)
	}
	for _, item := range *s {
		item.NamespaceWight.Percent = math.Round(float64(item.NamespaceWight.Weight)/total*100) / 100
	}
}

type ClustersSet []*Cluster
type Cluster struct {
	ClusterName string `json:"cluster_name"`
}
type GetProjectReq struct {
	ClusterName string `json:"cluster_name"`
}
type NamespaceWightSet []*NamespaceWight
type NamespaceWight struct {
	Namespace string  `json:"namespace"`
	Weight    int     `json:"weight"`
	Percent   float64 `json:"percent"`
}

/*
------------------------- ProjectRecordSet------------------------------------------------------
*/

type ProjectRecordSet []*ProjectRecord
type ProjectRecord struct {
	ID           int64         `json:"id" gorm:"primaryKey"`
	CreatedAt    int64         `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    int64         `json:"updatedAt" gorm:"autoUpdateTime"`
	Month        string        `json:"month"`
	ProjectWight *ProjectWight `json:"project_wight" gorm:"embedded"`
}
type ProjectWight struct {
	ClusterName string  `json:"cluster_name"`
	ProjectLine string  `json:"project_line"`
	ProjectDesc string  `json:"project_desc"`
	ProjectCode string  `json:"project_code"`
	Weight      int     `json:"weight"`
	Percent     float64 `json:"percent"`
}

func (s *ProjectRecordSet) GetPercent() {
	var total float64
	for _, item := range *s {
		total += float64(item.ProjectWight.Weight)
	}
	for _, item := range *s {
		item.ProjectWight.Percent = math.Round(float64(item.ProjectWight.Weight)/total*100) / 100
	}
}

/*
------------------------- ProjectRecordSet------------------------------------------------------
*/

type LineRecordSet []*LineRecord
type LineRecord struct {
	ID          int64   `json:"id" gorm:"primaryKey"`
	CreatedAt   int64   `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   int64   `json:"updatedAt" gorm:"autoUpdateTime"`
	ClusterName string  `json:"cluster_name"`
	ProjectLine string  `json:"project_line"`
	Month       string  `json:"month"`
	Weight      int     `json:"weight"`
	Percent     float64 `json:"percent"`
}

func (s *LineRecordSet) GetPercent() {
	var total float64
	for _, item := range *s {
		total += float64(item.Weight)
	}
	for _, item := range *s {
		item.Percent = math.Round(float64(item.Weight)/total*100) / 100
	}
}
