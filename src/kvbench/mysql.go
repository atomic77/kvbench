package main

import (
	"fmt"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct { }

func (m Mysql) Prepare(i int) {
	fmt.Printf("Called mysql.prepare with i = %d\n", i)
}

func (m Mysql) Select() {
	//db, err := sql.Open("mysql", "user:password@/dbname")
}
