package models

import "time"

type Chapter struct {
	ChapterID string    `bson:"chapter_id"`
	Name      string    `bson:"chapter_name"`
	Date      time.Time `bson:"date_posted"`
}
