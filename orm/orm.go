package orm

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pitakill/go_api/config"
)

func Connection() *sql.DB {
	config := config.GetConfig()

	params := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.IP, config.Port, config.Database)

	db, err := sql.Open("mysql", params)
	if err != nil {
		fmt.Print(err.Error())
	}

	err = db.Ping()

	if err != nil {
		fmt.Print(err.Error())
	}

	return db
}
