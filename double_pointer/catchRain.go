package main

import (
	"algorithm_collect/utils"
	"fmt"
)

// 称最多水的容器经典题目
// @Author: hyc
// @Description: 字节一面面试题，跪了，丢人，这么简单
// @Date: 2021/8/19 9:29 下午
func main() {

	var (
		nums = []int{3,2,5,4,6,2}
	)

	ans := SolverCatchRain(nums)

	fmt.Printf("%+v\n",ans)
}

func SolverCatchRain(nums []int) int {

	if len(nums) < 2 {
		return 0
	}

	right := len(nums) - 1
	left := 0

	max := (right - left) * utils.Min(nums[left],nums[right])

	for left < right {

		lh := nums[left]
		rh := nums[right]

		if lh >= rh {
			right--
			if nums[right] > rh {
				max = utils.Max(Area(nums,left,right),max)
			}
		} else {
			left++
			if nums[left] > lh {
				max = utils.Max(Area(nums,left,right),max)
			}
		}
	}

	return max
}

func Area(nums []int,a,b int) int {
	return (b-a) * utils.Min(nums[a],nums[b])
}


