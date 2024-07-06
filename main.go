package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RazvanBerbece/ant.dev/src/views/pages"
	"github.com/a-h/templ"
)

var Port = 8080

func main() {

	http.Handle("/", templ.Handler(pages.Index()))
	// http.Handle("/articles", templ.Handler(pages.Articles()))
	http.Handle("/about", templ.Handler(pages.About()))
	http.Handle("/contact", templ.Handler(pages.Contact()))

	fmt.Printf("Listening on port: %d | Index: /\n", Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)) // blocking call
}
