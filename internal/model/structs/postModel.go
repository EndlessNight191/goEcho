package structs

type Post struct {
	Title    string `json:"Title"`
	Content  string `json:"Content"`
	AuthorId int    `json:"AuthorId"`
	Image    string `json:"Image"`
}
