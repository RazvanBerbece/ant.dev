package commentsService

import (
	"fmt"
	"log/slog"

	"github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type CloudCommentsService struct {
	// Repository interface{}
	Logger slog.Logger
}

func NewCloudCommentsService(logger *slog.Logger) CloudCommentsService {
	return CloudCommentsService{
		Logger: *logger,
	}
}

func (s CloudCommentsService) Get(articleId int) ([]models.ArticleComment, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s CloudCommentsService) Store(comment models.ArticleComment) error {
	return fmt.Errorf("not implemented")
}
