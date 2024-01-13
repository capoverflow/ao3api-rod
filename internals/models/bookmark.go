package models

import "time"

type Bookmark struct {
	Title           string    `bson:"title"`
	URL             string    `bson:"url"`
	WorkID          string    `bson:"workid"`
	Author          []string  `bson:"author"`
	DateBookmarked  time.Time `bson:"datebookmarked"`
	DatePublished   time.Time `bson:"datepublished"`
	Updated         time.Time `bson:"updated"`
	Language        string    `bson:"language"`
	NBChapters      string    `bson:"nbchapters"`
	Words           string    `bson:"words"`
	Comments        int       `bson:"comments"`
	Kudos           int       `bson:"kudos"`
	Bookmarks       int       `bson:"bookmarks"`
	Hits            int       `bson:"hits"`
	RequiredTags    []string  `bson:"required"`
	Fandom          []string  `bson:"fandom"`
	Summary         []string  `bson:"summary"`
	Relationship    []string  `bson:"relationship"`
	AlternativeTags []string  `bson:"alternativetags"`
	Downloads       []string  `bson:"download"`
	Chapters        []Chapter `bson:"chapters"`
}
