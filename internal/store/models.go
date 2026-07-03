package store

import "time"

type Collection struct {
	ID        int
	Title     string
	CreatedAt time.Time
	LinkCount int // only populated when listing with counts
}

type Link struct {
	ID        int
	URL       string
	Tag       string
	CreatedAt time.Time
}
