package models

import "time"

type Employee struct {
	ID             uint
	Name           string
	Picture        []byte
	UpdateDateTime time.Time
	Department     string
	Remarks        string
}
