package main

import "fmt"

// 必备题库，练习一次
// @Author: hyc
// @Description:
// @Date: 2021/9/8 22:42
func main() {
	// 归并排序
	var (
		nums = []int{9, 8, 10, 7, 6, 5, 4, 3, 2, 1}
	)

	ans := mergeSort5(nums)

	fmt.Printf("ans = %+v", ans)
}

func mergeSort5(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	left := mergeSort5(nums[:mid])
	right := mergeSort5(nums[mid:])

	ans := merge5(left, right)

	return ans
}

//
func merge5(left []int, right []int) []int {

	tmp := make([]int, 0)

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}

	}

	for i < len(left) {
		tmp = append(tmp, left[i])
		i++
	}

	for j < len(right) {
		tmp = append(tmp, right[j])
		j++
	}

	return tmp
}
