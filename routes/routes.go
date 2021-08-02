package routes

import (
	controllers "golang3/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()

	r.POST("api/v1/article/add", controllers.CreateArticles)
	r.GET("api/v1/articleList", controllers.GetAllArticles)
	r.GET("api/v1/articleList/:id", controllers.GetArticlesByID)
	r.PUT("api/v1/article/:id", controllers.UpdateArticlesByID)
	r.DELETE("api/v1/article/:id", controllers.DeleteArticlesByID)

	return r
}
