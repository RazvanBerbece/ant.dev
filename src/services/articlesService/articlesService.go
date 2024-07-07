package articlesService

import (
	"github.com/RazvanBerbece/ant.dev/src/domain"
	"github.com/a-h/templ"
)

// Retrieves an article Templ component from the runtime-available published articles.
func TryGetArticleComponentFromMemory(articles []domain.Article, id int) (bool, templ.Component) {

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

// Retrieves an article from the runtime-available published articles.
func TryGetArticleFromMemory(articles []domain.Article, id int) (bool, *domain.Article) {

	if len(articles) == 0 {
		return false, nil
	}

	for _, article := range articles {
		if article.Id == id {
			return true, &article
		}
	}

	return false, nil

}
