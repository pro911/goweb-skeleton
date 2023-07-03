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

func newTeamPlace(db *gorm.DB, opts ...gen.DOOption) teamPlace {
	_teamPlace := teamPlace{}

	_teamPlace.teamPlaceDo.UseDB(db, opts...)
	_teamPlace.teamPlaceDo.UseModel(&model.TeamPlace{})

	tableName := _teamPlace.teamPlaceDo.TableName()
	_teamPlace.ALL = field.NewAsterisk(tableName)
	_teamPlace.ID = field.NewInt64(tableName, "id")
	_teamPlace.TeamID = field.NewInt64(tableName, "team_id")
	_teamPlace.UserID = field.NewInt64(tableName, "user_id")
	_teamPlace.IsAdmin = field.NewInt64(tableName, "is_admin")
	_teamPlace.BindDtime = field.NewInt64(tableName, "bind_dtime")
	_teamPlace.IsPresent = field.NewInt64(tableName, "is_present")
	_teamPlace.Status = field.NewInt64(tableName, "status")
	_teamPlace.IsReadonly = field.NewInt64(tableName, "is_readonly")
	_teamPlace.IsSuperAdmin = field.NewInt64(tableName, "is_super_admin")
	_teamPlace.InviterUserID = field.NewInt64(tableName, "inviter_user_id")
	_teamPlace.ExpiryTime = field.NewInt64(tableName, "expiry_time")
	_teamPlace.CreatedAt = field.NewTime(tableName, "created_at")

	_teamPlace.fillFieldMap()

	return _teamPlace
}

type teamPlace struct {
	teamPlaceDo

	ALL           field.Asterisk
	ID            field.Int64 // ID
	TeamID        field.Int64 // 团队ID
	UserID        field.Int64 // 绑定的用户ID
	IsAdmin       field.Int64 // 是否为管理员 1-是
	BindDtime     field.Int64 // 团队成员绑定时间
	IsPresent     field.Int64 // 是否为付费工位 -1-是 1-否
	Status        field.Int64 // 状态 1-正常
	IsReadonly    field.Int64 // 工位读写权限 1-读写 2-只读
	IsSuperAdmin  field.Int64 // 是否为超级管理员，1是，-1否
	InviterUserID field.Int64 // 邀请人
	ExpiryTime    field.Int64 // 过期时间
	CreatedAt     field.Time  // 创建时间

	fieldMap map[string]field.Expr
}

func (t teamPlace) Table(newTableName string) *teamPlace {
	t.teamPlaceDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t teamPlace) As(alias string) *teamPlace {
	t.teamPlaceDo.DO = *(t.teamPlaceDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *teamPlace) updateTableName(table string) *teamPlace {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.TeamID = field.NewInt64(table, "team_id")
	t.UserID = field.NewInt64(table, "user_id")
	t.IsAdmin = field.NewInt64(table, "is_admin")
	t.BindDtime = field.NewInt64(table, "bind_dtime")
	t.IsPresent = field.NewInt64(table, "is_present")
	t.Status = field.NewInt64(table, "status")
	t.IsReadonly = field.NewInt64(table, "is_readonly")
	t.IsSuperAdmin = field.NewInt64(table, "is_super_admin")
	t.InviterUserID = field.NewInt64(table, "inviter_user_id")
	t.ExpiryTime = field.NewInt64(table, "expiry_time")
	t.CreatedAt = field.NewTime(table, "created_at")

	t.fillFieldMap()

	return t
}

func (t *teamPlace) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *teamPlace) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 12)
	t.fieldMap["id"] = t.ID
	t.fieldMap["team_id"] = t.TeamID
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["is_admin"] = t.IsAdmin
	t.fieldMap["bind_dtime"] = t.BindDtime
	t.fieldMap["is_present"] = t.IsPresent
	t.fieldMap["status"] = t.Status
	t.fieldMap["is_readonly"] = t.IsReadonly
	t.fieldMap["is_super_admin"] = t.IsSuperAdmin
	t.fieldMap["inviter_user_id"] = t.InviterUserID
	t.fieldMap["expiry_time"] = t.ExpiryTime
	t.fieldMap["created_at"] = t.CreatedAt
}

func (t teamPlace) clone(db *gorm.DB) teamPlace {
	t.teamPlaceDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t teamPlace) replaceDB(db *gorm.DB) teamPlace {
	t.teamPlaceDo.ReplaceDB(db)
	return t
}

type teamPlaceDo struct{ gen.DO }

type ITeamPlaceDo interface {
	gen.SubQuery
	Debug() ITeamPlaceDo
	WithContext(ctx context.Context) ITeamPlaceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITeamPlaceDo
	WriteDB() ITeamPlaceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITeamPlaceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITeamPlaceDo
	Not(conds ...gen.Condition) ITeamPlaceDo
	Or(conds ...gen.Condition) ITeamPlaceDo
	Select(conds ...field.Expr) ITeamPlaceDo
	Where(conds ...gen.Condition) ITeamPlaceDo
	Order(conds ...field.Expr) ITeamPlaceDo
	Distinct(cols ...field.Expr) ITeamPlaceDo
	Omit(cols ...field.Expr) ITeamPlaceDo
	Join(table schema.Tabler, on ...field.Expr) ITeamPlaceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITeamPlaceDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITeamPlaceDo
	Group(cols ...field.Expr) ITeamPlaceDo
	Having(conds ...gen.Condition) ITeamPlaceDo
	Limit(limit int) ITeamPlaceDo
	Offset(offset int) ITeamPlaceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITeamPlaceDo
	Unscoped() ITeamPlaceDo
	Create(values ...*model.TeamPlace) error
	CreateInBatches(values []*model.TeamPlace, batchSize int) error
	Save(values ...*model.TeamPlace) error
	First() (*model.TeamPlace, error)
	Take() (*model.TeamPlace, error)
	Last() (*model.TeamPlace, error)
	Find() ([]*model.TeamPlace, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TeamPlace, err error)
	FindInBatches(result *[]*model.TeamPlace, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TeamPlace) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITeamPlaceDo
	Assign(attrs ...field.AssignExpr) ITeamPlaceDo
	Joins(fields ...field.RelationField) ITeamPlaceDo
	Preload(fields ...field.RelationField) ITeamPlaceDo
	FirstOrInit() (*model.TeamPlace, error)
	FirstOrCreate() (*model.TeamPlace, error)
	FindByPage(offset int, limit int) (result []*model.TeamPlace, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITeamPlaceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t teamPlaceDo) Debug() ITeamPlaceDo {
	return t.withDO(t.DO.Debug())
}

func (t teamPlaceDo) WithContext(ctx context.Context) ITeamPlaceDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t teamPlaceDo) ReadDB() ITeamPlaceDo {
	return t.Clauses(dbresolver.Read)
}

func (t teamPlaceDo) WriteDB() ITeamPlaceDo {
	return t.Clauses(dbresolver.Write)
}

func (t teamPlaceDo) Session(config *gorm.Session) ITeamPlaceDo {
	return t.withDO(t.DO.Session(config))
}

func (t teamPlaceDo) Clauses(conds ...clause.Expression) ITeamPlaceDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t teamPlaceDo) Returning(value interface{}, columns ...string) ITeamPlaceDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t teamPlaceDo) Not(conds ...gen.Condition) ITeamPlaceDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t teamPlaceDo) Or(conds ...gen.Condition) ITeamPlaceDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t teamPlaceDo) Select(conds ...field.Expr) ITeamPlaceDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t teamPlaceDo) Where(conds ...gen.Condition) ITeamPlaceDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t teamPlaceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITeamPlaceDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t teamPlaceDo) Order(conds ...field.Expr) ITeamPlaceDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t teamPlaceDo) Distinct(cols ...field.Expr) ITeamPlaceDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t teamPlaceDo) Omit(cols ...field.Expr) ITeamPlaceDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t teamPlaceDo) Join(table schema.Tabler, on ...field.Expr) ITeamPlaceDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t teamPlaceDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITeamPlaceDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t teamPlaceDo) RightJoin(table schema.Tabler, on ...field.Expr) ITeamPlaceDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t teamPlaceDo) Group(cols ...field.Expr) ITeamPlaceDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t teamPlaceDo) Having(conds ...gen.Condition) ITeamPlaceDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t teamPlaceDo) Limit(limit int) ITeamPlaceDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t teamPlaceDo) Offset(offset int) ITeamPlaceDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t teamPlaceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITeamPlaceDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t teamPlaceDo) Unscoped() ITeamPlaceDo {
	return t.withDO(t.DO.Unscoped())
}

func (t teamPlaceDo) Create(values ...*model.TeamPlace) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t teamPlaceDo) CreateInBatches(values []*model.TeamPlace, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t teamPlaceDo) Save(values ...*model.TeamPlace) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t teamPlaceDo) First() (*model.TeamPlace, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamPlace), nil
	}
}

func (t teamPlaceDo) Take() (*model.TeamPlace, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamPlace), nil
	}
}

func (t teamPlaceDo) Last() (*model.TeamPlace, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamPlace), nil
	}
}

func (t teamPlaceDo) Find() ([]*model.TeamPlace, error) {
	result, err := t.DO.Find()
	return result.([]*model.TeamPlace), err
}

func (t teamPlaceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TeamPlace, err error) {
	buf := make([]*model.TeamPlace, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t teamPlaceDo) FindInBatches(result *[]*model.TeamPlace, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t teamPlaceDo) Attrs(attrs ...field.AssignExpr) ITeamPlaceDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t teamPlaceDo) Assign(attrs ...field.AssignExpr) ITeamPlaceDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t teamPlaceDo) Joins(fields ...field.RelationField) ITeamPlaceDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t teamPlaceDo) Preload(fields ...field.RelationField) ITeamPlaceDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t teamPlaceDo) FirstOrInit() (*model.TeamPlace, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamPlace), nil
	}
}

func (t teamPlaceDo) FirstOrCreate() (*model.TeamPlace, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TeamPlace), nil
	}
}

func (t teamPlaceDo) FindByPage(offset int, limit int) (result []*model.TeamPlace, count int64, err error) {
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

func (t teamPlaceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t teamPlaceDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t teamPlaceDo) Delete(models ...*model.TeamPlace) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *teamPlaceDo) withDO(do gen.Dao) *teamPlaceDo {
	t.DO = *do.(*gen.DO)
	return t
}