package articlesService

import (
	"github.com/RazvanBerbece/ant.dev/src/domain"
	"github.com/a-h/templ"
)

// Retrieves an article Templ component from a local list of Articles.
// Usually the local variable is the runtime-available published articles slice.
func TryGetArticleComponentFromLocal(articles []domain.Article, id int) (bool, templ.Component) {

	if len(articles) == 0 {
		return false, nil
	}

	for _, article := range articles {
		if article.Id == id {
			return true, article.Component
		}
	}

	return false, nil

}
