package topic

import (
	"GoLaravel/pkg/app"
	"GoLaravel/pkg/database"
	"GoLaravel/pkg/paginator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func Get(idstr string) (topic Topic) {
	database.DB.Preload(clause.Associations).Where("id", idstr).First(&topic)
	return
}

func GetBy(field, value string) (topic Topic) {
	database.DB.Where("? = ?", field, value).First(&topic)
	return
}

func All() (topics []Topic) {
	database.DB.Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

// WhereAll 结构体查询
func WhereAll(_topics Topic) (topics []Topic) {
	database.DB.Where(&_topics).Find(&topics)
	return
}

// FindAll 多条件查询
func FindAll(userId interface{}, id interface{}) (topics []Topic) {
	database.DB.Where("user_id = ? AND id >= ?", userId, id).Find(&topics)
	return
}

// FindMap map 条件查询 _where := map[string]interface{}{"user_id": "2222", "title": "AAA"}
func FindMap(mapWhere map[string]interface{}) (topics []Topic) {
	database.DB.Where(mapWhere).Find(&topics)
	return
}

// Delete 删除
func (Topic *Topic) Delete() (RowsAffected int64) {
	result := database.DB.Delete(&Topic)
	return result.RowsAffected
}

func Paginate(c *gin.Context, perPage int) (users []Topic, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Topic{}),
		&users,
		app.V1URL(database.TableName(&Topic{})),
		perPage,
	)
	return
}
