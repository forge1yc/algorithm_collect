package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 定义运算函数类型
type operationFunc func(int, int) int

// 定义一个集合
var setMap = map[string]bool{
	"*": true,
	"/": true,
	"+": true,
	"-": true,
}

// 定义运算函数
func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }
func divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

// 将字符串运算符转换为实际的运算函数
func getOperation(op string) (operationFunc, error) {
	switch op {
	case "+":
		return add, nil
	case "-":
		return subtract, nil
	case "*":
		return multiply, nil
	case "/":
		return divide, nil
	default:
		return nil, errors.New("unsupported operator")
	}
}

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		re, err := strconv.Atoi(tokens[0])
		if err != nil {

		}
		return re
	}

	tmpS := s2(tokens)

	return evalRPN(tmpS)
}

func s2(tokens []string) []string {
	if len(tokens) <= 1 {
		return tokens
	}

	for i, v := range tokens {
		if i >= 2 {
			if setMap[v] {
				tmp := solve(tokens[i-2 : i+1])
				start := i - 2
				end := i + 1
				//fmt.Println("tmp", tmp, i, start, end)
				//tokens = append(tokens[:i-1], strconv.FormatInt(int64(tmp), 10))
				tokens = append(tokens[:start], tokens[end:]...)
				//fmt.Println("tokens", tokens)

				tokens = append(tokens[:start], append([]string{strconv.Itoa(tmp)}, tokens[start:]...)...)
				fmt.Println("tokens", tokens)
				return tokens
			}
		}

	}
	return []string{}
}

func solve(tokens []string) int {

	if len(tokens) <= 1 {
		result, err := strconv.Atoi(tokens[0])
		if err != nil {
			return result
		} else {
			return 0
		}
	}

	left, err := strconv.Atoi(tokens[0])
	if err != nil {

	}
	right, err := strconv.Atoi(tokens[1])
	if err != nil {

	}

	funcOperation, err := getOperation(tokens[2])
	if err != nil {

	}
	return funcOperation(left, right)

}

func main() {
	//tmp := 3
	//start := 2

	//tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	tokens := []string{"4", "13", "5", "/", "+"}
	//ss := append([]string{strconv.Itoa(tmp)}, tokens[start:]...)
	//fmt.Println("ss", ss)
	//tokens := []string{"2", "1", "+", "3", "*"}
	re := evalRPN(tokens)
	fmt.Print(re)
}
