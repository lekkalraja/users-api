package users

import (
	"strings"

	"github.com/lekkalraja/users-api/utils"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	EmailId     string `json:"email_id"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *utils.RestErr {
	fName := strings.Trim(user.FirstName, " ")
	lName := strings.Trim(user.LastName, " ")
	eId := strings.TrimSpace(user.EmailId)

	if len(fName) == 0 || len(lName) == 0 || len(eId) == 0 {
		return utils.NewBadRequest("Invalid Fname | LName | Emailid")
	}

	return nil
}
