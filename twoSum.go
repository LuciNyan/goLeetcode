package main

import "fmt"

func twoSum(numbers []int, target int) []int {


	l := 0
	r := len(numbers) - 1

	for {

		if numbers[l] + numbers[r] == target {

			break

		} else if numbers[l] + numbers[r] > target {

			r = r - 1

		} else {

			l = l + 1

		}

	}

	res := []int{l + 1, r + 1}

	return res
}

func main() {

	numbers := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(numbers, target)

	fmt.Print(res)
}
