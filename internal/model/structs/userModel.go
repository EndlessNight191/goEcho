package structs

type User struct {
	ID    string
	Login string `json:"Login"`
	Email string `json:"Email"`
}
