package models

import "fmt"

type ArticleTag struct {
	Value      string
	BgColour   string
	TextColour string
}

func (t ArticleTag) GetTailwindStyleString() string {
	return fmt.Sprintf("%s %s", t.BgColour, t.TextColour)
}
