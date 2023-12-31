// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProjectUser = "project_user"

// ProjectUser mapped from table <project_user>
type ProjectUser struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`
	UserID      int64  `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	ProjectID   int64  `gorm:"column:project_id;comment:项目ID" json:"project_id"`
	Status      int64  `gorm:"column:status;not null;default:1;comment:状态 1-启用" json:"status"`
	Sort        int64  `gorm:"column:sort;not null;comment:排序" json:"sort"`
	CreatedAt   int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`
	EditedAt    int64  `gorm:"column:edited_at;not null;comment:此处编辑时间以apis切换行时间为准" json:"edited_at"`
	EnvID       int64  `gorm:"column:env_id;not null;comment:最近使用的项目环境" json:"env_id"`
	Collect     int64  `gorm:"column:collect;not null;comment:收藏 1已收藏 0未收藏" json:"collect"`
	CollectedAt int64  `gorm:"column:collected_at;not null;comment:收藏时间" json:"collected_at"`
	Role        int64  `gorm:"column:role;not null;default:1;comment:角色 1读 2写" json:"role"`
	IsManager   int64  `gorm:"column:is_manager;not null;default:-1;comment:是否为项目拥有者 -1不是 1是" json:"is_manager"`
	LocalEnvID  string `gorm:"column:local_env_id;not null;default:-1;comment:最近使用的项目环境" json:"local_env_id"`
}

// TableName ProjectUser's table name
func (*ProjectUser) TableName() string {
	return TableNameProjectUser
}
