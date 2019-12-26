package generator

import (
	"goDoc/library/database"
	"strings"
	"goDoc/library/helper"
)

type TableStruct struct {
	Cid       int
	Name      string
	Type      string
	Notnull   int
	DfltValue interface{}
	Pk        int
}

var (
	TbName    string
	TableInfo []TableStruct
)

func init()  {
	TableInfo = make([]TableStruct,0)
}

func GetTableInfo(tbName string) interface{} {
	db := database.GetDbInstance()
	defer db.Close()

	TbName = tbName

	db.Raw("pragma table_info ('" + tbName + "')").Scan(&TableInfo)

	return &TableInfo
}



func MakeControl() {

}

func MakeService() {

}

func makeDataProfiles(hasOrmFlag bool) string {
	dataProfilesString := ""
	if hasOrmFlag {
		dataProfilesString += "gorm.Model\n"
	}
	for _,v := range TableInfo {
		nameText := ""
		nameMap := strings.Split(v.Name,"_")
		for _,v1 := range nameMap {
			nameText += helper.Ucfirst(v1)
		}
		v.Name = "\t" + nameText

		if v.Pk != 1 {
			typeText := strings.ToLower(v.Type)
			if strings.Contains(typeText,"integer") {
				dataProfilesString += v.Name + "\tint\n"
			}else if strings.Contains(typeText,"tinyint") {
				dataProfilesString += v.Name + "\tint\n"
			}else if strings.Contains(typeText,"varchar") {
				dataProfilesString += v.Name + "\tstring\n"
			}else if strings.Contains(typeText,"timestamp") || strings.Contains(typeText,"datetime") {
				dataProfilesString += v.Name + "\ttime.Time\n"
			}
		}
	}

	return dataProfilesString
}

func makeSearchParams()  {
	
}
