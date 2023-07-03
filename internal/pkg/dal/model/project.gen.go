// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProject = "project"

// Project mapped from table <project>
type Project struct {
	ProjectID      int64  `gorm:"column:project_id;primaryKey;autoIncrement:true;comment:ID" json:"project_id"`
	TeamID         int64  `gorm:"column:team_id;not null;comment:团队ID" json:"team_id"`
	UserID         int64  `gorm:"column:user_id;comment:用户ID(6版本及以上不要使用此字段去判断管理员，转为创建人)" json:"user_id"`
	LocalProjectID string `gorm:"column:local_project_id;not null;comment:前端localProjectID" json:"local_project_id"`
	IsDefault      int64  `gorm:"column:is_default;not null;comment:是否为默认团队 0-非默认团队 1-默认团队" json:"is_default"`
	Name           string `gorm:"column:name;comment:项目名称" json:"name"`
	ImgURL         string `gorm:"column:imgUrl;not null;comment:项目封面图" json:"imgUrl"`
	CreateDtime    int64  `gorm:"column:create_dtime;comment:创建时间" json:"create_dtime"`
	Status         int64  `gorm:"column:status;not null;default:1;comment:状态 1-正常 -1已删除" json:"status"`
	IsLock         int64  `gorm:"column:is_lock;default:-1;comment:是否锁定" json:"is_lock"`
	Intro          string `gorm:"column:intro;comment:项目介绍" json:"intro"`
	UpdatedAt      int64  `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt      int64  `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`
	DeletedID      int64  `gorm:"column:deleted_id;comment:删除人用户ID" json:"deleted_id"`
	Visibility     int64  `gorm:"column:visibility;not null;default:1;comment:可见性 1私有 2公开" json:"visibility"`
	IsOffline      int64  `gorm:"column:is_offline;not null;default:-1;comment:是否离线项目。-1非离线项目，1离线项目" json:"is_offline"`
	Logo           string `gorm:"column:logo;not null;comment:项目分享logo" json:"logo"`
	Category       string `gorm:"column:category;not null;comment:类别" json:"category"`
}

// TableName Project's table name
func (*Project) TableName() string {
	return TableNameProject
}
