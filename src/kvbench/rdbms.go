package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Rdbms struct {
	db *sql.DB
	createTable string
	dropTable string
	insertInto string
	valStr string
	selectByPk string
	updateByPk string
}

type Postgresql struct {
	Rdbms
}

type Mysql struct {
	Rdbms
}

func (m *Mysql) Init(host string, port int, user string, password string) {

	m.createTable = "CREATE TABLE tab (k INT, v TEXT, PRIMARY KEY (k)) ENGINE = InnoDB"
	m.dropTable = "DROP TABLE IF EXISTS tab"
	m.insertInto = "INSERT INTO tab VALUES (?, ?)"
	m.valStr = "1234567890ABCDEF"
	m.selectByPk = "SELECT k, v FROM tab WHERE k = ?"
	m.updateByPk = "UPDATE tab SET v = ? WHERE k = ?"

	/*
	For mysql, the DSN is formatted as:
	_connStr := "username:password@protocol(address)/dbname?param=value"
	 e.g: id:password@tcp(hostname.com:3306)/dbname
	 */

	_connStr := "%v:%v@tcp(%v:%v)/%v"
	connStr := fmt.Sprintf(_connStr, user, password, host, port, "test")
	db, err := sql.Open("mysql", connStr)
	m.db = db
	checkErr(err, "Failed to connect to mysql")
}


func (p *Postgresql) Init(host string, port int, user string, password string) {

	p.createTable = "CREATE TABLE tab (k INT, v TEXT, PRIMARY KEY (k))"
	p.dropTable = "DROP TABLE IF EXISTS tab"
	p.insertInto = "INSERT INTO tab VALUES ($1, $2)"
	p.valStr = "1234567890ABCDEF"
	p.selectByPk = "SELECT k, v FROM tab WHERE k = $1"
	p.updateByPk = "UPDATE tab SET v = $1 WHERE k = $2"

	_connStr := "user=%v password=%v dbname=%v host=%v port=%v sslmode=disable"
	connStr := fmt.Sprintf(_connStr, user, password, "postgres", host, port)
	db, err := sql.Open("postgres", connStr)
	p.db = db
	checkErr(err, "Failed to connect to pgsql")
}


func (r *Rdbms) Prepare(sz int, t chan time.Duration) {
	var err error
	tm_s := time.Now()
	_, err = r.db.Exec(r.dropTable)
	checkErr(err, "Failed to create table")

	_, err = r.db.Exec(r.createTable)
	checkErr(err, "Failed to create table")

	s, err2 := r.db.Prepare(r.insertInto)
	checkErr(err2, "Failed to prepare stmt")

	for i := 1; i<=sz ; i++  {
		s.Exec(i, r.valStr)
	}

	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}


func (r *Rdbms) SelectByPkRandom(start int, end int, t chan time.Duration) {

	l := end - start
	rnd := InitSampleSet(end - start, start)

	tm_s := time.Now()
	s, err := r.db.Prepare(r.selectByPk)
	checkErr(err, "Failed to query db")

	for i := 0; i < l; i++ {
		var k int
		var v string
		err2 := s.QueryRow(rnd[i]).Scan(&k, &v)
		checkErr(err2, "Failed on stmt exec")
		//fmt.Printf("k,v : %v , %v\n", k, v)
	}
	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}


func (r *Rdbms) UpdateByPkRandom(start int, end int, t chan time.Duration) {

	l := end - start
	rnd := InitSampleSet(end - start, start)

	tm_s := time.Now()
	s, err := r.db.Prepare(r.updateByPk)
	checkErr(err, "Failed to query db")

	v := "00012312301230"

	for i := 0; i < l; i++ {
		_, err2 := s.Exec(v, rnd[i])
		checkErr(err2, "Failed on stmt exec")
	}
	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}
