package fields

import "github.com/ulferts/sudogo/structure"

type Fielded interface {
	Fields(*structure.Field) []*structure.Field
}
