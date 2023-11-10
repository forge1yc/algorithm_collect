package main

import "fmt"

// 必备题库
// @Author: hyc
// @Description:
// @Date: 2021/8/16 3:33 下午
func main() {
	// 归并排序
	var (
		nums = []int{9,10,7,6,5,4,3,2,1}
	)

	ans := mergeSort(nums)


	fmt.Printf("%+v\n",ans)

}

func mergeSort(nums []int) []int {

	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2

	a := mergeSort(nums[:mid])
	b := mergeSort(nums[mid:])

	return merge(a,b)
}

func merge(a,b []int) []int {

	var (
		tmp = make([]int,0)
	)

	for len(a) != 0 && len(b) != 0 {
		if a[0] < b[0] {
			tmp = append(tmp,a[0])
			a = a[1:]
		} else {
			tmp = append(tmp,b[0])
			b = b[1:]
		}
	}

	tmp = append(tmp,a...)
	tmp = append(tmp,b...)

	return tmp
}