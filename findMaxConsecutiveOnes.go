package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {

	index := -1
	max := 0
	tmp := 0

	for i := 0; i <= len(nums); i++ {

		if i == len(nums) || nums[i] == 0 {
			tmp = i - index - 1
			index = i

			if tmp > max {
				max = tmp
			}

		}

	}

	return max

}

func main()  {

	nums := []int{1, 1, 0, 0, 0, 1, 1, 1}
	res := findMaxConsecutiveOnes(nums)

	fmt.Println(res)
}
