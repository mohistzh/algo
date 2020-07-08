package main

import (
	"fmt"
	"math"
)

/**
	Split to two parts to calculate maximum values, then to get the biggest one.
**/
func houseRobberSolution1(nums []int) int {
	maximum := 0
	odd, even := 0, 0
	for i := 0; i < len(nums); i++ {
		if (i & 1) == 0 { // even number
			even += nums[i]
			even = int(math.Max(float64(odd), float64(even)))
		} else {
			odd += nums[i]
			odd = int(math.Max(float64(odd), float64(even)))
		}
	}
	maximum = int(math.Max(float64(odd), float64(even)))
	return maximum
}

func main() {
	house1 := []int{
		1, 2, 3, 1,
	}
	house2 := []int{
		2, 7, 9, 3, 1,
	}
	house3 := []int{
		2, 1, 1, 2,
	}

	fmt.Println(houseRobberSolution1(house1))
	fmt.Println(houseRobberSolution1(house2))
	fmt.Println(houseRobberSolution1(house3))
}
