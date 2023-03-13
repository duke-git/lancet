package structutil

import (
	"github.com/duke-git/lancet/v2/validator"
	"strings"
)

type Tag struct {
	Name    string
	Options []string
}

func newTag(tag string) *Tag {
	res := strings.Split(tag, ",")
	return &Tag{
		Name:    res[0],
		Options: res[1:],
	}
}

func (t *Tag) HasOption(opt string) bool {
	for _, o := range t.Options {
		if o == opt {
			return true
		}
	}
	return false
}

func (t *Tag) IsEmpty() bool {
	return validator.IsEmptyString(t.Name)
}
