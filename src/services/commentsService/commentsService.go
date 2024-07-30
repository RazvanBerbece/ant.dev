package commentsService

import (
	"fmt"
	"log/slog"

	"github.com/RazvanBerbece/ant.dev/src/domain/models"
	"github.com/RazvanBerbece/ant.dev/src/repositories/articleComments"
)

type CommentsService struct {
	Logger                    slog.Logger
	articleCommentsRepository articleComments.ArticleCommentsDataRepository
}

func NewCommentsService(
	logger slog.Logger,
	articleCommentsRepository articleComments.ArticleCommentsDataRepository,
) CommentsService {
	return CommentsService{
		Logger:                    logger,
		articleCommentsRepository: articleCommentsRepository,
	}
}

func (s CommentsService) Get(articleId int) ([]models.ArticleComment, error) {

	commentsForArticle, err := s.articleCommentsRepository.Get(articleId)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, fmt.Errorf("could not retrieve comments for article with ID %d", articleId)
	}

	return commentsForArticle, nil
}

func (s CommentsService) Store(comment models.ArticleComment) error {

	err := s.articleCommentsRepository.Store(comment)
	if err != nil {
		s.Logger.Error(err.Error())
		return fmt.Errorf("could not store new comment for article with ID %d", comment.ArticleId)
	}

	go s.Logger.Info("Stored a new comment", "article_id", comment.ArticleId, "author", comment.Author)

	return nil
}
