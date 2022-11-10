package v1

import (
	"GoLaravel/app/models/topic"
	"GoLaravel/app/policies"
	"GoLaravel/app/requests"
	"GoLaravel/pkg/auth"
	"GoLaravel/pkg/response"
	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Index(c *gin.Context) {
	data, pager := topic.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *TopicsController) Show(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	response.Data(c, topicModel)
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Update(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID

	row := topicModel.Save()
	if row <= 0 {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}

	response.Data(c, row)
}

func (ctrl *TopicsController) Delete(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	row := topicModel.Delete()
	if row <= 0 {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}

	response.Data(c, row)
}

func (ctrl *TopicsController) FindTest(c *gin.Context) {
	//_topic := topic.Topic{
	//	UserID:     "1",
	//	CategoryID: "1",
	//	Title:      "Accusantium voluptatem perferendis sit consequatur aut.",
	//}
	//data := topic.WhereAll(_topic)

	//var _where map[string]string
	//_where["user_id"] = "1"
	//_where["title"] = "Sit perferendis consequatur voluptatem aut accusantium."
	_where := map[string]interface{}{
		// "user_id": c.Query("user_id"),
		"title":   c.Query("title"),
		"user_id": c.PostForm("user_id"),
	}
	_data := topic.FindMap(_where)
	response.Data(c, _data)
}
