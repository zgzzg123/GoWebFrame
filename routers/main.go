package routers

import (
	"github.com/gin-gonic/gin"
	"goDoc/services/catalogService"
	"goDoc/models/projectsModel"
	"fmt"
	"net/http"
)



func setRouter2(r *gin.Engine)  {
	//分类创建
	r.POST("catalog", func(c *gin.Context) {
		catalogInfo := catalogService.Create(c)
		c.JSON(200, &catalogInfo)
	})

	r.GET("projects/store", func(c *gin.Context) {
		params := projectsModel.DataProfiles{
			Name:        "中启行油品运营系统",
			Description: "中启行油品运营系统",
			CreatorName: "System",
			Status:      1,
		}

		res := projectsModel.DocProjectCreateOne(params)
		fmt.Println(res)

		c.JSON(200, res)
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	//	"foo":  "bar", // user:foo password:bar
	//	"manu": "123", // user:manu password:123
	//}))

	//authorized.POST("admin", func(c *gin.Context) {
	//
	//})
}

func SetupRouter(r *gin.Engine) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	//r := gin.Default()
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	r.Use(gin.Logger())



	r.Use(func(c *gin.Context) {
		//放行所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		if c.Request.Header.Get("Origin") == "http://dev.zqx.chinawayltd.com" {
			c.Request.Header.Del("Origin")
		}
	})

	r.LoadHTMLFiles("public/*")
	v1 := r.Group("/v1")
	{
		setUserRouter(v1)
		setCatalogRouter(v1)
		setWsRouter(v1)
	}


	return r
}