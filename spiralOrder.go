package main

import "fmt"

// 螺旋矩阵

func spiralOrder(matrix [][]int) []int {

	var res []int

	direction := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	directionIndex := 0

	x := 0
	y := 0

	n := len(matrix)

	if n == 0 {

		return res
	}

	m := len(matrix[0])

	if m == 0 {

		return res
	}


	all := m * n
	minX := 0
	maxX := m - 1
	minY := 0
	maxY := n - 1


	for i := 0; i < all; i++ {

		fmt.Printf("%d %d ", x, y)
		fmt.Printf("x:(%d~%d) ", minX, maxX)
		fmt.Printf("y:(%d~%d) \n", minY, maxY)

		res = append(res, matrix[y][x])

		switch directionIndex % 4 {
		case 0:
			if x >= maxX {
				x = maxX
				directionIndex++
				minY++
			}

		case 1:
			if y >= maxY {
				y = maxY
				directionIndex++
				maxX--
			}

		case 2:
			if x <= minX {
				x = minX
				directionIndex++
				maxY--
			}

		case 3:
			if y <= minY {
				y = minY
				directionIndex++
				minX++
			}


		default:


		}

		x += direction[directionIndex % 4][0]
		y += direction[directionIndex % 4][1]

	}

	return res

}

func main() {


	//matrix := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	matrix := [][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}, {13,14,15,16}}
	//matrix := [][]int{{1}}

	res := spiralOrder(matrix)

	fmt.Println(res)
}
