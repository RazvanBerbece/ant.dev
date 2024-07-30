package articlesService

import (
	"log/slog"

	"github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type LocalArticlesService struct {
	Logger         slog.Logger
	StoredArticles []models.Article
}

func NewLocalArticlesService(logger slog.Logger, localArticles []models.Article) LocalArticlesService {
	return LocalArticlesService{
		Logger:         logger,
		StoredArticles: localArticles,
	}
}

func (s LocalArticlesService) GetArticle(id int) *models.Article {

	if len(s.StoredArticles) == 0 {
		return nil
	}

	for _, article := range s.StoredArticles {
		if article.Id == id {
			go s.Logger.Info("Retrieved article from runtime slice", "article_id", id)
			return &article
		}
	}

	return nil

}

func (s LocalArticlesService) GetArticles() []models.Article {

	return s.StoredArticles

}

func (s LocalArticlesService) GetRecentArticles(n int) []models.Article {

	// TODO: Actually implement this, instead of just returning all articles
	return s.GetArticles()

}
