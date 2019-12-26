package routers

import (
	"github.com/gin-gonic/gin"
	"goDoc/controller/userControl"
)

func setUserRouter(r *gin.RouterGroup)  {
	// Get user value
	//r.GET("/test", func(c *gin.Context) {
	//	//res := docProjects.DocProjectList(docProjects.DocProjectSearchParams{})
	//	//res1 := docUser.DocUserList(docUser.DocUserSearchParams{})
	//
	//	res := docUser.GetByUserName("Kevin")
	//	fmt.Println(helper.CryptSha256Encode("123456"), res.Password)
	//	res1 := docProjects.GetList(docProjects.SearchParams{})
	//
	//	c.JSON(200, []interface{}{res, res1})
	//})

	//用户列表
	r.GET("userList", userControl.GetList)
	//r.GET("ListenSub", userControl.ListenSub)
	r.POST("userList", userControl.GetList)

	//用户创建
	r.POST("user", userControl.Create)

	r.GET("/web", userControl.Web)

}