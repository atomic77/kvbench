package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Postgresql struct {
	db *sql.DB
}

const CONN_STR = "user=%v dbname=%v host=%v sslmode=disable"

func (p *Postgresql) Init() {

	connStr := fmt.Sprintf(CONN_STR, "postgres", "postgres", "192.168.42.223")
	db, err := sql.Open("postgres", connStr)
	p.db = db

	checkErr(err, "Failed to connect to pgsql")
}

const CREATE_TABLE = "CREATE TABLE tab (k INT, v TEXT, PRIMARY KEY (k))"
const DROP_TABLE = "DROP TABLE IF EXISTS tab"
const INSERT_INTO = "INSERT INTO tab VALUES ($1, $2)"
const VAL_STR = "1234567890ABCDEF"

func (p *Postgresql) Prepare(sz int, t chan time.Duration) {
	var err error
	tm_s := time.Now()
	_, err = p.db.Exec(DROP_TABLE)
	checkErr(err, "Failed to create table")

	_, err = p.db.Exec(CREATE_TABLE)
	checkErr(err, "Failed to create table")

	s, err2 := p.db.Prepare(INSERT_INTO)
	checkErr(err2, "Failed to prepare stmt")

	for i := 1; i<=sz ; i++  {
		s.Exec(i, VAL_STR)
	}

	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}

const SELECT_BY_PK = "SELECT k, v FROM tab WHERE k = $1"

func (p *Postgresql) SelectByPkRandom(start int, end int, t chan time.Duration) {

	l := end - start
	r := InitSampleSet(end - start, start)

	tm_s := time.Now()
	s, err := p.db.Prepare(SELECT_BY_PK)
	checkErr(err, "Failed to query db")

	for i := 0; i < l; i++ {
		var k int
		var v string
		err2 := s.QueryRow(r[i]).Scan(&k, &v)
		checkErr(err2, "Failed on stmt exec")
		//fmt.Printf("k,v : %v , %v\n", k, v)
	}
	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}

const UPDATE_BY_PK = "UPDATE tab SET v = $1 WHERE k = $2"

func (p *Postgresql) UpdateByPkRandom(start int, end int, t chan time.Duration) {

	l := end - start
	r := InitSampleSet(end - start, start)

	tm_s := time.Now()
	s, err := p.db.Prepare(UPDATE_BY_PK)
	checkErr(err, "Failed to query db")

	v := "00012312301230"

	for i := 0; i < l; i++ {
		_, err2 := s.Exec(v, r[i])
		checkErr(err2, "Failed on stmt exec")
	}
	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}
