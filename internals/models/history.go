package models

import "time"

type History struct {
	Title           string    `bson:"title"`
	URL             string    `bson:"url"`
	WorkID          string    `bson:"workid"`
	Author          []string  `bson:"author"`
	DateVisited     time.Time `bson:"datevisited"`
	VisitNumber     int       `bson:"numbervisited"`
	VisitVersion    string    `bson:"visitversion"`
	DatePublished   time.Time `bson:"datepublished"`
	Updated         time.Time `bson:"updated"`
	Language        string    `bson:"language"`
	NBChapters      string    `bson:"nbchapters"`
	CurrentChapters string    `bson:"currentchapters"`
	MaxChapters     string    `bson:"maxchapters"`
	WordsInt        int       `bson:"wordsint"`
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
