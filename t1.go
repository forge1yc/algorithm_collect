package main

import "fmt"

func main() {

	var (
		nums = []int{9,8,7,6,5,4,3,2,1}
	)

	quickSort(nums,0,len(nums)-1)

	fmt.Printf("%+v\n",nums)
}

func quickSort(nums []int,left,right int) {

	if left < right {
		p := partition(nums,left,right)
		quickSort(nums,left,p-1)
		quickSort(nums,p+1,right)
	}
}

func partition(nums []int,left,right int) int {

	piv := nums[left]

	for left < right {
		for nums[right] > piv && left < right {
			right--
		}
		nums[left] = nums[right]

		for nums[left] < piv && left < right {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = piv

	return left
}