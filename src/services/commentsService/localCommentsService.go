package commentsService

import (
	"fmt"
	"log/slog"

	"github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type LocalCommentsService struct {
	StoredArticles []models.Article
	Logger         slog.Logger
}

func NewLocalCommentsService(articles []models.Article, logger *slog.Logger) LocalCommentsService {
	return LocalCommentsService{
		StoredArticles: articles,
		Logger:         *logger,
	}
}

func (s LocalCommentsService) Get(articleId int) ([]models.ArticleComment, error) {

	for _, article := range s.StoredArticles {
		if article.Id == articleId {
			return article.Comments, nil
		}
	}

	s.Logger.Info("comment failed to retrieve comments (article with ID not found)", slog.Int("ArticleId", articleId))

	return nil, fmt.Errorf("an error ocurred while reading the comments: article with ID %d does not exist", articleId)
}

func (s LocalCommentsService) Store(comment models.ArticleComment) error {

	for idx, article := range s.StoredArticles {
		if article.Id == comment.ArticleId {
			s.StoredArticles[idx].Comments = append(s.StoredArticles[idx].Comments, comment)
			return nil
		}
	}

	s.Logger.Info("comment failed to store (article with ID not found)", slog.Int("ArticleId", comment.ArticleId))

	return fmt.Errorf("an error ocurred while storing the comment: article with ID %d does not exist", comment.ArticleId)
}
