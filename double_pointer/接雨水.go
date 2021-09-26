package main

import (
	"fmt"
)

// 接雨水(一次过啦，看了一点提示）,宝刀未老！！！
// @Author: houyichao
// @Description: 这个是为了练习，保持go的手感，毕竟go的手感太好了，但是这道题也是面试的高频题
// @Date: 23:26
func main() {

	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	ans := trap(nums)

	fmt.Printf("%+v\n", ans)

}

func trap(height []int) int {

	if len(height) <= 2 {
		return 0
	}

	var (
		right = len(height) - 2
		areas = make([]int, 0)
	)

	// 每一个点都要做一次便利
	for i := 1; i <= right; i++ {

		// 从i开始做双指针，但是我觉得的这个复杂度会超出限制
		lHeight := maxHeight(height[0:i], height[i])
		//fmt.Printf("%+v\n",lHeight)
		rHeight := maxHeight(height[i+1:], height[i])
		//fmt.Printf("%+v\n",rHeight)

		//fmt.Printf("%+v\n","#######################")

		realHeight := min(lHeight, rHeight) - height[i]

		if realHeight <= 0 {
			continue
		} else {
			areas = append(areas, realHeight*1)
		}

	}
	fmt.Printf("%+v\n", areas)

	sum := 0
	for _, v := range areas {
		sum += v
	}

	return sum

}

func maxHeight(nums []int, tmp int) int {

	m := -1

	for _, v := range nums {
		if v > tmp {
			m = max(v, m)
		}
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
