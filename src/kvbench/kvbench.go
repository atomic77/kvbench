package main

import (
	"log"
	"fmt"
	"time"
	//"runtime"
	"flag"
)

type TestType int
const (
	Prepare TestType = iota
	Read
	Insert
	Update
	ReadWrite
)

type DatastoreTester interface {
	Init(host string, port int, user string, password string)
	Prepare(sz int, t chan time.Duration)
	SelectByPkRandom(start int, end int, t chan time.Duration)
	UpdateByPkRandom(start int, end int, t chan time.Duration)
}

var numConnections = flag.Int("num-connections", 0, "Number of connections to DB")
var dbType = flag.String("db", "postgres", "Database type")
var testType = flag.String("test", "", "Test type: prepare, select-by-pk, etc.")
var testSize = flag.Int("num-operations", 0, "Number of queries/updates/etc. per conn")
var host = flag.String("host", "192.168.42.223", "Target host")
var port = flag.Int("port", 0, "Port of db")
var user = flag.String("user", "u1", "Username if required")
var password = flag.String("password", "pw1pw1pw1", "Password if required")


func main() {
	flag.Parse()

	var tester DatastoreTester
	switch *dbType {
	case "postgres":
		var p Postgresql
		tester = DatastoreTester(&p)
	case "mysql":
		var m Mysql
		tester = DatastoreTester(&m)
	case "memcache":
		var m Memcache
		tester = DatastoreTester(&m)
	default:
		panic("Unsupported db type")
	}

	tester.Init(*host, *port, *user, *password)

	for i := 0; i < *numConnections; i++ {
		var dur time.Duration
		t := make(chan time.Duration)

		switch *testType {
		case "prepare":
			go tester.Prepare(*testSize, t)
		case "select-by-pk":
			go tester.SelectByPkRandom(1, *testSize, t)
		case "update-by-pk":
			go tester.UpdateByPkRandom(1, *testSize, t)

		}
		dur = <-t
		rate := float64(*testSize) / dur.Seconds()
		fmt.Printf("Routine called %v requests, %.2f sec, %.2f / s \n",
			*testSize, dur.Seconds(), rate)
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func assert(b bool, t string) {
	if !b {
		panic(t)
	}
}