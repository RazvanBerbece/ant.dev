package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RazvanBerbece/ant.dev/src/services/articlesService"
	errPages "github.com/RazvanBerbece/ant.dev/src/views/pages/err"
	pages "github.com/RazvanBerbece/ant.dev/src/views/pages/main"
	"github.com/RazvanBerbece/ant.dev/src/views/pages/publishings"
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
				w.WriteHeader(http.StatusBadRequest)
				errPages.ErrGeneric(fmt.Sprintf("Can't request an article with an invalid ID: %s%s?id=%s", r.Host, r.URL.Path, articleId), http.StatusBadRequest).Render(r.Context(), w)
				return
			}

			// Get the article component by ID from local memory
			articleExists, articleComponent := articlesService.TryGetArticleComponentFromMemory(publishings.PublishedArticles, idAsInt)

			if !articleExists {
				// Error - Non-existent article
				w.WriteHeader(http.StatusNotFound)
				errPages.ErrNotFound(fmt.Sprintf("Article with ID: %s", articleId)).Render(r.Context(), w)
			} else {
				articleComponent.Render(r.Context(), w)
			}
		} else {
			// Error - Article with empty ID
			w.WriteHeader(http.StatusNotFound)
			errPages.ErrGeneric(fmt.Sprintf("Can't request an article with an empty ID: %s%s?id=", r.Host, r.URL.Path), http.StatusNotFound).Render(r.Context(), w)
		}
	}
}
