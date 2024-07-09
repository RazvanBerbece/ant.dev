package commentsService

import "github.com/RazvanBerbece/ant.dev/src/domain/models"

type CommentsService interface {
	Get(articleId int) ([]models.ArticleComment, error)
	Store(comment models.ArticleComment) error
}
