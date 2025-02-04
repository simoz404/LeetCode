package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	var myMap = make(map[int]int)
	var result = []int{}
	for _, v := range nums {
		myMap[v]++
	}
	var top int
	for i := 1; i <= k; i++ {
		top = findFrequent(myMap)
		result = append(result, top)
		delete(myMap, top)
	}
	return result
}

func findFrequent(myMap map[int]int) int {
	var max = 0
	var top = 0
	for i, v := range myMap {
		if v > max {
			max = v
			top = i
		}
	}
	return top
}