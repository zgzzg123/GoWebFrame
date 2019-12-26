package excel

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"foss/library/helper"
)

func Read(titleMap map[string]string, filePath string) [][]interface{} {
	data := make([][]interface{}, 0)

	titleArr := make([]string, 0)
	titleMaps := make([]string, 0)

	_,err := helper.FileExist(filePath)
	if err != nil {
		panic(err.Error())
	}

	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, sheet := range xlFile.Sheets {
		dataSheet := make([]interface{}, 0)
		for kk, row := range sheet.Rows {
			dataRow := make(map[string]string, 0)
			if kk == 0 {
				for _, cell := range row.Cells {
					text := cell.String()
					if kk == 0 {
						titleArr = append(titleArr, text)
					}
				}
				titleMaps = preExcelTitleMap(titleMap, titleArr)
			}else if kk > 0 {
				for kkk, cell := range row.Cells {
					if titleMaps[kkk] != "" {
						text := cell.String()
						title := titleMaps[kkk]
						dataRow[title] = text
					}
				}
			}

			if kk > 0 {
				dataSheet = append(dataSheet, dataRow)
			}
		}

		data = append(data, dataSheet)
	}

	return data
}

func preExcelTitleMap(titleMap map[string]string, columnTitleMap []string) []string {
	titleMaps := make([]string, 0)

	for _, columnName := range columnTitleMap {
		for titleName, fieldName := range titleMap {
			if titleName == columnName {
				titleMaps = append(titleMaps,fieldName)
				break
			}
		}
	}

	return titleMaps
}
