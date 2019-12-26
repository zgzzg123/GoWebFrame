package excel

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"reflect"
)

var (
	columnsTitle = TitleMap{}
	fieldValue   string
)

type OutputOptions struct {
	FileName  string
	FilePath  string
	TitleMap  TitleMap
	SheetsMap []string
}

/**
全部Title
 */
type TitleMap []uniqueTitle

/**
单个Title结构
 */
type uniqueTitle struct {
	FieldName string
	TitleName string
	Type      string
}

func SetColumnsTitle(titleMap TitleMap) {
	columnsTitle = titleMap
}

func Write(outOptions OutputOptions, data []map[string]interface{}) string {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	//预处理第一行抬头和每列所对应的字段名

	row = sheet.AddRow()
	for _, v := range columnsTitle {
		cell = row.AddCell()
		cell.Value = string(v.TitleName)
	}

	for _, item := range data {
		row = sheet.AddRow()
		for _, v := range columnsTitle {
			cell = row.AddCell()

			str := item[v.FieldName]
			if str != nil {
				types := reflect.TypeOf(str)
				switch types.Name() {
				case "int64":
					cell.SetInt64(str.(int64))
				case "int":
					cell.SetInt(str.(int))
				case "float64":
					cell.SetFloat(str.(float64))
				case "string":
					cell.Value = str.(string)
				default:
					if str, ok := item[v.FieldName].(string); ok {
						cell.Value = str
					} else {
						cell.Value = ""
					}
				}
			} else {
				cell.Value = ""
			}
		}
	}

	err = file.Save(outOptions.FilePath + outOptions.FileName)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return outOptions.FilePath + outOptions.FileName
}
