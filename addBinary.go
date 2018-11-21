package main

import (
	"fmt"
	"strings"
	"strconv"
)

// 二进制求和

func addBinary(a string, b string) string {

	listA := strings.Split(a, "")
	listB := strings.Split(b, "")

	//fmt.Print(listA)
	//fmt.Print(listB)

	//return "22"

	lenA := len(listA)
	lenB := len(listB)
	vA := 0
	vB := 0
	maxIndexA := lenA - 1
	maxIndexB := lenB - 1

	lenC := lenA
	if lenB > lenA {
		lenC = lenB
	}
	maxIndexC := lenC - 1

	var c = make([]int, lenC)
	var d []int

	if lenB > lenA {
		lenC = lenB
	}


	for i := 0; i < lenC; i++ {

		if maxIndexA - i >= 0 {
			vA, _ = strconv.Atoi(listA[maxIndexA - i])
		} else {
			vA = 0
		}

		if maxIndexB - i >= 0 {
			vB, _ = strconv.Atoi(listB[maxIndexB - i])
		} else {
			vB = 0
		}
		fmt.Printf("vA + vB = %d + %d | \n", vA, vB)

		c[maxIndexC - i] = vA + vB
	}

	fmt.Println(c)


	for i := lenC - 1; i > 0; i-- {
		c[i - 1] += c[i] / 2
		c[i] = c[i] % 2
	}

	tmp1 := c[0] / 2
	c[0] = c[0] % 2

	if tmp1 != 0 {
		d = append(d, tmp1)
	}
	d = append(d, c...)

	var e string
	for _, v := range d {

		e = e + strconv.Itoa(v)
	}

	return e
}

func main() {

	a :=  "101"
	b := "1011"


	res := addBinary(a, b)

	fmt.Println(res)
}