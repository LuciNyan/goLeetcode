package main

import "fmt"

// 杨辉三角

func generate(numRows int) interface{} {

	var tmpArr []int

	res := [][]int{{1}}

	if numRows == 0 {

		//res = [][]int{{}}

		return tmpArr
	}

	if numRows == 1 {

		return res
	}
	if numRows == 2 {

		res = [][]int{{1}, {1, 1}}

		return res
	}

	res = [][]int{{1}, {1, 1}}
	prevArr := []int{1, 1}

	for i:= 3; i <= numRows; i++ {

		for j:= 0; j < i; j++ {

			if j == 0 || j == i - 1{
				tmpArr = append(tmpArr, 1)

				continue
			}

			tmpArr = append(tmpArr, prevArr[j] + prevArr[j - 1])


		}
		prevArr = tmpArr
		res = append(res, tmpArr)
		tmpArr = []int{}

	}

	return res
}

func main() {

	numRows := 5

	res := generate(numRows)

	fmt.Println(res)
}