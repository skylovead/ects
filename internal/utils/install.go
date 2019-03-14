package utils

import (
	"database/sql"
	"fmt"
	"github.com/betterde/ects/config"
	"log"
)

var (
	DB *sql.DB
	err error
)

func IsDatabaseExist() bool {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=%s",
		config.Conf.Database.User,
		config.Conf.Database.Pass,
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.Char,
	)

	DB, err = sql.Open("mysql", dsn)

	//defer func() {
	//	if err := DB.Close(); err != nil {
	//		// TODO
	//	}
	//}()

	statement := fmt.Sprintf("SHOW DATABASES LIKE '%s'", config.Conf.Database.Name)

	var (
		rows *sql.Rows
		Database string
	)

	rows, err = DB.Query(statement)

	for rows.Next() {
		err := rows.Scan(&Database)
		if err != nil {
			// TODO
		}
	}

	if err != nil {
		// TODO
	}

	return Database == config.Conf.Database.Name
}

func CreateDatabase()  {
	var rows *sql.Rows
	statement := fmt.Sprintf("CREATE DATABASE %s", config.Conf.Database.Name)
	rows, err = DB.Query(statement)
	if err != nil {
		log.Println(err)
	}

	if err := rows.Close(); err != nil {
		//TODO
	}
}