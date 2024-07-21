package articleComments

import (
	"github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type ArticleCommentsDataRepository interface {
	Get(articleId int) ([]models.ArticleComment, error)
	Store(comment models.ArticleComment) error
}
