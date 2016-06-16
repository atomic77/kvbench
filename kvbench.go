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
	"log"
	"fmt"
	"time"
	//"runtime"
	"flag"
)

type DatastoreTester interface {
	Init(host string, port int, user string, password string)
	InsertByPkRandom(start int, end int, t chan time.Duration)
	SelectByPkRandom(start int, end int, t chan time.Duration)
	UpdateByPkRandom(start int, end int, t chan time.Duration)
	DeleteByPkRandom(start int, end int, t chan time.Duration)
	CreateTables()
}

var numConnections = flag.Int("num-connections", 0, "Number of connections to DB")
var dbType = flag.String("db", "postgres", "Database type")
var testType = flag.String("test", "", "Test type: prepare, select-by-pk, etc.")
var numOpsPerConn = flag.Int("num-operations", 0, "Number of queries/updates/etc. per conn")
var host = flag.String("host", "192.168.42.223", "Target host")
var port = flag.Int("port", 0, "Port of db")
var user = flag.String("user", "u1", "Username if required")
var password = flag.String("password", "pw1pw1pw1", "Password if required")
var label = flag.String("label", "test", "Label to add to test result output")


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

	runTime := time.Now().Unix()
	tester.Init(*host, *port, *user, *password)
	if *testType == "insert-by-pk" {
		tester.CreateTables()
	}


	// A bit confusing : the first make creates a slice of channels of length *numCOnn
	// Then each individual make is there to create a new channel
	t := make([]chan time.Duration, *numConnections)
	for i := 0; i < *numConnections; i++ {
		t[i] = make(chan time.Duration)
	}

	for i := 0; i < *numConnections; i++ {
		start := *numOpsPerConn * i
		switch *testType {
		case "insert-by-pk":
			go tester.InsertByPkRandom(start, start + *numOpsPerConn, t[i])
		case "select-by-pk":
			go tester.SelectByPkRandom(start, start + *numOpsPerConn, t[i])
		case "update-by-pk":
			go tester.UpdateByPkRandom(start, start + *numOpsPerConn, t[i])
		case "delete-by-pk":
			go tester.DeleteByPkRandom(start, start + *numOpsPerConn, t[i])

		}
	}

	for i := 0; i< *numConnections; i++ {
		dur := <-t[i]
		rate := float64(*numOpsPerConn) / dur.Seconds()
		// Final output:
		// Runtime, DB, Test, ConnNum, Reqs, secs, rate
		fmt.Printf("%v,%v,%v,%v,%v,%v,%.2f,%.2f\n",
			runTime, *label, *dbType, *testType, i, *numOpsPerConn, dur.Seconds(), rate)
	}
}

func checkErr(err error, msg string, v ...interface{}) {
	if err != nil {
		log.Panic(msg, err, v)
	}
}

func assert(b bool, t string) {
	if !b {
		panic(t)
	}
}