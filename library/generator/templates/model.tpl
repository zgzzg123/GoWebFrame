package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"goDoc/library/database"
)

/**
* 数据表字段定义
*/
type DataProfiles struct {
	{{dataProfilesContent}}
}

/**
* 用于接收搜索参数
*/
type SearchParams struct {
	{{searchParamsContent}}
}

func (DataProfiles) TableName() string {
	return "{{tbName}}"
}

/**
* 聚集查询条件
*/
func ScopeFilter(params SearchParams, sqlObj *gorm.DB) *gorm.DB {
    {{scopeFilterContent}}

	return sqlObj
}

/**
列表查询
 */
func GetList(params SearchParams) []DataProfiles {
	db := database.GetDbInstance()
	defer db.Close()

	db = ScopeFilter(params, db)

	var result []DataProfiles
	db.Find(&result)

	return result
}

/**
按ID查询
 */
func GetById(id int) DataProfiles {
	db := database.GetDbInstance()
	defer db.Close()

	db.Where("id = ? ", id)

	var result DataProfiles
	db.First(&result)

	return result
}

/**
创建单条数据
 */
func CreateUnique(params *DataProfiles) *DataProfiles {
	db := database.GetDbInstance()
	defer db.Close()

	db.Create(&params)

	return params
}

/**
批量创建
 */
func CreateBatch(params []*DataProfiles) []interface{} {
	db := database.GetDbInstance()
	defer db.Close()

	result := make([]interface{},0)
	for _,v := range params {
		db.Create(&v)
		result = append(result,v)
	}

	return result
}

/**
按ID删除
 */
func DeleteById(id int) interface{} {
	db := database.GetDbInstance()
	defer db.Close()

	result := &DataProfiles{}
	db.Where("id = ?",id).Delete(&result)

	return &result
}

/**
批量删除
 */
func DeleteBatch(idMap []int) interface{} {
	db := database.GetDbInstance()
	defer db.Close()

	result := &DataProfiles{}
	db.Where("id in (?)",idMap).Delete(&result)

	return &result
}