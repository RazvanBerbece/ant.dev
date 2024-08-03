package articles

import (
	"fmt"
	"log/slog"

	"github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type LocalArticlesDataRepository struct {
	Logger        slog.Logger
	LocalArticles []models.Article
}

func NewLocalArticlesDataRepository(logger slog.Logger, localArticles []models.Article) LocalArticlesDataRepository {
	return LocalArticlesDataRepository{
		Logger:        logger,
		LocalArticles: localArticles,
	}
}

func (r LocalArticlesDataRepository) GetById(articleId int) (*models.Article, error) {

	for _, article := range r.LocalArticles {
		if article.Id == articleId {
			return &article, nil
		}
	}

	r.Logger.Info("failed to retrieve article by ID (article with ID not found)", slog.Int("ArticleId", articleId))

	return nil, fmt.Errorf("an error ocurred while retrieving an article: article with ID %d does not exist", articleId)
}

func (r LocalArticlesDataRepository) GetAll() ([]models.Article, error) {

	return r.LocalArticles, nil

}
