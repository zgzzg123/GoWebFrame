package projectsModel

import (
	"goDoc/library/database"
	"time"
	"github.com/jinzhu/gorm"
)

type DataProfiles struct {
	gorm.Model
	Name        string `form:"name" json:"name" binding:"required"`
	Description string
	CreatorName string
	Status      int8 `gorm:"not null;default:1"`
}

type SearchParams struct {
	ID          uint
	Name        string
	Description string
	CreatorName string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time

	Counts int
	Take   int
	Skip   int
}

func (DataProfiles) TableName() string {
	return "doc_projects"
}

func ScopeFilter(params SearchParams, db *gorm.DB) *gorm.DB {

	return db
}

func GetList(params SearchParams) []DataProfiles {
	db := database.GetDbInstance()
	defer db.Close()

	var docs []DataProfiles
	//DocProjects := projectsScopeFilter(params,db)
	db.Find(&docs)

	return docs
}

func GetCount(params SearchParams) *gorm.DB {
	db := database.GetDbInstance()
	defer db.Close()

	sqlObj := ScopeFilter(params,db)
	total := sqlObj.Count("*")

	return total
}

func DocProjectCreateOne(params DataProfiles) *gorm.DB {
	db := database.GetDbInstance()
	defer db.Close()

	return db.Create(&params)
}
