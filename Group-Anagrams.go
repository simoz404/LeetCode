package main

// import "fmt"

// func main() {
// 	fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
// }

func groupAnagrams(strs []string)[][]string {
	var fianlAtt [][]string
	var myMap = make(map[string][]string)
	var s string
	var bigS string
	var n = 0
	for _, v := range strs {
		for i := 'a'; i <= 'z'; i++ {
			for _, char := range v {
				if char == i {
					n++
				}
			}
			if n != 0 {
				s = fmt.Sprintf("%c%d", i, n)
				bigS += s
			}
			s = ""
			n = 0
		}
		myMap[bigS] = append(myMap[bigS], v)
		bigS = ""
	}
	for _, arr := range myMap {
		fianlAtt = append(fianlAtt, arr)
	}
	return fianlAtt
}
