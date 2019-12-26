package userControl

import (
	"github.com/gin-gonic/gin"
	"goDoc/library/request"
	"goDoc/models/userModel"
	"fmt"
	//"goDoc/library/generator"
	//"goDoc/services/userService"
	"goDoc/library/weather"
)

func Create(c *gin.Context) {
	user := userModel.DataProfiles{
		UserName: c.PostForm("username"),
		Password: c.PostForm("password"),
		Status:   1,
	}

	result := userModel.CreateUnique(&user)

	c.JSON(200, result)
}

func GetList(c *gin.Context) {
	paramsInterface := request.All(c, userModel.SearchParams{})

	fmt.Println(paramsInterface)
	//tbInfo := generator.GetTableInfo("doc_users")

	//generator.MakeModel("doc_users")

	//params := request.GetPostParams(c,userModel.SearchParams{})
	//
	//getParams := request.GetQueryParams(c,userModel.SearchParams{})

	//jsonParams := request.GetJsonParams(c,userModel.SearchParams{})

	//users := userService.GetList(paramsInterface)

	users := weather.GetWeatherByCity("CHCQ000600")

	c.JSON(200, &users)
}





func Web(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
