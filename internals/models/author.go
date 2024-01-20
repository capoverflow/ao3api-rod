package models

type AuthorParams struct {
	Addr     string
	Author   string
	AllWorks bool
}

type Author struct {
	Author       string
	AuthorParams AuthorParams
	Profile      AuthorProfile
	Fandoms      []AuthorFandom
	Works        []Work
	Bookmarks    []Bookmark
	Gift         []Work
	Series       []Series
}

type AuthorFandom struct {
	Fandom        string
	URL           string
	NumberOfWorks int
}

type AuthorProfile struct {
	Pseuds   []string
	JoinDate string
	Bio      string
}
