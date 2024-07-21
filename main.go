package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/RazvanBerbece/ant.dev/publish"
	"github.com/RazvanBerbece/ant.dev/src/handlers"
	"github.com/RazvanBerbece/ant.dev/src/repositories/articleComments"
	"github.com/RazvanBerbece/ant.dev/src/services/articlesService"
	"github.com/RazvanBerbece/ant.dev/src/services/commentsService"
	"github.com/RazvanBerbece/ant.dev/src/startup"
	pages "github.com/RazvanBerbece/ant.dev/src/views/pages/main"
	"github.com/a-h/templ"
)

// Application configuration
var Environment = startup.NewEnvironment(
	os.Getenv("ENV"),
	os.Getenv("PORT"),
	os.Getenv("USE_LOCAL_STORAGE_COMMENTS"),
)

// Infrastructure
var Logger = slog.Default()

// Repositories
var ArticleCommentsRepository articleComments.ArticleCommentsDataRepository

// Services
var ArticlesService articlesService.ArticlesService
var CommentsService commentsService.CommentsService

func main() {

	InjectRuntimeDependencies(Environment, *Logger)

	// Serve static content (CSS, JS, images, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Simple page handlers
	http.Handle("/", templ.Handler(pages.Index(ArticlesService.GetRecentArticles(3))))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/contact", templ.Handler(pages.Contact()))

	// Complex page handlers
	http.HandleFunc("/articles", handlers.HandleArticleRequest(ArticlesService))  // GET /articles, GET /articles?id=
	http.HandleFunc("/comments", handlers.HandleCommentsRequest(CommentsService)) // GET /comments?articleId=, POST Form-Data /comments?articleId=

	Environment.LogStartupStatus(*Logger)

	// Blocking call! - Listen and serve
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Environment.Port), nil))
}

// Runtime injections and swaps
// (i.e use local or DB storage for comments under articles, use local list of articles as available publishings, other configurations, etc.)
func InjectRuntimeDependencies(env startup.Environment, logger slog.Logger) {

	ArticlesService = articlesService.NewLocalArticlesService(publish.PublishedArticles)

	if env.UsesLocalStorageForComments() {
		//
		// Use local runtime memory for storing & retrieving comments and other artifacts
		//
		ArticleCommentsRepository = articleComments.NewLocalArticleCommentsDataRepository(logger, ArticlesService.GetArticles())
	} else {
		//
		// Use a managed DB to store & retrieve comments and other artifacts
		//
		// TODO: Support storage of comments in a DB ? blob storage ?
		ArticleCommentsRepository = articleComments.NewManagedArticleCommentsDataRepository(logger, "NotSupportedYet")
	}

	CommentsService = commentsService.NewCommentsService(Logger, ArticleCommentsRepository)
}
