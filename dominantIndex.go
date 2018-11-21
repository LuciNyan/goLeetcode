package main

import "fmt"

// 至少是其他数字两倍的最大数

func dominantIndex(nums []int) int {

	lenNums := len(nums)
	a := 0

	if lenNums == 0 {

		return -1
	}

	if lenNums == 1 {

		return 0
	}

	for i := 0; i < lenNums; i++ {
		a = nums[i]

		for j := 0; j < lenNums; j++ {
			//if i == j {
			//	continue
			//}

			if a < 2 * nums[j] && i != j{

				break
			}

			if j == lenNums - 1 {

				return i
			}
		}
	}

	return -1

}

func main() {


	//nums := []int{1,2,3,4,5,6,7,21}
	//nums := []int{3,6,1,0}
	//nums := []int{1,2,3,4}
	nums := []int{0,0,0,1} // 3
	//nums := []int{1}

	res := dominantIndex(nums)

	fmt.Println(res)
}
