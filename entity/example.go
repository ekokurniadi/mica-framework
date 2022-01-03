package entity

import "time"

type Example struct {
	ID          int
	ExampleName string
	CreatedAt   time.Time
	Updated     time.Time
}
