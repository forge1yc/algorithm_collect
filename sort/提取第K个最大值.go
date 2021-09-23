package main

import (
	"fmt"
	"sort"
)

// 提取第k个最大值
// @Author: hyc
// @Description:
// @Date: 2021/9/16 22:08
func main() {

	var (
		nums = []int{3,4,2,7,1,19,5,1,2,2,3,3,1}
		k = 5
	)

	ans := findKthLargest(nums,k)

	fmt.Printf("%+v\n",ans)

}


func findKthLargest( nums []int ,  k int ) int {
	// write code here
	//m := make(map[int]struct{},0)
	//
	//for _,v := range nums {
	//	m[v] = struct{}{}
	//}
	//
	//tmp := make([]int,0)
	//for k,_ := range m {
	//	tmp = append(tmp,k)
	//}

	sort.Ints(nums)


	return nums[len(nums) - k]

}