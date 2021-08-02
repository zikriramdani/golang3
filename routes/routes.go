package routes

import (
	controllers "golang3/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()

	// Router Articles
	r.POST("api/v1/article/add", controllers.CreateArticles)       // Create
	r.GET("api/v1/articleList", controllers.GetAllArticles)        // Read
	r.GET("api/v1/articleList/:id", controllers.GetArticlesByID)   // Read ByID
	r.PUT("api/v1/article/:id", controllers.UpdateArticlesByID)    // Update ByID
	r.DELETE("api/v1/article/:id", controllers.DeleteArticlesByID) // Delete ByID

	return r
}
