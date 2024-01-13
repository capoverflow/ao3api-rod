package models

import "time"

type FanficParams struct {
	Scheme    string
	Addr      string
	WorkID    string
	ChapterID string
	Debug     bool
	ProxyURLs []string
	Login     Login
}

type Work struct {
	Title           string     `bson:"title"`
	URL             string     `bson:"url"`
	WorkID          string     `bson:"workid"`
	Author          []string   `bson:"author"`
	DatePublished   time.Time  `bson:"datepublished"`
	Updated         time.Time  `bson:"updated"`
	Language        string     `bson:"language"`
	NBChapters      string     `bson:"nbchapters"`
	Words           string     `bson:"words"`
	Comments        string     `bson:"comments"`
	Kudos           string     `bson:"kudos"`
	Bookmarks       string     `bson:"bookmarks"`
	Hits            string     `bson:"hits"`
	RequiredTags    []string   `bson:"required"`
	Fandom          []string   `bson:"fandom"`
	Summary         []string   `bson:"summary"`
	Relationship    []string   `bson:"relationship"`
	AlternativeTags []string   `bson:"alternativetags"`
	Downloads       []Download `bson:"download"`
	Chapters        []Chapter  `bson:"chapters"`
}

type Download struct {
	Format string `bson:"format"`
	URL    string `bson:"url"`
}

type RequiredTags struct {
	ContentRating  string `bson:""`
	ContentWarning string `bson:""`
	Relationship   string `bson:""`
	WorkStatus     string `bson:""`
}
