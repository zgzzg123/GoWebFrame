package request

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"goDoc/library/helper"
	"encoding/json"
)

func All(c *gin.Context,fields interface{}) map[string]interface{} {
	params := GetQueryParams(c,fields)
	postParams := GetPostParams(c,fields)
	if len(postParams) > 0 {
		for k,v := range postParams {
			params[k] = v
		}
	}

	jsonParams := GetJsonParams(c,fields)
	if len(jsonParams) > 0 {
		for k,v := range jsonParams {
			params[k] = v
		}
	}

	return params
}

func GetPostParams(c *gin.Context,fields interface{}) map[string]interface{} {
	params := make(map[string]interface{},0)

	fieldsMap := helper.Struct2Map(fields)
	for k,_ := range fieldsMap {
		val := c.PostForm(k)
		if val != "" {
			params[k] = val
		}
	}

	return params
}

func GetQueryParams(c *gin.Context,fields interface{}) map[string]interface{} {
	params := make(map[string]interface{},0)
	fieldsMap := helper.Struct2Map(fields)
	for k,_ := range fieldsMap {
		val := c.Query(k)
		if val != "" {
			params[k] = val
		}
	}

	return params
}

func GetJsonParams(c *gin.Context,fields interface{}) map[string]interface{} {
	var preParamsMap map[string]interface{}
	params := make(map[string]interface{},0)

	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &preParamsMap)

	fieldsMap := helper.Struct2Map(fields)
	for k,_ := range fieldsMap {
		if preParamsMap[k] != nil {
			params[k] = preParamsMap[k]
		}
	}

	return preParamsMap
}


