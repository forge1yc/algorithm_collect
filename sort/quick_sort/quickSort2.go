package main

import "fmt"

// 必备题库
// @Author: hyc
// @Description:
// @Date: 2021/8/16 3:33 下午
func main() {

	var (
		nums = []int{5, 6, 4, 10, 8, 9, 7, 5, 6, 4, 3, 2, 1}
	)

	quickSort2(nums, 0, len(nums)-1)

	fmt.Printf("%+v\n", nums)
}

func quickSort2(nums []int, left, right int) {

	if left <= right {
		p := partition2(nums, left, right)

		quickSort2(nums, 0, p-1)
		quickSort2(nums, p+1, right)
	}

}

func partition2(nums []int, left, right int) int {

	midValue := nums[left]

	for left < right {

		for nums[right] >= midValue && left < right {
			right--
		}
		nums[left] = nums[right]

		for nums[left] < midValue && left < right {
			left++
		}
		nums[right] = nums[left]

	}

	nums[left] = midValue

	return left
}
