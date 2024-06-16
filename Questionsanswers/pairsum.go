package main

import "fmt"

func PairSum(arr []int, sum int) [][2]int {
	var pairs [][2]int
	seen := make(map[int]bool)
	for _, num := range arr {
		complement := sum - num
		if seen[complement] {
			pairs = append(pairs, [2]int{complement, num})

		}
		seen[num] = true
	}
	return pairs
}

func main() {
	arr := []int{20, 15, 17, 18, 19}
	sum := 35
	pairs := PairSum(arr, sum)

	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Println("Pairs with sum", sum, ":")
		for _, pair := range pairs {
			fmt.Println(pair[0], pair[1])
		}
	}
}
