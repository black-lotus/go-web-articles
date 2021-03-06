package domain

import "time"

// Article structure
type Article struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	IsDeleted bool      `json:"is_deleted"`
	Created   time.Time `json:"created"`
}

// Filter model
type Filter struct {
	Query  string `json:"query"`
	Author string `json:"author"`
}
