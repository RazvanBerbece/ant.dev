package commentsService

import (
	"fmt"
	"log/slog"

	"github.com/RazvanBerbece/ant.dev/src/domain"
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

func (s CloudCommentsService) Get(articleId int) ([]domain.ArticleComment, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s CloudCommentsService) Store(comment domain.ArticleComment) error {
	return fmt.Errorf("not implemented")
}
