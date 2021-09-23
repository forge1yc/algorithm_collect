package main

import "fmt"

// 提取最长字符串
// @Author: hyc
// @Description: 来自华为二面，算是比较简单的题
// @Date: 2021/9/13 16:39
func main() {

	var (
		inputStr = "aabbbbbbbccdd"
	)

	ans := Solver(inputStr)

	fmt.Printf("%+v\n",ans)
}

func Solver(input string) string {

	if len(input) <= 1 {
		return input
	}

	left := 0
	right := 0
	minLength := 1

	tmpLeft := 0
	tmpRight := 1

	for tmpRight < len(input) {
		if input[tmpRight] == input[tmpRight-1] {
			tmpRight++
			if tmpRight - tmpLeft + 1 > minLength {
				left = tmpLeft
				right = tmpRight
				// 这个面试的时候忘记写了，但是面试官没有注意到，因为答案对了
				minLength = tmpRight - tmpLeft + 1
			}
		} else {
			tmpLeft = tmpRight
			tmpRight++
		}
	}

	return input[left:right]

}
