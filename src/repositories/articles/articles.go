package articles

import (
	"github.com/RazvanBerbece/ant.dev/src/domain/models"
)

type ArticlesDataRepository interface {
	GetAll() ([]models.Article, error)
	GetById(articleId int) (*models.Article, error)
}
