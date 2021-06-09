package domain

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	EmailId   string `json:"email_id"`
}
