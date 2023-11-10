package main

import "fmt"

// 必备题库
// @Author: hyc
// @Description:
// @Date: 2021/8/16 3:33 下午
func main() {

	var (
		nums = []int{8, 9, 7, 5, 6, 4, 3, 2, 1}
	)

	quickSort3(nums, 0, len(nums)-1)

	fmt.Printf("%+v\n", nums)
}

func quickSort3(nums []int, left, right int) {

	if left < right {
		p := partition3(nums, left, right)
		quickSort3(nums, 0, p-1)
		quickSort3(nums, p+1, right)
	}

}

func partition3(nums []int, left, right int) int {

	mid := nums[right]

	for left < right {

		for nums[left] < mid && left < right {
			left++
		}
		nums[right] = nums[left]

		for nums[right] >= mid && left < right {
			right--
		}
		nums[left] = nums[right]

	}

	nums[right] = mid

	return right
}
