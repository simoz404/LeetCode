package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	var product = make([]int, len(nums))
	var pre = 1
	for i, v := range nums {
		product[i] = pre
		pre *= v
	}
	var suf = 1
	for i := len(nums)-1; i >= 0; i-- {
		product[i]*=suf
		suf *= nums[i]
	}
	return product
}



// Sample Solution!!

// func productExceptSelf(nums []int) []int {
// 	var product = make([]int, len(nums))
// 	var res = 1
// 	for i, _ := range nums {
// 		for i1, v := range nums {
// 			if i != i1 {
// 				res *= v
// 			}
// 		}
// 		product[i] = res
// 		res = 1
// 	}
// 	return product
// }
