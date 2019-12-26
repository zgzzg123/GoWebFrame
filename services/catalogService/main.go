package catalogService

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"goDoc/models/catalogModel"
)

func Create(c *gin.Context) interface{} {
	pid, _ := strconv.Atoi(c.PostForm("pid"))
	projectId, _ := strconv.Atoi(c.PostForm("project_id"))
	level, _ := strconv.Atoi(c.DefaultPostForm("level", "1"))
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "1"))
	sNumber, _ := strconv.Atoi(c.DefaultPostForm("s_number", "99"))
	creatorName := c.DefaultPostForm("creator_name", "system")

	catalogInfo := catalogModel.DataProfiles{
		Pid:         pid,
		ProjectId:   projectId,
		Name:        c.PostForm("name"),
		Level:       level,
		Status:      status,
		CreatorName: creatorName,
		SNumber:     sNumber,
	}
	catalogModel.CreateUnique(&catalogInfo)

	return &catalogInfo
}

func GetList(c *gin.Context) interface{} {
	params := catalogModel.SearchParams{}

	return catalogModel.GetList(params)
}