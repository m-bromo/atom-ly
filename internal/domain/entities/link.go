package entities

import "time"

type Link struct {
	ShortCode string
	Url       string
	CreatedAt time.Time
}
