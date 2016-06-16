/*
Copyright 2016, Alex Tomic

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
	deleteByPk string
	commitInterval int
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
	m.deleteByPk = "DELETE FROM tab WHERE k = ?"
	m.commitInterval = 100

	/*
	For mysql, the DSN is formatted as:
	_connStr := "username:password@protocol(address)/dbname?param=value"
	 e.g: id:password@tcp(hostname.com:3306)/dbname
	 */

	_connStr := "%v:%v@tcp(%v:%v)/%v?readTimeout=%v&writeTimeout=%v"
	connStr := fmt.Sprintf(_connStr, user, password, host, port, "test",
		"2s", "2s")
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
	p.deleteByPk = "DELETE FROM tab WHERE k = $1"
	p.commitInterval = 100

	_connStr := "user=%v password=%v dbname=%v host=%v port=%v sslmode=disable connect_timeout=%v "
	connStr := fmt.Sprintf(_connStr, user, password, "postgres", host, port, "2")
	db, err := sql.Open("postgres", connStr)
	p.db = db

	checkErr(err, "Failed to connect to pgsql")

}

func (m *Mysql) CreateTables() {
	var err error

	_, err = m.db.Exec("CREATE SCHEMA IF NOT EXISTS test")
	checkErr(err, "Failed to create test schema")

	_, err = m.db.Exec(m.dropTable)
	checkErr(err, "Failed to create table")

	_, err = m.db.Exec(m.createTable)
	checkErr(err, "Failed to create table")
}

func (p *Postgresql) CreateTables() {
	var err error
	_, err = p.db.Exec(p.dropTable)
	checkErr(err, "Failed to create table")

	_, err = p.db.Exec(p.createTable)
	checkErr(err, "Failed to create table")
}

func (r *Rdbms) InsertByPkRandom(start int, end int, t chan time.Duration) {
	tm_s := time.Now()

	s, err := r.db.Prepare(r.insertInto)
	checkErr(err, "Failed to prepare stmt")

	l := end - start
	rnd := InitSampleSet(end - start, start)

	for i := 0; i < l; i++ {
		s.Exec(rnd[i], r.valStr)
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
		//fmt.Printf("k,v : %v , %v\n", k, v)
		checkErr(err2, "Failed on stmt exec")
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

func (r *Rdbms) DeleteByPkRandom(start int, end int, t chan time.Duration) {

	l := end - start
	rnd := InitSampleSet(end - start, start)

	tm_s := time.Now()
	s, err := r.db.Prepare(r.deleteByPk)
	checkErr(err, "Failed to query db")

	for i := 0; i < l; i++ {
		_, err2 := s.Exec(rnd[i])
		checkErr(err2, "Failed on stmt exec")
	}
	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}
