package domain

import "time"

// Article structure
type Article struct {
	ID      int64     `json:"id"`
	Author  string    `json:"author"`
	Title   int64     `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

// Filter model
type Filter struct {
	Query  string `json:"query"`
	Author string `json:"author"`
}
