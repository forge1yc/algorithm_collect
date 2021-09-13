package main

import "fmt"

// 斐波那契: N个台阶，每次只能跳1个或者3个，问有几种跳法，有一个变种，AB拿数字，剩下最后一个的是失败的人，我没有做出来
// @Author: hyc
// @Description: 来自华为一面的第二道算法题
// @Date: 2021/9/13 16:39
func main() {

	var (
		n = 20
		nums = make([]int,n+1)
	)

	nums[1] = 1
	nums[2] = 1
	nums[3] = 2

	for i := 4; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-3]
	}

	fmt.Printf("%+v\n",nums[n])


}
