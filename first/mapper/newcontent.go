package mapper

import (
	. "rdani2005.com/GolangAdventageCode/first/entities"
)

func NewTextContent(texts []string) *TextContent {
	return &TextContent{
		Texts:   texts,
		Results: nil,
	}
}
