package users

import (
	"github.com/lekkalraja/users-api/datasource/my_sql"
	"github.com/lekkalraja/users-api/utils"

	"github.com/lekkalraja/users-api/utils/date_utils"
)

const (
	userInsertionQuery string = "INSERT INTO users(first_name, last_name, email_id, date_created) values (?, ?, ?, ?);"
)

var (
	userDB = make(map[int64]*User)
)

func (u *User) Save() *utils.RestErr {
	stmt, err := my_sql.PrepareStatement(userInsertionQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	u.DateCreated = date_utils.GetNowString()
	res, saveErr := stmt.Exec(u.FirstName, u.LastName, u.EmailId, u.DateCreated)
	if saveErr != nil {
		return my_sql.HandleError(saveErr)
	}

	id, idErr := res.LastInsertId()
	if idErr != nil {
		return my_sql.HandleError(idErr)
	}

	u.Id = id
	return nil
}

func GetUsers() []*User {
	users := make([]*User, 0, len(userDB))

	for _, value := range userDB {
		users = append(users, value)
	}

	return users
}

func FindUser(id int64) (*User, *utils.RestErr) {
	user, ok := userDB[id]
	if !ok {
		return nil, utils.UserNotFound(id)
	}
	return user, nil
}
