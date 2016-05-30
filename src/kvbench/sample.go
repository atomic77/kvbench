package main

import (
	"github.com/montanaflynn/stats"
)

func InitSampleSet(n int, start_key int) (stats.Float64Data) {
	// Create a randomized slice of ints in the given range

	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = start_key + i;
	}

	f := stats.LoadRawData(s)
	r, err := stats.Sample(f, n, false)
	checkErr(err, "Error when creating randomized array")
	return r
}


