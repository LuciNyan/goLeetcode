package main

import "fmt"

/**
	旋转图像
 */
func rotate(matrix [][]int)  {

	var x int
	var y int

	n := len(matrix)

	//i := 0
	//j := 0
	//x = j
	//y = n - 1 - i
	//
	//for k:= 0; k < 3; k++ {
	//	fmt.Println(x)
	//	fmt.Println(y)
	//	matrix[i][j], matrix[x][y] = matrix[x][y], matrix[i][j]
	//	x, y = y, n - 1 - x
	//}

	for i := 0; i * 2 <= n; i++ {
		for j := i; j < n - i - 1; j++ {

			fmt.Println(matrix[i][j])
			x = j
			y = n - 1 - i

			for k:= 0; k < 3; k++ {
				matrix[i][j], matrix[x][y] = matrix[x][y], matrix[i][j]
				x, y = y, n - 1 - x

			}
		}
	}

	for i:= 0; i < n; i++ {

		fmt.Println(matrix[i])
	}

}

func main()  {

	matrix := [][]int{
		{  5,  1,  9, 11, 57},
		{  2,  4,  8, 10, 58},
		{ 13,  3,  6,  7, 99},
		{ 15, 14, 12, 16, 98},
		{ 61, 54, 42, 31, 55},
	}

	rotate(matrix)

}
