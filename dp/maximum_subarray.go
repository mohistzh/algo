package main
import(
	"fmt"
	"math"
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

	maximum := -10086 // always catch maximum sum
	maxSofar := 0 // Maximum sum of subarray

	for i := 0; i < len(nums); i++ {
		maxSofar = int(math.Max(float64(nums[i]), float64(nums[i] + maxSofar)))
		if maximum < maxSofar {
			maximum = maxSofar
		}

	}
	return maximum
}

func kadaneAlgoPrintIndices(nums []int) int {
	maximum := -10086
	maxSofar := 0
	start, end := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[i] + maxSofar {
			maxSofar = nums[i]
			start = i 
		} else {
			maxSofar = nums[i] + maxSofar
		}
		if maximum < maxSofar {
			maximum = maxSofar
			end = i  
		}
	}
	fmt.Println("Maximum subarray index {", start, end, "}")
	return maximum
}

func main() {
	nums := []int{
		-2,1,-3,4,-1,2,1,-5,4,
	}
	fmt.Println(maxSubArrayBF(nums))
	fmt.Println(kadaneAlgo(nums))
	fmt.Println(kadaneAlgoPrintIndices(nums))
}