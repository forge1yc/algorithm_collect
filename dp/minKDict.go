package main

import "fmt"

// abcdeafg 取k个字符，求最小字典序
// @Author: hyc
// @Description: 只会暴力解，一点思路没有，更丢人  | 暂时没有思路，搁置
// @Date: 2021/8/19 9:24 下午
func main() {

	var (
		input = "abcdeafg"
	)

	ans := SolverK(input)

	fmt.Printf("%+v\n",ans)
}

func SolverK(str string) string {
	return str
}
