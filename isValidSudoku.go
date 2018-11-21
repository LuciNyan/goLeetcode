package main

import "fmt"

/**
	有效的数独
 */
func checkValid9(nums []byte) bool {
	map1 := make(map[byte]int)
	map1['1'] = 1
	map1['2'] = 2
	map1['3'] = 3
	map1['4'] = 4
	map1['5'] = 5
	map1['6'] = 6
	map1['7'] = 7
	map1['8'] = 8
	map1['9'] = 9

	flag := []int{0,0,0,0,0,0,0,0,0,0}

	res := true

	for _, v := range nums {

		if v == '.' {

			continue
		}

		if flag[map1[v]] == 0 {
			flag[map1[v]] = 1
		} else {
			res = false

			break
		}

	}

	return res
}

func isValidSudoku(board [][]byte) bool {

	res := true

	for _, v := range board {

		res = checkValid9(v)
		if res == false {

			return res
		}

	}

	var tmp []byte
	for i := 0; i < 9; i++ {

		tmp = []byte{}

		for _, v := range board {

			tmp = append(tmp, v[i])

		}

		res = checkValid9(tmp)
		if res == false {

			return res
		}
	}

	for i := 0; i < 9; i += 3 {

		for j := 0; j < 9; j += 3{

			tmp = []byte{}
			tmp = append(tmp, board[i][j])
			tmp = append(tmp, board[i][j + 1])
			tmp = append(tmp, board[i][j + 2])
			tmp = append(tmp, board[i + 1][j])
			tmp = append(tmp, board[i + 2][j])
			tmp = append(tmp, board[i + 1][j + 1])
			tmp = append(tmp, board[i + 1][j + 2])
			tmp = append(tmp, board[i + 2][j + 1])
			tmp = append(tmp, board[i + 2][j + 2])
			res = checkValid9(tmp)
			if res == false {

				return res
			}

		}
	}

	return true
}

func main()  {

	board := [][]byte{
		{'8','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}
	res := isValidSudoku(board)

	fmt.Print(res)

}
