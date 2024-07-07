package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RazvanBerbece/ant.dev/src/domain"
	"github.com/RazvanBerbece/ant.dev/src/domain/contracts"
	"github.com/RazvanBerbece/ant.dev/src/services/commentsService"
	"github.com/RazvanBerbece/ant.dev/src/views/components"
	errPages "github.com/RazvanBerbece/ant.dev/src/views/pages/err"
)

func HandleCommentsRequest(commentsService commentsService.CommentsService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var targetArticleId int

		// Validate and extract query params (i.e article ID)
		if len(r.URL.Query()) == 0 {
			// Can't retrieve comments without an article id
			// so redirect to index page
			indexFullUrl := fmt.Sprintf("%v://%v/", GetHttpRequestScheme(r), r.Host)
			http.Redirect(w, r, indexFullUrl, http.StatusSeeOther)
			return
		} else {
			// Comments requested for an article with a specified ID
			articleId := r.URL.Query().Get("articleId")

			if articleId != "" {

				idAsInt, err := strconv.Atoi(articleId)
				if err != nil {
					// Invalid ID for an Article
					w.WriteHeader(http.StatusBadRequest)
					errPages.ErrGeneric(fmt.Sprintf("Can't request comments for an article with an invalid ID: %s%s?articleId=%s", r.Host, r.URL.Path, articleId), http.StatusBadRequest).Render(r.Context(), w)
					return
				}

				targetArticleId = idAsInt
			} else {
				// Error - Article with empty ID
				w.WriteHeader(http.StatusBadRequest)
				errPages.ErrGeneric(fmt.Sprintf("Can't request comments for an article with an empty ID: %s%s?articleId=", r.Host, r.URL.Path), http.StatusBadRequest).Render(r.Context(), w)
			}
		}

		switch r.Method {
		case "GET":
			comments, err := commentsService.Get(targetArticleId)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				errPages.ErrGeneric(fmt.Sprintf("An error ocurred while retrieving comments for article with ID %d: %v", targetArticleId, err), http.StatusInternalServerError).Render(r.Context(), w)
				return
			}

			components.CommentsUnderSection(comments).Render(r.Context(), w)
		case "POST":
			err := r.ParseForm()
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Web form parse failure: %v", err)
				return
			}

			var dto contracts.PostCommentRequest
			if r.Form.Has("comment_content") {
				dto.Content = r.FormValue("comment_content")
			}
			if r.Form.Has("comment_author") {
				dto.Author = r.FormValue("comment_author")
			}

			err = commentsService.Store(domain.ArticleComment{
				Author:    dto.Author,
				Content:   dto.Content,
				ArticleId: targetArticleId,
				CreatedAt: time.Now().Unix(),
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				components.DefaultOperationStatus("failed to submit comment", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			components.DefaultOperationStatus("comment submitted succesfully!", http.StatusCreated).Render(r.Context(), w)
		default:
			fmt.Printf("HTTP method %s not supported for /comments handler", r.Method)
			return
		}
	}
}

func GetHttpRequestScheme(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme
}
