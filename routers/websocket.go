package routers

import (
	"github.com/gin-gonic/gin"
	"goDoc/controller/webSocketControl"
)

func setWsRouter(r *gin.RouterGroup)  {
	r.GET("/ws", webSocketControl.WsHandler)
	r.GET("/pub", webSocketControl.PubMessage)
	r.POST("/pub", webSocketControl.PubMessage)
}