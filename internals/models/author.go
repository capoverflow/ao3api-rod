package models

type AuthorParams struct {
	Addr           string
	Author         string
	Debug          bool
	WorkPageSpleep int
	Bookmarks      bool
	ProxyURLs      []string
	Login          Login
}

type Author struct {
	Author    string
	Profile   AuthorProfile
	Works     []Work
	Bookmarks []Bookmark
	Gift      []Work
	Series    []Series
}

type AuthorProfile struct {
	Pseuds   []string
	JoinDate string
	Bio      string
}
