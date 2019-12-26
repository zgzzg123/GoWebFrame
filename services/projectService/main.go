package projectService

import (
	"github.com/gin-gonic/gin"
	"goDoc/models/projectsModel"
)

func Create(c *gin.Context)  {
	params := projectsModel.DataProfiles{
		Name:        "中启行油品运营系统",
		Description: "中启行油品运营系统",
		CreatorName: "System",
		Status:      1,
	}
}