package catalogModel

import (
	"github.com/jinzhu/gorm"
	"time"
	"goDoc/library/database"
)

type DataProfiles struct {
	gorm.Model
	Pid         int
	ProjectId   int    `gorm:"not null"`
	Name        string `gorm:"not null;unique"`
	Level       int    `gorm:"not null;default:1"`
	SNumber     int    `gorm:"not null;default:99"`
	CreatorName string `gorm:"not null;default:'system'"`
	Status      int    `gorm:"not null;default:1"`
}

type SearchParams struct {
	ID          int
	Pid         int
	ProjectId   int
	Name        string
	Level       int
	SNumber     int
	CreatorName string
	Status      int
	CreatedAtGe time.Time
	CreatedAtLe time.Time
}

func (DataProfiles) TableName() string {
	return "doc_catalog"
}

func ScopeFilter(params SearchParams, sqlObj *gorm.DB) *gorm.DB {

	return sqlObj
}

/**
列表查询
 */
func GetList(params SearchParams) []DataProfiles {
	db := database.GetDbInstance()
	defer db.Close()


	db = ScopeFilter(params, db)
	db = db.Where("id = ?",8)

	var result []DataProfiles
	db.Find(&result)

	return result
}

func GetById(id uint) {

}

/**
创建单条
 */
func CreateUnique(params *DataProfiles) *DataProfiles {
	db := database.GetDbInstance()
	defer db.Close()

	db.Create(&params)

	return params
}

func DeleteById(id uint) {

}
