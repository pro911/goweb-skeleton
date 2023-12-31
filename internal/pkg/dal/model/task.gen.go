// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTask = "task"

// Task mapped from table <task>
type Task struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:任务id" json:"id"`
	TaskID        string `gorm:"column:task_id;not null;comment:任务id" json:"task_id"`
	Name          string `gorm:"column:name;not null;comment:任务名称" json:"name"`
	ProjectID     int64  `gorm:"column:project_id;not null;comment:项目id" json:"project_id"`
	Vid           int64  `gorm:"column:vid;not null;comment:资产id" json:"vid"`
	TriggerEnable int64  `gorm:"column:trigger_enable;not null;default:1;comment:触发开关 1开 -1关闭" json:"trigger_enable"`
	TriggerUnit   int64  `gorm:"column:trigger_unit;not null;default:1;comment:触发单位: 1分 2时 3日 4周(月去掉了)" json:"trigger_unit"`
	TriggerAmount int64  `gorm:"column:trigger_amount;not null;comment:触发数值" json:"trigger_amount"`
	ServiceNode   int64  `gorm:"column:service_node;not null;comment:服务节点" json:"service_node"`
	OriginAt      int64  `gorm:"column:origin_at;not null;comment:起始时间" json:"origin_at"`
	NextAt        int64  `gorm:"column:next_at;not null;comment:下一次执行时间" json:"next_at"`
	Status        int64  `gorm:"column:status;not null;default:1;comment:状态: 1正常 -1删除" json:"status"`
	CreatedAt     int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`
	CreatedID     int64  `gorm:"column:created_id;not null;comment:创建人" json:"created_id"`
	UpdatedAt     int64  `gorm:"column:updated_at;not null;comment:修改时间" json:"updated_at"`
	UpdatedID     int64  `gorm:"column:updated_id;not null;comment:修改人" json:"updated_id"`
}

// TableName Task's table name
func (*Task) TableName() string {
	return TableNameTask
}
