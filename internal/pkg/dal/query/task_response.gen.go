// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/pro911/golang-demo/internal/pkg/dal/model"
)

func newTaskResponse(db *gorm.DB, opts ...gen.DOOption) taskResponse {
	_taskResponse := taskResponse{}

	_taskResponse.taskResponseDo.UseDB(db, opts...)
	_taskResponse.taskResponseDo.UseModel(&model.TaskResponse{})

	tableName := _taskResponse.taskResponseDo.TableName()
	_taskResponse.ALL = field.NewAsterisk(tableName)
	_taskResponse.ID = field.NewInt64(tableName, "id")
	_taskResponse.Rid = field.NewString(tableName, "rid")
	_taskResponse.TaskID = field.NewString(tableName, "task_id")
	_taskResponse.ProjectID = field.NewInt64(tableName, "project_id")
	_taskResponse.Vid = field.NewInt64(tableName, "vid")
	_taskResponse.ExecAt = field.NewInt64(tableName, "exec_at")
	_taskResponse.ExecType = field.NewInt64(tableName, "exec_type")
	_taskResponse.TaskStatus = field.NewInt64(tableName, "task_status")
	_taskResponse.Report = field.NewString(tableName, "report")
	_taskResponse.CreatedAt = field.NewInt64(tableName, "created_at")
	_taskResponse.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_taskResponse.fillFieldMap()

	return _taskResponse
}

type taskResponse struct {
	taskResponseDo

	ALL        field.Asterisk
	ID         field.Int64  // 默认主键
	Rid        field.String // 日志uuidv4主键
	TaskID     field.String // 任务id主键
	ProjectID  field.Int64  // 项目id
	Vid        field.Int64  // 项目版本主键
	ExecAt     field.Int64  // 执行时间
	ExecType   field.Int64  // 执行类型:-1不执行 0定时执行 1立即执行
	TaskStatus field.Int64  // 任务状态: 1成功 -1失败
	Report     field.String
	CreatedAt  field.Int64 // 创建时间
	UpdatedAt  field.Int64 // 更新时间

	fieldMap map[string]field.Expr
}

func (t taskResponse) Table(newTableName string) *taskResponse {
	t.taskResponseDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t taskResponse) As(alias string) *taskResponse {
	t.taskResponseDo.DO = *(t.taskResponseDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *taskResponse) updateTableName(table string) *taskResponse {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.Rid = field.NewString(table, "rid")
	t.TaskID = field.NewString(table, "task_id")
	t.ProjectID = field.NewInt64(table, "project_id")
	t.Vid = field.NewInt64(table, "vid")
	t.ExecAt = field.NewInt64(table, "exec_at")
	t.ExecType = field.NewInt64(table, "exec_type")
	t.TaskStatus = field.NewInt64(table, "task_status")
	t.Report = field.NewString(table, "report")
	t.CreatedAt = field.NewInt64(table, "created_at")
	t.UpdatedAt = field.NewInt64(table, "updated_at")

	t.fillFieldMap()

	return t
}

func (t *taskResponse) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *taskResponse) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 11)
	t.fieldMap["id"] = t.ID
	t.fieldMap["rid"] = t.Rid
	t.fieldMap["task_id"] = t.TaskID
	t.fieldMap["project_id"] = t.ProjectID
	t.fieldMap["vid"] = t.Vid
	t.fieldMap["exec_at"] = t.ExecAt
	t.fieldMap["exec_type"] = t.ExecType
	t.fieldMap["task_status"] = t.TaskStatus
	t.fieldMap["report"] = t.Report
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
}

func (t taskResponse) clone(db *gorm.DB) taskResponse {
	t.taskResponseDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t taskResponse) replaceDB(db *gorm.DB) taskResponse {
	t.taskResponseDo.ReplaceDB(db)
	return t
}

type taskResponseDo struct{ gen.DO }

type ITaskResponseDo interface {
	gen.SubQuery
	Debug() ITaskResponseDo
	WithContext(ctx context.Context) ITaskResponseDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITaskResponseDo
	WriteDB() ITaskResponseDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITaskResponseDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITaskResponseDo
	Not(conds ...gen.Condition) ITaskResponseDo
	Or(conds ...gen.Condition) ITaskResponseDo
	Select(conds ...field.Expr) ITaskResponseDo
	Where(conds ...gen.Condition) ITaskResponseDo
	Order(conds ...field.Expr) ITaskResponseDo
	Distinct(cols ...field.Expr) ITaskResponseDo
	Omit(cols ...field.Expr) ITaskResponseDo
	Join(table schema.Tabler, on ...field.Expr) ITaskResponseDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITaskResponseDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITaskResponseDo
	Group(cols ...field.Expr) ITaskResponseDo
	Having(conds ...gen.Condition) ITaskResponseDo
	Limit(limit int) ITaskResponseDo
	Offset(offset int) ITaskResponseDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITaskResponseDo
	Unscoped() ITaskResponseDo
	Create(values ...*model.TaskResponse) error
	CreateInBatches(values []*model.TaskResponse, batchSize int) error
	Save(values ...*model.TaskResponse) error
	First() (*model.TaskResponse, error)
	Take() (*model.TaskResponse, error)
	Last() (*model.TaskResponse, error)
	Find() ([]*model.TaskResponse, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TaskResponse, err error)
	FindInBatches(result *[]*model.TaskResponse, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TaskResponse) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITaskResponseDo
	Assign(attrs ...field.AssignExpr) ITaskResponseDo
	Joins(fields ...field.RelationField) ITaskResponseDo
	Preload(fields ...field.RelationField) ITaskResponseDo
	FirstOrInit() (*model.TaskResponse, error)
	FirstOrCreate() (*model.TaskResponse, error)
	FindByPage(offset int, limit int) (result []*model.TaskResponse, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITaskResponseDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t taskResponseDo) Debug() ITaskResponseDo {
	return t.withDO(t.DO.Debug())
}

func (t taskResponseDo) WithContext(ctx context.Context) ITaskResponseDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t taskResponseDo) ReadDB() ITaskResponseDo {
	return t.Clauses(dbresolver.Read)
}

func (t taskResponseDo) WriteDB() ITaskResponseDo {
	return t.Clauses(dbresolver.Write)
}

func (t taskResponseDo) Session(config *gorm.Session) ITaskResponseDo {
	return t.withDO(t.DO.Session(config))
}

func (t taskResponseDo) Clauses(conds ...clause.Expression) ITaskResponseDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t taskResponseDo) Returning(value interface{}, columns ...string) ITaskResponseDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t taskResponseDo) Not(conds ...gen.Condition) ITaskResponseDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t taskResponseDo) Or(conds ...gen.Condition) ITaskResponseDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t taskResponseDo) Select(conds ...field.Expr) ITaskResponseDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t taskResponseDo) Where(conds ...gen.Condition) ITaskResponseDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t taskResponseDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITaskResponseDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t taskResponseDo) Order(conds ...field.Expr) ITaskResponseDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t taskResponseDo) Distinct(cols ...field.Expr) ITaskResponseDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t taskResponseDo) Omit(cols ...field.Expr) ITaskResponseDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t taskResponseDo) Join(table schema.Tabler, on ...field.Expr) ITaskResponseDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t taskResponseDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITaskResponseDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t taskResponseDo) RightJoin(table schema.Tabler, on ...field.Expr) ITaskResponseDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t taskResponseDo) Group(cols ...field.Expr) ITaskResponseDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t taskResponseDo) Having(conds ...gen.Condition) ITaskResponseDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t taskResponseDo) Limit(limit int) ITaskResponseDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t taskResponseDo) Offset(offset int) ITaskResponseDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t taskResponseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITaskResponseDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t taskResponseDo) Unscoped() ITaskResponseDo {
	return t.withDO(t.DO.Unscoped())
}

func (t taskResponseDo) Create(values ...*model.TaskResponse) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t taskResponseDo) CreateInBatches(values []*model.TaskResponse, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t taskResponseDo) Save(values ...*model.TaskResponse) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t taskResponseDo) First() (*model.TaskResponse, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskResponse), nil
	}
}

func (t taskResponseDo) Take() (*model.TaskResponse, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskResponse), nil
	}
}

func (t taskResponseDo) Last() (*model.TaskResponse, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskResponse), nil
	}
}

func (t taskResponseDo) Find() ([]*model.TaskResponse, error) {
	result, err := t.DO.Find()
	return result.([]*model.TaskResponse), err
}

func (t taskResponseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TaskResponse, err error) {
	buf := make([]*model.TaskResponse, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t taskResponseDo) FindInBatches(result *[]*model.TaskResponse, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t taskResponseDo) Attrs(attrs ...field.AssignExpr) ITaskResponseDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t taskResponseDo) Assign(attrs ...field.AssignExpr) ITaskResponseDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t taskResponseDo) Joins(fields ...field.RelationField) ITaskResponseDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t taskResponseDo) Preload(fields ...field.RelationField) ITaskResponseDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t taskResponseDo) FirstOrInit() (*model.TaskResponse, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskResponse), nil
	}
}

func (t taskResponseDo) FirstOrCreate() (*model.TaskResponse, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TaskResponse), nil
	}
}

func (t taskResponseDo) FindByPage(offset int, limit int) (result []*model.TaskResponse, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t taskResponseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t taskResponseDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t taskResponseDo) Delete(models ...*model.TaskResponse) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *taskResponseDo) withDO(do gen.Dao) *taskResponseDo {
	t.DO = *do.(*gen.DO)
	return t
}