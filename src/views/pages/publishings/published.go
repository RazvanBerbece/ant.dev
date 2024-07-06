package publishings

import (
	"github.com/RazvanBerbece/ant.dev/src/domain"
)

// New articles are created by first writing a Templ component for them.
// Then, a new entry is added in here, where the ID is an incrementing integer and the component is the new templ.Component.

var PublishedArticles = []domain.Article{
	// Article 1
	{
		Id:        1,
		Component: Article__1BrcCsharp("06/07/2024"),
	},
	// Article 2
	// etc...
}
