package main

import "fmt"

// 寻找数组中心索引

func pivotIndex(nums []int) int {

	numsLen := len(nums)
	leftSum := 0
	rightSum := 0

	if numsLen == 0 {

		return -1
	}

	if numsLen == 1 {

		return 0
	}

	if numsLen == 2 {

		return -1
	}

	midIndex := 0


	for i := 0; i < midIndex; i++ {
		leftSum += nums[i]
	}

	for i := numsLen - 1; i > midIndex; i-- {
		fmt.Println(i)
		rightSum += nums[i]
	}

	for true {

		if leftSum == rightSum {

			return midIndex
		}

		midIndex++

		if midIndex == numsLen {

			return -1
		}

		leftSum += nums[midIndex - 1]
		rightSum -= nums[midIndex]

	}

	return -1

}

func main() {


	//nums := []int{1,2,3,4,5,6,7,21}
	nums := []int{-1,-1,-1,-1,-1,0}
	//nums := []int{1}

	res := pivotIndex(nums)

	fmt.Println(res)
}