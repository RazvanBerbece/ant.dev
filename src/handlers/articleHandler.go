package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RazvanBerbece/ant.dev/src/domain"
	errPages "github.com/RazvanBerbece/ant.dev/src/views/pages/err"
	pages "github.com/RazvanBerbece/ant.dev/src/views/pages/main"
	"github.com/RazvanBerbece/ant.dev/src/views/pages/publishings"
	"github.com/a-h/templ"
)

func HandleArticleRequest(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()) == 0 {
		// Index page for Articles
		pages.Articles().Render(r.Context(), w)
	} else {
		// A specific article is requested by ID
		articleId := r.URL.Query().Get("id")

		if articleId != "" {

			idAsInt, err := strconv.Atoi(articleId)
			if err != nil {
				// Invalid ID for an Article
				errPages.ErrGeneric(fmt.Sprintf("Can't request an article with an invalid ID: %s%s?id=%s", r.Host, r.URL.Path, articleId), 400).Render(r.Context(), w)
				return
			}

			// Get the article component by ID from local memory
			articleExists, articleComponent := TryGetArticleFromMemory(idAsInt)

			if !articleExists {
				// Error - Non-existent article
				errPages.ErrNotFound(fmt.Sprintf("Article with ID: %s", articleId)).Render(r.Context(), w)
			} else {
				articleComponent.Render(r.Context(), w)
			}
		} else {
			// Error - Article with empty ID
			errPages.ErrGeneric(fmt.Sprintf("Can't request an article with an empty ID: %s%s?id=", r.Host, r.URL.Path), 400).Render(r.Context(), w)
		}
	}
}

func TryGetArticleFromMemory(id int) (bool, templ.Component) {

	exists, article := ArticleWithIdExists(id, publishings.PublishedArticles)

	if !exists {
		return false, nil
	}

	return true, article.Component

}

func ArticleWithIdExists(id int, slice []domain.Article) (bool, *domain.Article) {
	if len(slice) == 0 {
		return false, nil
	}

	for _, article := range slice {
		if article.Id == id {
			return true, &article
		}
	}

	return false, nil
}