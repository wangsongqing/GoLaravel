package v1

import (
	"GoLaravel/app/models/link"
	"GoLaravel/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	links := link.AllCached()
	response.Data(c, links)
}
