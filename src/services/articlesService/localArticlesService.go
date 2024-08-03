package articlesService

import (
	"log/slog"

	"github.com/RazvanBerbece/ant.dev/src/domain/models"
	"github.com/RazvanBerbece/ant.dev/src/repositories/articles"
)

type LocalArticlesService struct {
	Logger             slog.Logger
	articlesRepository articles.ArticlesDataRepository
}

func NewLocalArticlesService(logger slog.Logger, articlesRepository articles.ArticlesDataRepository) LocalArticlesService {
	return LocalArticlesService{
		Logger:             logger,
		articlesRepository: articlesRepository,
	}
}

func (s LocalArticlesService) GetArticle(id int) *models.Article {

	article, err := s.articlesRepository.GetById(id)
	if err != nil {
		s.Logger.Error(err.Error(),
			slog.Group("context",
				"func", "GetArticle",
				"article_id", id,
			),
		)
		return nil
	}

	if article != nil {
		go s.Logger.Info(
			"Retrieved article from runtime slice",
			slog.Group("context",
				"func", "GetArticle",
				"article_id", id,
			),
		)
		return article
	}

	return nil

}

func (s LocalArticlesService) GetArticles() []models.Article {

	articles, err := s.articlesRepository.GetAll()
	if err != nil {
		s.Logger.Error(err.Error(),
			slog.Group("context",
				"func", "GetArticles",
			),
		)
		return nil
	}

	return articles

}

func (s LocalArticlesService) GetRecentArticles(n int) []models.Article {

	// TODO: Actually implement this, instead of just returning all articles
	return s.GetArticles()

}
