package articleComments

import (
	"fmt"
	"log/slog"

	"github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type LocalArticleCommentsDataRepository struct {
	Logger                    slog.Logger
	LocalArticlesWithComments []models.Article
}

func NewLocalArticleCommentsDataRepository(logger slog.Logger, localArticlesWithComments []models.Article) LocalArticleCommentsDataRepository {
	return LocalArticleCommentsDataRepository{
		Logger:                    logger,
		LocalArticlesWithComments: localArticlesWithComments,
	}
}

func (r LocalArticleCommentsDataRepository) Get(articleId int) ([]models.ArticleComment, error) {

	for _, article := range r.LocalArticlesWithComments {
		if article.Id == articleId {
			return article.Comments, nil
		}
	}

	r.Logger.Info("comment failed to retrieve comments (article with ID not found)", slog.Int("ArticleId", articleId))

	return nil, fmt.Errorf("an error ocurred while reading the comments: article with ID %d does not exist", articleId)
}

func (r LocalArticleCommentsDataRepository) Store(comment models.ArticleComment) error {

	for idx, article := range r.LocalArticlesWithComments {
		if article.Id == comment.ArticleId {
			r.LocalArticlesWithComments[idx].Comments = append(r.LocalArticlesWithComments[idx].Comments, comment)
			return nil
		}
	}

	r.Logger.Info("comment failed to store (article with ID not found)", slog.Int("ArticleId", comment.ArticleId))

	return fmt.Errorf("an error ocurred while storing the comment: article with ID %d does not exist", comment.ArticleId)
}
