package models

import "github.com/volatiletech/null/v8"

type Image struct {
	ID             uint
	Name           string
	Data           string
	UpdateDateTime string
	Tagid          null.Int
	Description    null.String
}
