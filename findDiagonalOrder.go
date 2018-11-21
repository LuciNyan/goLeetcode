package main

// 对角线遍历

import (
	"fmt"
)

type haru struct {
	m int
	n int

	sgnM int
	sgnN int

	lenM int
	lenN int

	matrix [][]int
	arr []int
}

func (h haru) getSgn() (m int, n int) {

	return h.sgnM, h.sgnN
}

func (h haru) getIndex() (int, int) {

	return h.m, h.n
}

func (h *haru) changeSgn() {
	h.sgnM *= -1
	h.sgnN *= -1
}

func (h *haru) setLen() {
	h.lenN = len(h.matrix)
	h.lenM = len(h.matrix[0])
}

func (h haru) getLen() (int, int) {

	return h.lenM, h.lenN
}

func (h haru) getValue() int {

	fmt.Println(h.n, h.m)
	return h.matrix[h.n][h.m]
}

func (h *haru) next() bool {
	res := true

	h.m = h.m + h.sgnM
	h.n = h.n + h.sgnN

	flag := 0

	//fmt.Println(h.m, h.n)
	//fmt.Println(h.lenM, h.lenN)


	if h.m >= h.lenM {

		h.m = h.lenM - 1
		h.n += 2

		if h.n >= h.lenN {

			return false

		}

		flag = 1

	}

	if h.n >= h.lenN {

		h.n = h.lenN - 1
		h.m += 2

		if h.m >= h.lenM {

			return false

		}

		flag = 1

	}

	if h.m < 0 {
		h.m = 0
		flag = 1
	}

	if h.n < 0 {
		h.n = 0
		flag = 1
	}

	if flag == 1 {
		h.changeSgn()
	}

	if h.m == h.lenM - 1 && h.n == h.lenN - 1 {
		res = false
	}

	h.arr = append(h.arr, h.getValue())

	return res
}

func (h haru) handle() []int {
	var res []int
	check := true

	if len(h.matrix) == 0 {

		return res
	}

	h.setLen()

	if h.lenM == 0 || h.lenN == 0 {

		return res
	}

	//fmt.Println(h.getIndex())
	//fmt.Println(h.getSgn())

	h.arr = append(h.arr, h.getValue())
	//check = h.next()


	for check == true {

		//fmt.Println(h.getIndex())
		//fmt.Println(h.getSgn())

		//res = append(res, h.getValue())
		check = h.next()
	}

	//res = append(res, h.getValue())


	//h.next()
	//fmt.Println(h.getIndex())
	//fmt.Println(h.getSgn())
	//res = append(res, h.getValue())

	return h.arr
}


func findDiagonalOrder(matrix [][]int) []int {


	haru := haru{
		m: 0,
		n: 0,
		sgnM: 1,
		sgnN: -1,
		matrix: matrix,
	}

	return haru.handle()
}


func main() {


	//matrix := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	matrix := [][]int{{1}}

	res := findDiagonalOrder(matrix)

	fmt.Println(res)
}
