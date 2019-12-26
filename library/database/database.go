package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)


var (
	Connection	*gorm.DB
)

func GetDbInstance() *gorm.DB {
	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return "gos_" + defaultTableName
	//}


	//Connection, err := gorm.Open("mysql", "cat:cathyr@(172.16.1.38:3306)/gos_dev?charset=utf8&parseTime=True&loc=Local")
	Connection, err := gorm.Open("sqlite3", "./database/doc.db")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	Connection.SingularTable(true)

	return Connection
}

func CloseDbConnection(){
	Connection.Close()
}

