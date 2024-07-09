package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/RazvanBerbece/ant.dev/src/handlers"
	"github.com/RazvanBerbece/ant.dev/src/services/commentsService"
	pages "github.com/RazvanBerbece/ant.dev/src/views/pages/main"
	"github.com/RazvanBerbece/ant.dev/src/views/pages/publishings"
	"github.com/a-h/templ"
)

// Environment Configuration
var Port = os.Getenv("PORT")
var UsesLocalStorage = os.Getenv("USE_LOCAL_STORAGE")

// Infrastructure
var Logger = slog.Default()

// Services
var CommentsService commentsService.CommentsService

func main() {

	// Runtime injections and swaps
	// i.e use Cloud dependencies when on Cloud Run,
	// and locals when running locally
	if UsesLocalStorage != "1" {
		CommentsService = commentsService.NewCloudCommentsService(Logger)
	} else {
		CommentsService = commentsService.NewLocalCommentsService(publishings.PublishedArticles, Logger)
	}

	// Simple page handlers
	http.Handle("/", templ.Handler(pages.Index()))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/contact", templ.Handler(pages.Contact()))

	// Complex page handlers
	http.HandleFunc("/articles", handlers.HandleArticleRequest(publishings.PublishedArticles)) // /articles, /articles?id=
	http.HandleFunc("/comments", handlers.HandleCommentsRequest(CommentsService))              // GET /comments?articleId=, POST Form-Data /comments?articleId=

	fmt.Printf("Listening on port: %s | Index: /\n", Port)

	// Blocking call! - Listen and serve
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Port), nil))
}
