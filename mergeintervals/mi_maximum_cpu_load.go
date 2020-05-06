package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

// CPULoad data structure
type CPULoad struct {
	start int
	end   int
	load  int
}

// MinHeap1D to find out the smallest end time
// Minheap restrictions are !mh.Less(j, i) for 0 <= i < mh.Len() and 2*i+1 <= j <= 2*i+2 and j < mh.Len()
type MinHeap1D []CPULoad

func (mh MinHeap1D) Len() int {
	return len(mh)
}

func (mh MinHeap1D) Less(i, j int) bool {
	return mh[i].end < mh[j].end
}

func (mh MinHeap1D) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

// Push add an item
func (mh *MinHeap1D) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*mh = append(*mh, x.(CPULoad))
}

// Pop popup an item
func (mh *MinHeap1D) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

// Peak get the peak element
func (mh *MinHeap1D) Peak() interface{} {
	old := *mh
	n := old.Len()
	if n < 1 {
		return CPULoad{start: -1, end: -1, load: -1}
	}
	return old[n-1]
}

/*
We are given a list of Jobs. Each job has a Start time, an End time, and a CPU load when it is running.
Our goal is to find the maximum CPU load at any time if all the jobs are running on the same machine.
*/
func findMaximumCPULoad(jobs []CPULoad) int {
	sortJobs(jobs)
	currentCPULoad, maxCPULoad := 0, 0
	minHeap := &MinHeap1D{}

	for i := 0; i < len(jobs); i++ {
		job := jobs[i]
		minPeak := minHeap.Peak().(CPULoad)
		for minHeap.Len() > 0 && job.start >= minPeak.end {
			currentCPULoad -= minPeak.load
			heap.Pop(minHeap)
		}
		heap.Push(minHeap, job)
		heap.Fix(minHeap, minHeap.Len()-1)
		currentCPULoad += job.load
		maxCPULoad = int(math.Max(float64(maxCPULoad), float64(currentCPULoad)))
		fmt.Println(minHeap)
	}
	return maxCPULoad
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
