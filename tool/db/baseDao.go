package db

import (
	"fmt"
	"reflect"
	"time"

	"github.com/shixinshuiyou/mayo/tool/gconv"
	"github.com/shixinshuiyou/mayo/tool/util"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
)

// TODO 可以引入jaeger 进行链路追踪
// 引入jaeger 需要了解 gorm 上下文 context

type BaseDao struct {
	CreateID   int       `gorm:"column:create_id;default:0" json:"createID"`                     // 创建人
	OperateID  int       `gorm:"column:operate_id;default:0" json:"operateID"`                   // 操作、更新人
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"updateTime"` // 更新时间
}

func (dao *BaseDao) BeforeCreate(tx *gorm.DB) (err error) {
	if dao.CreateTime.IsZero() {
		dao.CreateTime = time.Now()
	}
	if dao.UpdateTime.IsZero() {
		dao.UpdateTime = time.Now()
	}
	return
}

func (dao *BaseDao) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (dao *BaseDao) BeforeUpdate(tx *gorm.DB) (err error) {
	// 有实际值改变，才会改变
	changes := change(tx)
	if len(changes) > 0 {
		dao.UpdateTime = time.Now()
		tx.Statement.SetColumn("UpdateTime", time.Now())
	}

	return
}

func (dao *BaseDao) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

func change(tx *gorm.DB) (result map[string]string) {
	stmt := tx.Statement
	modelValue := stmt.ReflectValue
	switch modelValue.Kind() {
	case reflect.Slice, reflect.Array:
		modelValue = stmt.ReflectValue.Index(stmt.CurDestIndex)
	}
	selectColumns, restricted := stmt.SelectAndOmitColumns(false, true)
	changed := func(field *schema.Field) (bool, interface{}, interface{}) {
		fieldValue, _ := field.ValueOf(tx.Statement.Context, modelValue)
		if v, ok := selectColumns[field.DBName]; (ok && v) || (!ok && !restricted) {
			if v, ok := stmt.Dest.(map[string]interface{}); ok {
				if fv, ok := v[field.Name]; ok {
					return !utils.AssertEqual(fv, fieldValue), fv, fieldValue
				} else if fv, ok := v[field.DBName]; ok {
					return !utils.AssertEqual(fv, fieldValue), fv, fieldValue
				}
			} else {
				destValue := reflect.ValueOf(stmt.Dest)
				for destValue.Kind() == reflect.Ptr {
					destValue = destValue.Elem()
				}
				changedValue, zero := field.ValueOf(tx.Statement.Context, destValue)
				return !zero && !utils.AssertEqual(changedValue, fieldValue), changedValue, fieldValue
			}
		}
		return false, "", ""
	}

	result = make(map[string]string)
	for _, field := range stmt.Schema.FieldsByDBName {
		b, source, to := changed(field)
		if b {
			result[util.ToLower(field.Name)] = fmt.Sprintf("%s -> %s", gconv.String(to), gconv.String(source))
		}
	}

	return result
}
