package api_server

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"net/http"
	"fmt"
	"LearnOnline/ui/api-server/controller"
	"LearnOnline/ui/api-server/classes"
	"LearnOnline/ui/api-server/admins"
	"LearnOnline/ui/api-server/accesstoken"
)

type APIServer struct {
	engine * gin.Engine
}

func (a *APIServer) registry()  {
	APIServerInit(a.engine)
}

func (a *APIServer) init()  {
	
}

type Welcome struct {
	Greet string `json:"greet" binding:"required"`
	Words string `json:"words" binding:"required"`
}

func APIServerInit(r *gin.Engine)  {
	// docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 使用curl测试：
	// curl -X POST http://localhost:5005/welcome -H 'content-type:application/json' -d '{"greet":"jjjjjjjjj","words":"888888"}'
	r.POST("/welcome", func(context *gin.Context) {
		var welcome Welcome
		if err := context.ShouldBindJSON(&welcome); err == nil {
			context.JSON(
				http.StatusOK, gin.H{
					"status": fmt.Sprintf("%s : %s", welcome.Words, welcome.Greet),})
		} else {
			context.JSON(
				http.StatusAccepted,
				gin.H{
					"err": err.Error(),
				},
			)
		}

	})

	v1 := r.Group("/v1/api")
	// 欢迎页
	indexRegistry(v1)
	// 管理员登录
	adminsRegistry(v1)

	v1.Use(controller.AuthRequired())
	{
		//获取accesstoken
		classGetAccessToken(v1)
		//课程管理API
		classMessageRegistry(v1)
		//直播（stream）管理API
		radioStreamRegistry(v1)
	}
}

func Hello(c *gin.Context)  {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Hello everyone! Welcome to class"},
	)
}

func indexRegistry(r *gin.RouterGroup) {
	r.GET("", Hello)
}

func adminsRegistry(r *gin.RouterGroup)  {
	r.POST("/admins/sign_in", admins.SignIn)
	r.POST("/admins/sign_up", admins.SignUp)
}

func classMessageRegistry(r *gin.RouterGroup)  {
	r.POST("/classes/create", classes.CreateClassHandler)
	r.POST("/classes/edit", classes.EditClassHandler)
	r.GET("/classes/find", classes.ListClassHandler)
}

func classGetAccessToken(r *gin.RouterGroup)  {
	r.POST("/accesstoken", accesstoken.TrapGetAccessToken)
}

func radioStreamRegistry(r *gin.RouterGroup)  {
	r.POST("/stream/create", )
}






func (a *APIServer) Start()  {
	a.registry()
	a.engine.Run(":5005")
}

func New() *APIServer {
	return &APIServer{
		engine: gin.Default(),
	}
}
