package main

import "fmt"

// 必备题库，练习一次
// @Author: hyc
// @Description:
// @Date: 2021/9/8 22:42
func main() {
	// 归并排序
	var (
		nums = []int{9, 10, 7, 6, 5, 4, 3, 2, 1}
	)

	ans := mergeSort4(nums)

	fmt.Printf("ans = %+v", ans)
}

func mergeSort4(nums []int) []int {

	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2

	left := mergeSort4(nums[:mid])
	right := mergeSort4(nums[mid:])

	fmt.Printf("left = %v \n", left)
	fmt.Printf("right = %v \n", right)

	ans := merge4(left, right)

	return ans
}

func merge4(left []int, right []int) []int {

	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			tmp = append(tmp, left[i])
			i++
			fmt.Printf("i = %d, len(left) = %d\n", i, len(left))

		} else {
			tmp = append(tmp, right[j])
			j++
			fmt.Printf("j = %d, len(right) = %d\n", j, len(right))
		}
	}

	// 上面没有进行截取，所以双指针的方法不能这样添加，否则会添加全部
	//tmp = append(tmp, left...)
	//tmp = append(tmp, right...)

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
