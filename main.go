package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/RazvanBerbece/ant.dev/publish"
	"github.com/RazvanBerbece/ant.dev/src/handlers"
	"github.com/RazvanBerbece/ant.dev/src/services/articlesService"
	"github.com/RazvanBerbece/ant.dev/src/services/commentsService"
	pages "github.com/RazvanBerbece/ant.dev/src/views/pages/main"
	"github.com/a-h/templ"
)

// Environment Configuration
var Port = os.Getenv("PORT")
var UseLocalStorageForComments = os.Getenv("USE_LOCAL_STORAGE_COMMENTS")

// Infrastructure
var Logger = slog.Default()

// Services
var ArticlesService articlesService.ArticlesService
var CommentsService commentsService.CommentsService

func main() {

	// Runtime injections and swaps
	// i.e use Cloud dependencies when on Cloud Run,
	// and locals when running locally
	ArticlesService = articlesService.NewLocalArticlesService(publish.PublishedArticles)
	if UseLocalStorageForComments != "1" {
		// TODO: Support storage of comments on-Cloud
		CommentsService = commentsService.NewCloudCommentsService(Logger)
	} else {
		CommentsService = commentsService.NewLocalCommentsService(publish.PublishedArticles, Logger)
	}

	// Simple page handlers
	http.Handle("/", templ.Handler(pages.Index(ArticlesService.GetArticles())))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/contact", templ.Handler(pages.Contact()))

	// Complex page handlers
	http.HandleFunc("/articles", handlers.HandleArticleRequest(ArticlesService))  // GET /articles, GET /articles?id=
	http.HandleFunc("/comments", handlers.HandleCommentsRequest(CommentsService)) // GET /comments?articleId=, POST Form-Data /comments?articleId=

	if UseLocalStorageForComments != "1" {
		fmt.Printf("Serving on port: %s | Index: /\n | With Cloud storage\n", Port)
	} else {
		fmt.Printf("Serving on port: %s | Index: /\n | With local storage\n", Port)
	}

	// Blocking call! - Listen and serve
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Port), nil))
}
