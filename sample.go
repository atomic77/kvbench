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


