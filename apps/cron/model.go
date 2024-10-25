package cron

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

type ClustersSet []*Cluster
type Cluster struct {
	ClusterName string `json:"cluster_name"`
}
type GetProjectReq struct {
	ClusterName string `json:"cluster_name"`
}
type NamespaceWightSet []*NamespaceWight
type NamespaceWight struct {
	Namespace string `json:"namespace"`
	Weight    int    `json:"weight"`
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
	ClusterName string `json:"cluster_name"`
	ProjectLine string `json:"project_line"`
	ProjectDesc string `json:"project_desc"`
	ProjectCode string `json:"project_code"`
	Weight      int    `json:"weight"`
}

/*
------------------------- ProjectRecordSet------------------------------------------------------
*/

type LineRecordSet []*LineRecord
type LineRecord struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	CreatedAt   int64  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   int64  `json:"updatedAt" gorm:"autoUpdateTime"`
	ClusterName string `json:"cluster_name"`
	ProjectLine string `json:"project_line"`
	Month       string `json:"month"`
	Weight      int    `json:"weight"`
}
