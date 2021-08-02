package repo

import (
	config "golang3/config"
	models "golang3/models"
)

func GetAllArticles(article *models.Article, pagination *models.Pagination) (*[]models.Article, int64, error) {
	var articles []models.Article
	var totalRows int64 = 0
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := config.DB.Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.Article{}).Where(article).Find(&articles)
	if result.Error != nil {
		msg := result.Error
		return nil, totalRows, msg
	}
	errCount := config.DB.Model(&models.Article{}).Count(&totalRows).Error
	if errCount != nil {
		return nil, totalRows, errCount
	}
	return &articles, totalRows, nil
}
