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

// Configuration
var Port = os.Getenv("PORT")

// Infrastructure
var Logger = slog.Default()

// Services
var CommentsService = commentsService.NewLocalCommentsService(publishings.PublishedArticles, Logger)

func main() {

	// Runtime injections and swaps
	// e.g replace the CommentsService local strategy with a cloud one when in Cloud Run
	// TODO

	// Simple page handlers
	http.Handle("/", templ.Handler(pages.Index()))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/contact", templ.Handler(pages.Contact()))

	// Complex page handlers
	http.HandleFunc("/articles", handlers.HandleArticleRequest)                   // /articles, /articles?id=
	http.HandleFunc("/comments", handlers.HandleCommentsRequest(CommentsService)) // GET /comments?articleId=, POST Form-Data /comments?articleId=

	fmt.Printf("Listening on port: %s | Index: /\n", Port)

	// Blocking call! - Listen and serve
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Port), nil))
}
