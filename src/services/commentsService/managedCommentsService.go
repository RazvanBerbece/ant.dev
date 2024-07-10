package commentsService

import (
	"fmt"
	"log/slog"

	"github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type ManagedCommentsService struct {
	// Repository interface{}
	Logger slog.Logger
}

func NewManagedCommentsService(logger *slog.Logger) ManagedCommentsService {
	return ManagedCommentsService{
		Logger: *logger,
	}
}

func (s ManagedCommentsService) Get(articleId int) ([]models.ArticleComment, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s ManagedCommentsService) Store(comment models.ArticleComment) error {
	return fmt.Errorf("not implemented")
}
