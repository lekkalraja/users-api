package my_sql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lekkalraja/users-api/utils"
)

const (
	MYSQL_USER     string = "MYSQL_USER"
	MYSQL_PASSWORD string = "MYSQL_PASSWORD"
	MYSQL_HOST     string = "MYSQL_HOST"
	MYSQL_DB       string = "MYSQL_DB"
)

var Client *sql.DB

func init() {
	user := os.Getenv(MYSQL_USER)
	pwd := os.Getenv(MYSQL_PASSWORD)
	host := os.Getenv(MYSQL_HOST)
	db := os.Getenv(MYSQL_DB)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pwd, host, db)

	var err error

	Client, err = sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	if connErr := Client.Ping(); connErr != nil {
		log.Panicf("Failed to Ping : %v", connErr)
	}
	log.Printf("Successfully Established Connection")
}

func PrepareStatement(query string) (*sql.Stmt, *utils.RestErr) {
	stmt, err := Client.Prepare(query)
	if err != nil {
		return nil, HandleError(err)
	}
	return stmt, nil
}

func HandleError(err error) *utils.RestErr {
	mysqlError, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return utils.NewBadRequest("Requested User Didn't found")
		}
		return utils.InternalServerError(err.Error())
	}

	switch mysqlError.Number {
	case 1146:
		return utils.InternalServerError("Table Not Found")
	case 1062:
		return utils.InternalServerError("User Already Exist")
	default:
		return utils.InternalServerError(err.Error())
	}
}
