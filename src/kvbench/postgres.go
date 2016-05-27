package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

type Postgresql struct {
	db *sql.DB
}

const CREATE_TABLE = "CREATE TABLE huy (k INT, v TEXT, PRIMARY KEY (k))"
const INSERT_INTO = "INSERT INTO huy VALUES ($1, $2)"
const CONN_STR = "user=%v dbname=%v host=%v sslmode=disable"

func (p Postgresql) Init() {
	var err error

	connStr := fmt.Sprintf(CONN_STR, "postgres", "postgres", "192.168.42.223")
	p.db, err = sql.Open("postgres", connStr)

	checkErr(err, "Failed to connect to pgsql")
}

func (p Postgresql) Prepare(sz int) {
	fmt.Printf("Called pgsql.prepare with i = %d\n", sz)
}

func (p Postgresql) Select() {

	//rows, err := db.Query("SELECT k,v FROM huy WHERE k = $1", key)
	rows, err := p.db.Query("SELECT k,v FROM huy ")
	checkErr(err, "Failed to query db")

	defer rows.Close()
	for rows.Next() {
		var k int
		var v string
		if err := rows.Scan(&k, &v); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("k,v : %d : %v\n", k, v)
	}

}

