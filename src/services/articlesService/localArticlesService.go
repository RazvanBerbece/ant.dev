package articlesService

import (
	"github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type LocalArticlesService struct {
	StoredArticles []models.Article
}

func NewLocalArticlesService(localArticles []models.Article) LocalArticlesService {
	return LocalArticlesService{
		StoredArticles: localArticles,
	}
}

func (s LocalArticlesService) GetArticle(id int) *models.Article {

	if len(s.StoredArticles) == 0 {
		return nil
	}

	for _, article := range s.StoredArticles {
		if article.Id == id {
			return &article
		}
	}

	return nil

}

func (s LocalArticlesService) GetArticles() []models.Article {

	return s.StoredArticles

}

func (s LocalArticlesService) GetRecentArticles(n int) []models.Article {

	return nil

}
