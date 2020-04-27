package router

import (
	"in-world-server/middleware"
	"in-world-server/pkg/setting"
	"in-world-server/router/api"
	v1 "in-world-server/router/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	r.GET("/users", api.GetUsers)
	r.GET("/users/:id", api.GetUser)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
