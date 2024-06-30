package main

import "fmt"

func twoSum(arr []int, target int) (int, int) {
	seen := make(map[int]int)
	for i, num := range arr {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return i, j
		}
		seen[num] = i
	}
	return -1, -1
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	idx1, idx2 := twoSum(arr, target)
	if idx1 != -1 && idx2 != -1 {
		fmt.Printf("Indices of the two numbers that sum up to %d: %d, %d\n", target, idx1, idx2)
	} else {
		fmt.Printf("No such pair found that sums up to %d\n", target)
	}
}
