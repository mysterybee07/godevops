package main

import "fmt"

func NoRepeat(arr []int) int {
	seen := make(map[int]int)
	for _, num := range arr {
		seen[num]++
	}
	for n, count := range seen {
		if count == 1 {
			return n
		}
	}
	return 0
}
func main() {
	arr := []int{1, 1, 3, 2, 3}
	n := NoRepeat(arr)
	fmt.Println(n)

}
