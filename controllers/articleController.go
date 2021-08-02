package Controllers

import (
	"fmt"
	config "golang3/config"
	models "golang3/models"
	"golang3/repo"
	utils "golang3/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Create
func CreateArticles(c *gin.Context) {
	var articles models.Article
	c.BindJSON(&articles)
	config.DB.Create(&articles)
	c.JSON(http.StatusOK, articles)
}

// Read
func GetAllArticles(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var article models.Article
	articleLists, totalRows, err := repo.GetAllArticles(&article, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":      articleLists,
		"totalList": totalRows,
	})
}

func GetArticlesByID(c *gin.Context) {
	var articles models.Article
	id := c.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&articles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, articles)
	}
}

func UpdateArticlesByID(c *gin.Context) {
	var articles models.Article
	id := c.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&articles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&articles)
	config.DB.Save(&articles)
	c.JSON(200, articles)
}

func DeleteArticlesByID(c *gin.Context) {
	var articles models.Article
	id := c.Params.ByName("id")
	d := config.DB.Where("id = ?", id).Delete(&articles)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
