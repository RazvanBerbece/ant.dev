package domain

import "github.com/a-h/templ"

type Article struct {
	Id          int
	Title       string
	Component   templ.Component // the Templ component rendered for an article model
	CreatedAt   int64           // UNIX timestamp for when the article was created
	TextPreview string          // a preview of the content of the article
}
