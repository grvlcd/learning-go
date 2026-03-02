package main

import "time"

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Save()
