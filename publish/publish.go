package publish

import (
	"github.com/RazvanBerbece/ant.dev/publish/articles"
	"github.com/RazvanBerbece/ant.dev/src/domain"
)

//
// New articles are created with the following process :
// 	1. Write a Templ page component for the new article
// 	2. Add a new entry to the slice below, where the ID is an incrementing integer and the component is the newly created one from step 1.
//
// Note: This list dictates which routes for articles with a given ID actually exist,
// 		 so if a user requests an article with ID 99999, and that ID is not in this list,
// 		 then the server will respond with a 404 Not Found.
//

var PublishedArticles = []domain.Article{
	// Article 1
	{
		Id:        1,
		Title:     "1BRC: My Not-Very-Sophisticated C# Attempt",
		CreatedAt: 1720285638,
		TextPreview: `I'm generally not a wizard when it comes to squeezing every little ms of performance out of my code. 
		With that in mind, I went on an odyssey to create a potential solution to the 1 Billion Row Challenge using C# and .NET.`,
		Component: articles.Article__1BrcCsharp(1, 1720285638),
		Comments:  []domain.ArticleComment{},
	},
	// Article 2
	// etc...
}
