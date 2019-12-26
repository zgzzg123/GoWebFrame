package main

import (
	"goDoc/routers"
	"goDoc/controller/webSocketControl"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	go webSocketControl.ListenSub()

	r := gin.New()
	//http.Handle("/", http.FileServer(http.Dir("public")))
	r = routers.SetupRouter(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8085")
}