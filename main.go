package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RazvanBerbece/ant.dev/src/handlers"
	pages "github.com/RazvanBerbece/ant.dev/src/views/pages/main"
	"github.com/a-h/templ"
)

var Port = os.Getenv("PORT")

func main() {

	// Simple page handlers
	http.Handle("/", templ.Handler(pages.Index()))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/contact", templ.Handler(pages.Contact()))

	// Complex page handlers
	http.HandleFunc("/articles", handlers.HandleArticleRequest) // /articles, /articles?id=

	fmt.Printf("Listening on port: %s | Index: /\n", Port)

	// Blocking call! - Listen and serve
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", Port), nil))
}
