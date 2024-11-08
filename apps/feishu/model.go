package feishu

type Meta struct {
	ID          int64 `gorm:"primaryKey"`
	CreatedTime int64 `json:"createdTime" gorm:"autoCreateTime"`
}

type Spec struct {
	User    map[string]string `json:"user" gorm:"json"`
	Content string            `json:"content"`
}
type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type UrgentAuditRequest struct {
	Meta   *Meta   `gorm:"embedded"`
	Spec   *Spec   `gorm:"embedded"`
	Status *Status `gorm:"embedded"`
}
