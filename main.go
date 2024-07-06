package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RazvanBerbece/ant.dev/src/handlers"
	pages "github.com/RazvanBerbece/ant.dev/src/views/pages/main"
	"github.com/a-h/templ"
)

var Port = 8080

func main() {

	// Simple page handlers
	http.Handle("/", templ.Handler(pages.Index()))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/contact", templ.Handler(pages.Contact()))

	// Complex page handlers
	http.HandleFunc("/articles", handlers.HandleArticleRequest)

	fmt.Printf("Listening on port: %d | Index: /\n", Port)

	// Blocking call! - Listen and server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
}
