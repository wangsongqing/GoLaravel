//Package topic 模型
package topic

import (
	"GoLaravel/app/models"
	"GoLaravel/app/models/category"
	"GoLaravel/app/models/user"
	"GoLaravel/pkg/database"
)

type Topic struct {
	models.BaseModel

	Title      string `json:"title,omitempty" `
	Body       string `json:"body,omitempty" `
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	// 通过 user_id 关联用户
	User user.User `json:"user"`

	// 通过 category_id 关联分类
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (Topic *Topic) Create() {
	database.DB.Create(&Topic)
}

func (Topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&Topic)
	return result.RowsAffected
}
