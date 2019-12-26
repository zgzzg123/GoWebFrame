package templates

import (
	"github.com/gin-gonic/gin"
	"goDoc/library/request"
	"goDoc/models/userModel"
)

/**
列表查询
 */
func GetList(c *gin.Context)  {
	params := request.All(c, userModel.SearchParams{})


	data := userModel.GetList(params)
}

/**
创建
 */
func Create(c *gin.Context)  {
	
}

/**
详情
 */
func Show(c *gin.Context)  {

}

/**
更新
 */
func Update(c *gin.Context)  {
	
}

/**
删除
 */
func Destroy(c *gin.Context)  {

}