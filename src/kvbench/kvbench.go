package main

import (
	"log"
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
	Prepare(sz int)
	Select()
}

func main() {

	testers := []DatastoreTester{Postgresql{}}
	for _, tester := range testers {
		tester.Prepare(10)
	}

}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
