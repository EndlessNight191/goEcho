package structs

type Post struct {
	ID       string
	Title    string `json:"Title"`
	Content  string `json:"Content"`
	AuthorId int    `json:"author_id"`
	Image    string `json:"Image"`
}
