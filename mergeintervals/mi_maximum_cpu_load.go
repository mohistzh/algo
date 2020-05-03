package main

import (
	"fmt"
)

// CPULoad data structure
type CPULoad struct {
	start int
	end   int
	load  int
}

/*
We are given a list of Jobs. Each job has a Start time, an End time, and a CPU load when it is running.
Our goal is to find the maximum CPU load at any time if all the jobs are running on the same machine.
*/
func findMaximumCPULoad(jobs []CPULoad) int {
	fmt.Println(jobs)
	return -1
}

func main() {
	jobs := []CPULoad{CPULoad{start: 1, end: 4, load: 3}, CPULoad{start: 2, end: 5, load: 4}, CPULoad{start: 7, end: 9, load: 6}}
	fmt.Println(findMaximumCPULoad(jobs))
}
