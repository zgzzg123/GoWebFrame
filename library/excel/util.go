package excel

import (
	"fmt"
	"sort"
)

/**
预处理后所返回的结构体数据类型
 */
type excelColumnsMap struct {
	ColumnFieldsMap []string
	ColumnTitleMap  []string
}

/**
预处理列与抬头、列与字段
 */
func PreColumns(titleMap map[string]string) excelColumnsMap {
	fmt.Println(titleMap)

	tempTitleArr := make([]string, 0)
	tempFieldsArr := make([]string, 0)
	for k, v := range titleMap {
		tempTitleArr = append(tempTitleArr, v)
		tempFieldsArr = append(tempFieldsArr, k)
	}

	sort.Strings(tempFieldsArr)
	sort.Strings(tempTitleArr)
	return excelColumnsMap{
		ColumnFieldsMap: tempFieldsArr,
		ColumnTitleMap:  tempTitleArr,
	}
}
