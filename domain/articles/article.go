package articles

// Article defines the article struct
type Article struct {
	ID         int64    `json:"id"`
	Title      string   `json:"title"`
	Blurb      string   `json:"blurb"`
	Content    string   `json:"content"`
	Categories []string `json:"categories"`
}
