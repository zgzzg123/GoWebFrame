package routers

import (
	"github.com/gin-gonic/gin"
	"goDoc/controller/catalogControl"
)

func setCatalogRouter(r *gin.RouterGroup)  {
	r.GET("/catalog", catalogControl.GetList)
	r.POST("/catalog", catalogControl.Create)
	r.PUT("/catalog", catalogControl.Update)
	r.DELETE("/catalog", catalogControl.Destroy)
}