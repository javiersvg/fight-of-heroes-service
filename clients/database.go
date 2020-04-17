package clients

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func MySqlDatabaseFactory() *sql.DB {
	db,err :=sql.Open("mysql", "root:password@tcp(database:3306)/events")
	if err != nil {
		log.Fatal("unable to use data source name", err)
		panic(err)
	}
	return db
}