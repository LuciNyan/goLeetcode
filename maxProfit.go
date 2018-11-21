package main

import "fmt"

/**
	买卖股票的最佳时机 II
 */
func maxProfit(prices []int) int {

	lenPrices := len(prices)
	res := 0

	if lenPrices == 0 || lenPrices == 1  {

		return res
	}

	if lenPrices == 2 {

		res = prices[1] - prices[0]
		if res > 0 {

			return res

		} else {

			return 0
		}

	}

	tmp := []int{prices[0], prices[1]}

	for i := 2; i < lenPrices; i++ {

		if tmp[0] > tmp[1] {

			tmp[0] = tmp[1]
			tmp[1] = prices[i]

			continue
		}
		// ----以下都是 如果第一个数 比第二个数小

		if prices[i] > tmp[1] {

			tmp[1] = prices[i]

			continue
		}

		if prices[i] < tmp[1] {

			res += tmp[1] - tmp[0]
			tmp[1] = prices[i]
			tmp[0] = prices[i]
		}

	}
	haru := tmp[1] - tmp[0]
	if haru > 0 {
		res += haru
	}

	return res
}

func main()  {

	prices := []int{7,6,4,3,1}
	res := maxProfit(prices)

	fmt.Println(res)

}
