package entity

type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Address   []string `json:"address"`
	Hobbies   []string `json:"hobbies"`
}
