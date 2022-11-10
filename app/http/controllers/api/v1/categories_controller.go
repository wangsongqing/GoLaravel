package v1

import (
	"GoLaravel/app/models/category"
	"GoLaravel/app/requests"
	"GoLaravel/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseAPIController
}

// Store 添加分类
func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

// Update 更新分类
func (ctrl *CategoriesController) Update(c *gin.Context) {

	// 验证 url 参数 id 是否正确
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	if res := categoryModel.Save(); res <= 0 {
		response.Abort500(c, "更新失败~")
		return
	}

	response.Created(c, categoryModel)
}

// Index 分类列表
func (ctrl *CategoriesController) Index(c *gin.Context) {
	data, pager := category.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

// Delete 删除分类
func (ctrl *CategoriesController) Delete(c *gin.Context) {
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := categoryModel.Delete(); ok <= 0 {
		response.Abort500(c, "删除失败~")
		return
	}

	response.Success(c)
}
