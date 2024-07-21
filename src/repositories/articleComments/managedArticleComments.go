package articleComments

import (
	"fmt"
	"log/slog"

	"github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type ManagedArticleCommentsDataRepository struct {
	Logger           slog.Logger
	connectionString string
}

func NewManagedArticleCommentsDataRepository(logger slog.Logger, connString string) ManagedArticleCommentsDataRepository {
	return ManagedArticleCommentsDataRepository{
		Logger:           logger,
		connectionString: connString,
	}
}

func (r ManagedArticleCommentsDataRepository) Get(articleId int) ([]models.ArticleComment, error) {
	return nil, fmt.Errorf("ManagedArticleCommentsDataRepository not implemented")
}

func (r ManagedArticleCommentsDataRepository) Store(comment models.ArticleComment) error {
	return fmt.Errorf("ManagedArticleCommentsDataRepository not implemented")
}
