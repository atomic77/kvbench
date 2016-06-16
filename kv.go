package main

// Testers for the K-V stores we will compare against - memcached for now

import (
	"github.com/bradfitz/gomemcache/memcache"
	"time"
	"strconv"
	"fmt"
)


type Memcache struct {

	mc *memcache.Client
	valStr string
}

func (m *Memcache) Init(host string, port int, user string, password string) {
	_connStr := "%v:%v"
	connStr := fmt.Sprintf(_connStr, host, port)
	m.mc = memcache.New(connStr)
	// We can get i/o timeouts with parallel executions without this
	m.mc.Timeout = time.Second * 2
	m.valStr = "1234567890ABCDEF"
}

func (m *Memcache) CreateTables() {
	err := m.mc.DeleteAll()
	checkErr(err, "Error purging memcache for prepare")

}
func (m *Memcache) InsertByPkRandom(start int, end int, t chan time.Duration) {

	// There are no 'tables' to set up in memcached so just pre-write the keys
	var err error
	tm_s := time.Now()
	for i := start; i< end ; i++  {
		k := strconv.Itoa(i)
		err = m.mc.Add(&memcache.Item{Key: k, Value: []byte(m.valStr)})
		checkErr(err, "Failed to add data to memcache: ", k)
	}

	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}

func (m *Memcache) SelectByPkRandom(start int, end int, t chan time.Duration) {
	l := end - start
	rnd := InitSampleSet(end - start, start)

	tm_s := time.Now()

	for i := 0; i < l; i++ {
		var k string
		// float -> int -> string, yay
		k = strconv.Itoa(int(rnd[i]))
		_, err := m.mc.Get(k)
		checkErr(err, "Failed on stmt exec for k/v: ", k)
		//v := string(item.Value)
		//assert(v == m.valStr, "Value string not same")
		//fmt.Printf("k,v : %v , %v\n", k, v)
	}
	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}

func (m *Memcache) UpdateByPkRandom(start int, end int, t chan time.Duration) {

	l := end - start
	rnd := InitSampleSet(end - start, start)

	tm_s := time.Now()
	m.valStr = "00012312301230"

	for i := 0; i<l ; i++  {
		k := strconv.Itoa(int(rnd[i]))
		err := m.mc.Replace(&memcache.Item{Key: k, Value: []byte(m.valStr)})
		checkErr(err, "Failed to update data to memcache:", k)
	}

	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}

func (m *Memcache) DeleteByPkRandom(start int, end int, t chan time.Duration) {

	l := end - start
	rnd := InitSampleSet(end - start, start)

	tm_s := time.Now()
	m.valStr = "00012312301230"

	for i := 0; i<l ; i++  {
		k := strconv.Itoa(int(rnd[i]))
		err := m.mc.Delete(k)
		checkErr(err, "Failed to delete key from memcache", k)
	}

	tm_e := time.Now()
	t <- tm_e.Sub(tm_s)
}
