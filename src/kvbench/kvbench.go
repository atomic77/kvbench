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
	Init()
	Prepare(sz int, t chan time.Duration)
	SelectByPkRandom(start int, end int, t chan time.Duration)
	UpdateByPkRandom(start int, end int, t chan time.Duration)
}

var numConnections = flag.Int("num-connections", 0, "Number of connections to DB")
var dbType = flag.String("db", "postgres", "Database type")
var testType = flag.String("test", "", "Test type: prepare, select-by-pk, etc.")
var testSize = flag.Int("num-operations", 0, "Number of queries/updates/etc. per conn")


func main() {
	flag.Parse()

	var tester DatastoreTester

	if *dbType == "postgres" {
		var p Postgresql
		tester = DatastoreTester(&p)
	} else {
		panic("Unsupported db type")
	}

	tester.Init()

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