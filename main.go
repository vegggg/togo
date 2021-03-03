package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	sql_helper "github.com/manabie-com/togo/internal/storages/sql"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "user"
	password = "pass"
	dbname   = "tasks_service"
)

func newPsqlDB() (*sql.DB, sql_helper.Stmt) {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user, password, host, port, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	return db, sql_helper.Psql
}

func newSqlite3DB() (*sql.DB, sql_helper.Stmt) {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	return db, sql_helper.Sqlite
}

func main() {
	db, stmtConf := newSqlite3DB()
	defer db.Close()

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sql_helper.Helper{
			DB:   db,
			Stmt: stmtConf,
		},
	})
}
