package classes

import (
	"github.com/pkg/errors"
	"github.com/gin-gonic/gin"
	"LearnOnline/infra/model"
	"strconv"
	"LearnOnline/infra/init"
	"net/http"
	"time"
)

var (
	ErrorClass      = errors.New("not found record")
	ErrorClassId    = errors.New("id is not allow")
	ErrorClassParam = errors.New("list team param error")
)

type CreateClassParam struct {
	Name string 	`form:"name"`
	Time string 	`form:"time"`
	Desc string 	`form:"desc"`
	Icon string 	`form:"icon"`
	Status string 	`form:"status"`
}

// CreateClassHandler will create classes
// @Summary create classes
// @Accept json
// @Tags classes
// @Security Bearer
// @Produce  json
// @Param classname query string true "name"
// @Param classtime query string true "time"
// @Param classdesc query string true "desc"
// @Param classicon query string true "icon"
// @Param classstatus query string true "status"
// @Resource classes
// @Router /classes/create [post]
// @Success 200 {string} string "success!"
func CreateClassHandler(c *gin.Context) {
	var param CreateClassParam
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(400, c.AbortWithError(400, err))
		return
	}

	var class model.ClassManage
	class.ClassName = param.Name
	class.ClassTime,_ = time.Parse("2006-01-02 15:04:05",param.Time)
	class.ClassDesc = param.Desc
	class.ClassIcon = param.Icon
	class.ClassStatus, _ = strconv.Atoi(param.Status)
	class.ClassVis = 0
	class.ClassVie = 0

	if dbError := initiator.POSTGRES.Create(&class).Error; dbError != nil{
		c.AbortWithError(400, dbError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success!",
	})
}


type EditClassParam struct {
	ClassID string 	`form:"id"`
	Name 	string 	`form:"name"`
	Time 	string 	`form:"time"`
	Desc 	string 	`form:"desc"`
	Icon 	string 	`form:"icon"`
	Status 	string 	`form:"status"`
}
// EditClassHandler will edit classes
// @Summary edit classes
// @Accept json
// @Tags classes
// @Security Bearer
// @Produce  json
// @Param classid query string true "id"
// @Param classname query string true "name"
// @Param classtime query string true "time"
// @Param classdesc query string true "desc"
// @Param classicon query string true "icon"
// @Param classstatus query string true "status"
// @Resource classes
// @Router /classes/create [post]
// @Success 200 {string} string "success!"
func EditClassHandler(c *gin.Context) {
	var param EditClassParam
	if err := c.ShouldBindQuery(&param); err != nil{
		c.JSON(400, c.AbortWithError(400, err))
		return
	}

	var class model.ClassManage
	classid,_ := strconv.Atoi(param.ClassID)
	if dbError := initiator.POSTGRES.Where("id = ?", classid).First(&class).Error; dbError != nil{
		c.JSON(400, c.AbortWithError(400, dbError))
		return
	}
	t ,_ := time.Parse("2006-01-02 15:04:05",param.Time)
	s, _ := strconv.Atoi(param.Status)

	if class.ClassName != param.Name {
		class.ClassName = param.Name
	}
	if t != class.ClassTime {
		class.ClassTime = t
	}
	if class.ClassDesc != param.Desc{
		class.ClassDesc = param.Desc
	}
	if s != class.ClassStatus{
		class.ClassStatus = s
	}
	initiator.POSTGRES.Save(&class)

}


type ListClassParam struct {
	Page string 	`form:"page"`
	Number string 	`form:"number"`
	Status string 	`form:"status"`
	Keyword string 	`form:"key"`
}

// ListClassHandler will list classes
// @Summary List classes
// @Accept json
// @Tags classes
// @Security Bearer
// @Produce  json
// @Param page query string true "page"
// @Param number query string true "limit"
// @Param status query string false "status"
// @Param keyword query string false "key"
// @Resource classes
// @Router /classes/find [get]
// @Success 200 {array} model.ClassSerializer
func ListClassHandler(c *gin.Context) {

	var param ListClassParam
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(400, c.AbortWithError(400, err))
		return
	}

	var classes []model.ClassManage

	page, _ := strconv.Atoi(param.Page)
	num, _ := strconv.Atoi(param.Number)
	sta, _ := strconv.Atoi(param.Status)

	if page == 0 || num == 0 {
		c.JSON(400, c.AbortWithError(400, ErrorClassParam))
		return
	}else if param.Keyword == "" {
		if dbError := initiator.POSTGRES.Where("class_status = ?",sta).Limit(num).Offset((page-1)*num).Find(&classes).Error; dbError != nil{
			c.AbortWithError(404, ErrorClass)
		}
	}else {
		if dbError := initiator.POSTGRES.Where("class_status = ?",sta, "class_name = ?",param.Keyword).Limit(num).Offset((page-1)*num).Find(&classes).Error; dbError != nil{
			c.AbortWithError(404,ErrorClass)
		}
	}

	result := make([]model.ClassSerializer, len(classes))
	for index, cla := range classes{
		result[index] = cla.Serializer()
	}
	c.JSON(
		http.StatusOK,
		result,
	)
}