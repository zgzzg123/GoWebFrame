package catalogControl

import (
	"goDoc/services/catalogService"
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context)  {
	result := catalogService.GetList(c)

	c.JSON(200,&result)
}

func Create(c *gin.Context)  {

}

func Update(c *gin.Context)  {

}

func Destroy(c *gin.Context)  {

}