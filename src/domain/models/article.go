package models

import "github.com/a-h/templ"

type Article struct {
	Id    int
	Title string

	// The Templ component rendered for an article model
	Component templ.Component

	// UNIX timestamp for when the article was created
	CreatedAt int64

	// A preview of the content of the article
	TextPreview string

	// List of tags associated with the article
	Tags []ArticleTag

	// The URL or statics path which points to the thumbnail image of the article
	ThumbnailImagePath string

	// Alt value to use for the thumbnail image
	ThumbnailImageAlt string

	// Locally-stored list of comments associated with an article (TODO: redundant once cloud storage is supported)
	Comments []ArticleComment
}
