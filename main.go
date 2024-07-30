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

// Default logging
var JsonHandler = slog.NewJSONHandler(os.Stdout, nil)
var DefaultLogger = slog.New(JsonHandler).With("layer", "Root")

// Infrastructure
// var Database = Database()...

// Repositories
var ArticleCommentsRepository articleComments.ArticleCommentsDataRepository

// Services
var ArticlesService articlesService.ArticlesService
var CommentsService commentsService.CommentsService

func main() {

	InjectRuntimeDependencies()

	// Serve static content (CSS, JS, images, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Simple page handlers
	http.Handle("/", templ.Handler(pages.Index(ArticlesService.GetRecentArticles(3))))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/contact", templ.Handler(pages.Contact()))

	// Complex page handlers
	http.HandleFunc("/articles", handlers.HandleArticleRequest(ArticlesService))  // GET /articles, GET /articles?id=
	http.HandleFunc("/comments", handlers.HandleCommentsRequest(CommentsService)) // GET /comments?articleId=, POST Form-Data /comments?articleId=

	Environment.LogStartupStatus(*DefaultLogger)

	// Blocking call! - Listen and serve
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Environment.Port), nil))
}

// Runtime injections and swaps
// (i.e use local or DB storage for comments under articles, use local list of articles as available publishings, other configurations, etc.)
func InjectRuntimeDependencies() {

	/////////////////////////////////////////////////////////////////
	// LOGGERS
	/////////////////////////////////////////////////////////////////

	// Data repository loggers
	var ArticleCommentsRepositoryLogger = slog.New(JsonHandler).With(
		"layer", "Repository",
		"repository_name", "ArticleCommentsRepository",
	)

	// Service loggers
	var ArticlesServiceLogger = slog.New(JsonHandler).With(
		"layer", "Service",
		"service_name", "ArticlesService",
	)
	var CommentsServiceLogger = slog.New(JsonHandler).With(
		"layer", "Service",
		"service_name", "CommentsService",
	)

	/////////////////////////////////////////////////////////////////
	// DATA REPOSITORIES & SERVICES
	/////////////////////////////////////////////////////////////////

	ArticlesService = articlesService.NewLocalArticlesService(*ArticlesServiceLogger, publish.PublishedArticles)

	if Environment.UsesLocalStorageForComments() {
		//
		// Use local runtime memory for storing & retrieving comments and other artifacts
		//
		ArticleCommentsRepository = articleComments.NewLocalArticleCommentsDataRepository(*ArticleCommentsRepositoryLogger, ArticlesService.GetArticles())
	} else {
		//
		// Use a managed DB to store & retrieve comments and other artifacts
		//
		// TODO: Support storage of comments in a DB ? blob storage ?
		ArticleCommentsRepository = articleComments.NewManagedArticleCommentsDataRepository(*ArticleCommentsRepositoryLogger, "NotSupportedYet")
	}

	CommentsService = commentsService.NewCommentsService(*CommentsServiceLogger, ArticleCommentsRepository)
}
