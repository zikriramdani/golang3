package controllers

import (
	"encoding/json"
	"golang1/database"
	"golang1/entity"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	pageSize = 10
)

// Pagination
const getPagination = (page, size) => {
    const limit = size ? +size : 8;
    const offset = page ? page * limit : 0;
  
    return { limit, offset };
};
const getPagingData = (data, page, limit) => {
    const { count: totalItems, rows: getData } = data;
    const currentPage = page ? +page : 0;
    const totalPages = Math.ceil(totalItems / limit);
  
    return { totalItems, getData, totalPages, currentPage };
};

// GetAllArticle get all article data
func GetAllArticle(w http.ResponseWriter, r *http.Request) {
	var article []entity.Article
	database.Connector.Find(&article)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)
}

// GetArticleByID returns article with specific ID
func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var article entity.Article
	database.Connector.First(&article, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// GetArticles returns all articles from the database
// func GetArticlePagination(w http.ResponseWriter, r *http.Request) {
// 	// articles := &types.ArticleList{}
// 	// c.Client.Where("id >= ?", pageID).Order("id").Limit(pageSize + 1).Find(&articles.Items)
// 	// if len(articles.Items) == pageSize+1 {
// 	// 	articles.NextPageID = articles.Items[len(articles.Items)-1].ID
// 	// 	articles.Items = articles.Items[:pageSize]
// 	// }
// 	// return articles

// 	const { page, limits } = req.query;
//     const { limit, offset } = getPagination(page, limits);

// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	var article entity.Article
// 	id, _ := strconv.ParseInt(key, 10, 64)
// 	database.Connector.Where("id = ?", id).Limit(pageSize + 1).Find(&article)
// 	// if len(article) == pageSize+1 {
// 	// 	articles.NextPageID = article[len(article)-1].ID
// 	// 	article = article[:pageSize]
// 	// }
// 	// return article
// }

// GetArticles returns all articles from the database
func (c *Client) GetArticlePagination(pageID int) *types.ArticleList {
	articles := &types.ArticleList{}
	c.Client.Where("id >= ?", pageID).Order("id").Limit(pageSize + 1).Find(&articles.Items)
	if len(articles.Items) == pageSize+1 {
		articles.NextPageID = articles.Items[len(articles.Items)-1].ID
		articles.Items = articles.Items[:pageSize]
	}
	return articles
}

// CreateArticle creates article
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var article entity.Article
	json.Unmarshal(requestBody, &article)

	database.Connector.Create(article)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

// UpdateArticleByID updates article with respective ID
func UpdateArticleByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var article entity.Article
	json.Unmarshal(requestBody, &article)
	database.Connector.Save(&article)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)
}

// DeletArticleByID delete's article with specific ID
func DeleteArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var article entity.Article
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&article)
	w.WriteHeader(http.StatusNoContent)
}
