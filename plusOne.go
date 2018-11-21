package main

import "fmt"

// 加一

func plusOne(digits []int) []int {

	lenDigits := len(digits)
	res := []int{1}

	if lenDigits == 1 && digits[0] == 0 {

		return res
	}

	digits[lenDigits - 1] += 1

	for i := lenDigits - 1; i > 0; i-- {

		if digits[i] == 10 {

			digits[i] = 0
			digits[i - 1] += 1
		}
	}

	if digits[0] == 10 {
		digits[0] = 0
		tmp := []int{1}
		res = append(tmp, digits...)

		return res
	}

	return digits

}

func main() {


	//nums := []int{9,9,3,4,5,6,7,1}
	nums := []int{9,9,9}

	res := plusOne(nums)

	fmt.Println(res)
}