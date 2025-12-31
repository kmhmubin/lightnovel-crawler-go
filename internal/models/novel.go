package models

// novel represent the base metadata of a book.
type Novel struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	CoverURL string `json:"cover_url"`

	// if a novel has a list of volumes
	Volumes string `json:"volumes"`
}

// vloume represent a collection of chapters
type Volume struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Chapters []Chapter `json:"chapters"`
}

// chapters represent a single realable story
type Chapter struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Body  string `json:"body"`
}
