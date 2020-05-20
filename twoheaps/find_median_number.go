package main

/*
Design a class to calculate the median of a number stream. The class should have the following two methods:

insertNum(int num): stores the number in the class
findMedian(): returns the median of all numbers inserted in the class

If the count of numbers inserted in the class is even, the median will be the average of the middle two numbers.
*/
import (
	"fmt"
)

// MedianOfAStream object
type MedianOfAStream struct {
}

func (obj *MedianOfAStream) insertNumber(number int) {

}

func (obj *MedianOfAStream) findMedian() int {
	return -1
}

func main() {
	medianObj := MedianOfAStream{}
	medianObj.insertNumber(3)
	medianObj.insertNumber(1)
	fmt.Println("The median is:", medianObj.findMedian())
}
