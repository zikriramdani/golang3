package controllers

import (
	"encoding/json"
	"golang3/database"
	"golang3/entity"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	limit = 8
	page  = 1
)

func Pagination(article *entity.Article, pagination *entity.Pagination) {
	var articles []entity.Article
	offset := (pagination.Page - 1) * pagination.Limit
	database.Connector.Where(article).Limit(pagination.Limit).Offset(offset).Find(&articles)
}

// GetAllArticle get all article data
func GetAllArticle(w http.ResponseWriter, r *http.Request) {
	var article []entity.Article

	database.Connector.Where(article).Limit(limit).Offset(page).Find(&article)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(article)

	// RequestURI: https://example.com/?limit=10&page=1
	// p := pagination.ParseQuery(r.URL.RequestURI())
	// fetcher := newFruitFetcher()

	// totalCount, totalPages, res, err := pagination.Fetch(fetcher, &pagination.Setting{
	// 	Limit: p.Limit,
	// 	Page:  p.Page,
	// })

	// if err != nil {
	// 	w.Header().Set("Content-Type", "text/html; charset=utf8")
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "something wrong: %v", err)
	// 	return
	// }

	// w.Header().Set("X-Total-Count", strconv.Itoa(totalCount))
	// w.Header().Set("X-Total-Pages", strconv.Itoa(totalPages))
	// w.Header().Set("Access-Control-Expose-Headers", "X-Total-Count,X-Total-Pages")
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// w.WriteHeader(200)
	// resJSON, _ := json.Marshal(res)
	// w.Write(resJSON)
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
// func (c *Client) GetArticlePagination(pageID int) *types.ArticleList {
// 	articles := &types.ArticleList{}
// 	c.Client.Where("id >= ?", pageID).Limit(pageSize + 1).Find(&articles.Items)
// 	if len(articles.Items) == pageSize+1 {
// 		articles.NextPageID = articles.Items[len(articles.Items)-1].ID
// 		articles.Items = articles.Items[:pageSize]
// 	}
// 	return articles
// }

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
