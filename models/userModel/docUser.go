package userModel

import (
	"github.com/jinzhu/gorm"
	"goDoc/library/database"
	"goDoc/library/helper"
	"time"
	"strings"
)

type DataProfiles struct {
	gorm.Model
	UserName string
	Password string
	Status   int8
}

type SearchParams struct {
	ID        int
	UserName  string
	Password  string
	Status    int
	StatusIn  []interface{}
	CreatedAt time.Time
	CreatedAtGe time.Time
	CreatedAtLe time.Time
}

func ScopeFilter(params map[string]interface{}, sqlObj *gorm.DB) *gorm.DB {
	if params["ID"] != nil {
		sqlObj = sqlObj.Where("ID = ? ", params["ID"])
	}

	if params["UserName"] != nil {
		sqlObj = sqlObj.Where(" user_name like ?", "%"+params["UserName"].(string)+"%")
	}

	if params["Status"] != nil {
		sqlObj = sqlObj.Where(" status in (?) ", strings.Split(params["Status"].(string),","))
	}

	return sqlObj
}

func (DataProfiles) TableName() string {
	return "doc_users"
}

func GetList(params map[string]interface{}) interface{} {
	db := database.GetDbInstance()
	defer db.Close()

	db = ScopeFilter(params, db)

	var docUsers []DataProfiles
	db.Find(&docUsers)

	return docUsers
}

/**
创建单条记录
 */
func CreateUnique(params *DataProfiles) *DataProfiles {
	db := database.GetDbInstance()
	defer db.Close()

	params.Password = helper.CryptSha256Encode(params.Password)

	db.Create(&params)

	return params
}

/**
按ID获取记录
 */
func GetById(id int) *DataProfiles {
	db := database.GetDbInstance()
	defer db.Close()

	var user DataProfiles

	db.Where("id =?",id).First(&user)

	return &user
}

/**
按UserName获取
 */
func GetByUserName(userName string) *DataProfiles {
	db := database.GetDbInstance()
	defer db.Close()

	var user DataProfiles

	db.Where("user_name = ?",userName).First(&user)

	return &user
}
