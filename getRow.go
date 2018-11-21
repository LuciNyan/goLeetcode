package main

import "fmt"

/**
	杨辉三角
 */
func getRow(rowIndex int) []int {

	var res []int

	if rowIndex == 0 {

		res = []int{1}

		return res
	}

	if rowIndex == 1 {

		res = []int{1, 1}

		return res
	}

	lastRound := []int{1, 1}

	for i := 0; i < rowIndex - 1; i++ {

		res = []int{}

		for j := 0; j < len(lastRound); j++ {

			if j == 0 {

				res = append(res, 1)
				continue
			}

			res = append(res, lastRound[j - 1] + lastRound[j])

		}

		res = append(res, 1)

		lastRound = res
	}

	return res

}

func main()  {

	rowIndex := 3

	res := getRow(rowIndex)

	fmt.Println(res)

}
