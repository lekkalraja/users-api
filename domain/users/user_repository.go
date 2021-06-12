package users

import (
	"github.com/lekkalraja/users-api/datasource/my_sql"
	"github.com/lekkalraja/users-api/utils"

	"github.com/lekkalraja/users-api/utils/date_utils"
)

const (
	userInsertionQuery string = "INSERT INTO users(first_name, last_name, email_id, date_created) VALUES (?, ?, ?, ?);"
	getAllUsers        string = "SELECT * FROM users;"
	getByUserId        string = "SELECT * FROM users where id = ?"
	deleteUser         string = "DELETE FROM users where id = ?"
	updateUser         string = "UPDATE users set first_name = ?, last_name = ?, email_id = ? WHERE id = ?"
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

func (u *User) Update(id int64) (int64, *utils.RestErr) {
	_, err := FindUser(id)
	if err != nil {
		return 0, err
	}
	stmt, err := my_sql.PrepareStatement(updateUser)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, updateErr := stmt.Exec(u.FirstName, u.LastName, u.EmailId, id)
	if updateErr != nil {
		return 0, my_sql.HandleError(updateErr)
	}

	affectedRows, idErr := res.RowsAffected()
	if idErr != nil {
		return 0, my_sql.HandleError(idErr)
	}
	return affectedRows, nil
}

func GetUsers() ([]*User, *utils.RestErr) {
	stmt, err := my_sql.PrepareStatement(getAllUsers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, getErr := stmt.Query()
	if getErr != nil {
		return nil, my_sql.HandleError(getErr)
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := &User{}
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.EmailId, &user.DateCreated)
		users = append(users, user)
	}

	return users, nil
}

func FindUser(id int64) (*User, *utils.RestErr) {
	stmt, err := my_sql.PrepareStatement(getByUserId)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	user := &User{}
	if getErr := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.EmailId, &user.DateCreated); getErr != nil {
		return nil, my_sql.HandleError(getErr)
	}

	return user, nil
}

func Delete(id int64) (int64, *utils.RestErr) {
	stmt, err := my_sql.PrepareStatement(deleteUser)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row, getErr := stmt.Exec(id)
	if getErr != nil {
		return -1, my_sql.HandleError(getErr)
	}

	affectedRows, rowsErr := row.RowsAffected()
	if rowsErr != nil {
		return -1, my_sql.HandleError(rowsErr)
	}

	return affectedRows, nil
}
