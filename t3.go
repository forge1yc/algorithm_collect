package main

import "fmt"

// 无敌
func main() {
	// 归并排序
	var (
		nums = [][]int{{7,8,9},{1,2,3},{4,5,6}}
	)

	ans := mergeSort2(nums)

	fmt.Printf("%+v\n",ans)
}

func mergeSort2(nums [][]int) []int {

	if len(nums) == 1 {
		return nums[0]
	}

	mid := len(nums) / 2

	a := mergeSort2(nums[:mid])
	b := mergeSort2(nums[mid:])

	return merge2(a,b)
}

func merge2(a,b []int) []int {

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