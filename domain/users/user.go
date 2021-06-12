package users

import (
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	EmailId     string `json:"email_id"`
	DateCreated string `json:"date_created"`
}

func (user *User) Format() {
	user.FirstName = strings.Trim(user.FirstName, " ")
	user.LastName = strings.Trim(user.LastName, " ")
	user.EmailId = strings.TrimSpace(user.EmailId)
}
