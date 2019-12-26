package export

import (
	"os"
	"fmt"
	"syscall"
	"foss/library/helper"
	"foss/library/aliyunOss"
	"foss/models"
	"math"
	"foss/library/excel"
)

var (
	TitleColumns = excel.TitleMap{
		{
			FieldName: "id",
			TitleName: "ID",
			Type:      "string",
		},
		{
			FieldName: "main_no",
			TitleName: "主卡号",
			Type:      "string",
		},
		{
			FieldName: "card_owner",
			TitleName: "持卡人",
			Type:      "string",
		},
		{
			FieldName: "qz_drivername",
			TitleName: "司机姓名",
			Type:      "string",
		},
		{
			FieldName: "qz_drivertel",
			TitleName: "司机电话",
			Type:      "string",
		},
		{
			FieldName: "oil_com",
			TitleName: "油卡类型",
			Type:      "string",
		},
		{
			FieldName: "card_from",
			TitleName: "卡来源",
			Type:      "string",
		},
		{
			FieldName: "regions_name",
			TitleName: "积分可用地区",
			Type:      "string",
		},
		{
			FieldName: "consume_region",
			TitleName: "消费地区",
			Type:      "string",
		},
		{
			FieldName: "org_name",
			TitleName: "机构名称",
			Type:      "string",
		},
		{
			FieldName: "orgcode",
			TitleName: "机构编码",
			Type:      "string",
		},
		{
			FieldName: "main_operator_name",
			TitleName: "主卡运营商",
			Type:      "string",
		},
		{
			FieldName: "oil_operator_name",
			TitleName: "机构运营商",
			Type:      "string",
		},
		{
			FieldName: "vice_no",
			TitleName: "卡号",
			Type:      "string",
		},
		{
			FieldName: "consume_type",
			TitleName: "消费类型",
			Type:      "string",
		},
		{
			FieldName: "truck_no",
			TitleName: "车牌号",
			Type:      "string",
		},
		{
			FieldName: "trade_time",
			TitleName: "交易时间",
			Type:      "string",
		},
		{
			FieldName: "trade_type",
			TitleName: "交易类型",
			Type:      "string",
		},
		{
			FieldName: "trade_money",
			TitleName: "金额",
			Type:      "float64",
		},
		{
			FieldName: "use_fanli_money",
			TitleName: "使用返利",
			Type:      "string",
		},
		{
			FieldName: "oil_name",
			TitleName: "油品",
			Type:      "string",
		},
		{
			FieldName: "oil_type",
			TitleName: "油品类型",
			Type:      "string",
		},
		{
			FieldName: "trade_num",
			TitleName: "数量",
			Type:      "float64",
		},
		{
			FieldName: "trade_price",
			TitleName: "单价",
			Type:      "float64",
		},
		{
			FieldName: "trade_jifen",
			TitleName: "奖励分值",
			Type:      "float64",
		},
		{
			FieldName: "fanli_money",
			TitleName: "返利金额",
			Type:      "float64",
		},
		{
			FieldName: "fanli_jifen",
			TitleName: "返利积分",
			Type:      "float64",
		},
		{
			FieldName: "balance",
			TitleName: "余额",
			Type:      "float64",
		},
		{
			FieldName: "trade_place",
			TitleName: "地点",
			Type:      "string",
		},
		{
			FieldName: "fetch_time",
			TitleName: "抓取时间",
			Type:      "string",
		},
		{
			FieldName: "createtime",
			TitleName: "创建时间",
			Type:      "string",
		},
		{
			FieldName: "is_fanli",
			TitleName: "计算返利",
			Type:      "string",
		},
		{
			FieldName: "is_test",
			TitleName: "测试机构",
			Type:      "string",
		},
	}
)

const PageSize = 10000

func TradesExport(params models.TradesSearchParams) (string, string) {
	localFile := ""
	ossFileUrl := ""

	excel.SetColumnsTitle(TitleColumns)

	countParams := params
	countParams.Counts = 1
	_, total := models.GetTradesList(countParams)
	if total == 0 {
		fmt.Println("没有可以导出的数据")
	} else {
		totalPageTmp := float64(total / PageSize)
		totalPage := math.Ceil(totalPageTmp)

		osFileMap := make([]*os.File, 0)
		i := 1
		params.Take = PageSize

		for i <= int(totalPage) {
			params.Skip = (i - 1) * PageSize

			ch := make(chan string)
			go exportTrades(params, ch)

			xlsFilePath := <-ch
			f1, err := os.Open(xlsFilePath)

			if err != nil {
				fmt.Println(err.Error())
			}

			osFileMap = append(osFileMap, f1)
			defer f1.Close()

			defer syscall.Unlink(xlsFilePath)

			i++
		}

		fileName := helper.Guid() + ".zip"
		dest := "./download/" + fileName
		err := helper.Compress(osFileMap, dest)
		if err != nil {
			fmt.Println(err.Error())
		}

		ossFileUrl, err := aliyunOss.PutObjectFromFile(fileName, dest)
		if err != nil {
			fmt.Print(err.Error())
		}

		return dest, ossFileUrl
	}

	return localFile, ossFileUrl

}

func exportTrades(params models.TradesSearchParams, ch chan string) string {
	var filePath string
	tradesRecords := make([]map[string]interface{}, 0)

	subPageSize := 10000
	totalCount := PageSize / subPageSize

	i := 0
	for i < totalCount {
		if i > 0 {
			params.Skip += subPageSize
		}

		params.Take = subPageSize

		fmt.Println(params)
		res, total := models.GetTradesList(params)
		if total == 0 {
			return ""
		}
		for _, v := range res {
			tradesRecords = append(tradesRecords, v)
		}

		fmt.Println("thisTime:",len(res),"total:",len(tradesRecords))

		res = nil
		i++
	}

	filePath = excel.Write(excel.OutputOptions{TitleMap: TitleColumns, FilePath: "./download/", FileName: helper.Guid() + ".xlsx"}, tradesRecords)
	ch <- filePath

	return filePath
}
