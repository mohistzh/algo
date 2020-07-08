package main
import(
	"fmt"
)
/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
**/
/**
	O(n)^2
**/
func maxSubArrayBF(nums []int) int {
	n := len(nums)
	maximun := -10086
	for left := 0; left < n; left++ {
		rightWindowSum := 0
		for right := left; right < n; right++ {
			rightWindowSum += nums[right]
			if rightWindowSum > maximun {
				maximun = rightWindowSum
			}
		}
	}
	return maximun
}

/**
	O(n)
**/
func kadaneAlgo(nums []int) int {

	maximumSofar := 0
	maximumEndingHere := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[i] + maximumEndingHere {
			maximumEndingHere = nums[i]
		} else {
			maximumEndingHere = nums[i] + maximumEndingHere

			if maximumEndingHere > maximumSofar {
				maximumSofar = maximumEndingHere
			}
		}

	}
	return maximumSofar
}

func main() {
	nums := []int{
		-2,1,-3,4,-1,2,1,-5,4,
	}
	fmt.Println(maxSubArrayBF(nums))
	fmt.Println(kadaneAlgo(nums))
}