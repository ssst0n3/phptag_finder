package phptag_finder

import (
	"github.com/z7zmey/php-parser/position"
)

type Tag struct {
	Position *position.Position
}

func NewTag(position *position.Position) Tag {
	return Tag{
		Position: position,
	}
}
