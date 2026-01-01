package domain

// the complete novel data model
type Novel struct {
	Title    string
	Author   string
	Source   string
	Chapters []Chapter
}
