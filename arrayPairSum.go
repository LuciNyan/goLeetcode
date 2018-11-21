package main

import (
	"fmt"
	"sort"
)



func arrayPairSum(nums []int) int {

	var res int

	//fmt.Print(nums)
	sort.Ints(nums)
	//fmt.Print(nums)

	for i:= 0; i <= len(nums) - 1; i+=2 {

		res += nums[i]
	}

	return res
}


func main() {

	nums := []int{1, 4, 3, 2}

	res := arrayPairSum(nums)

	fmt.Print(res)
}
