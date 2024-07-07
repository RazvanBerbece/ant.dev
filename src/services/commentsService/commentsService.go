package commentsService

import "github.com/RazvanBerbece/ant.dev/src/domain"

type CommentsService interface {
	Get(articleId int) ([]domain.ArticleComment, error)
	Store(comment domain.ArticleComment) error
}
