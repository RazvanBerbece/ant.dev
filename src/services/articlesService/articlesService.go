package articlesService

import (
	domain "github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type ArticlesService interface {
	GetArticle(id int) *domain.Article
	GetArticles() []domain.Article
	GetRecentArticles(n int) []domain.Article
}
