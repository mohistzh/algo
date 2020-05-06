package main

import (
	"fmt"
	"sort"
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
	sortJobs(jobs)
	fmt.Println(jobs)
	return -1
}
func sortJobs(jobs []CPULoad) {
	sort.Slice(jobs[:], func(i, j int) bool {
		return jobs[i].start < jobs[j].start
	})
}

func main() {
	jobs1 := []CPULoad{CPULoad{start: 1, end: 4, load: 3}, CPULoad{start: 2, end: 5, load: 4}, CPULoad{start: 7, end: 9, load: 6}}
	jobs2 := []CPULoad{CPULoad{start: 6, end: 7, load: 10}, CPULoad{start: 2, end: 4, load: 11}, CPULoad{start: 8, end: 12, load: 15}}
	jobs3 := []CPULoad{CPULoad{start: 1, end: 4, load: 2}, CPULoad{start: 2, end: 4, load: 1}, CPULoad{start: 3, end: 6, load: 5}}
	fmt.Println(findMaximumCPULoad(jobs1))
	fmt.Println(findMaximumCPULoad(jobs2))
	fmt.Println(findMaximumCPULoad(jobs3))
}
