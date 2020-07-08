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

/**
	DP tabulation solution
**/
func houseRobberSolution2(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < n; i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+nums[i])))
	}
	return dp[n-1]

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
	fmt.Println("------------DP---------------")
	fmt.Println(houseRobberSolution2(house1))
	fmt.Println(houseRobberSolution2(house2))
	fmt.Println(houseRobberSolution2(house3))
}
