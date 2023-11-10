package main

import "fmt"

// 必备题库
// @Author: hyc
// @Description:
// @Date: 2021/8/16 3:33 下午
func main() {

	var (
		nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	)

	quickSortCopy2(nums, 0, len(nums)-1)

	fmt.Printf("%+v\n", nums)

}

func quickSortCopy2(nums []int, left, right int) {
	if left < right {
		p := partitionCopy2(nums, left, right)
		quickSortCopy2(nums, left, p-1)
		quickSortCopy2(nums, p+1, right)
	}
}

func partitionCopy2(nums []int, left, right int) int {

	piv := nums[left]

	for left < right {

		for nums[right] >= piv && left < right {
			right--
		}
		nums[left] = nums[right]
		for nums[left] < piv && left < right {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = right

	return left
}
