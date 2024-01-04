package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x int
	fmt.Scanf("%d %d", &n, &x)

	aboba := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &aboba[i])
	}

	result := findPairWithSum(aboba, x)

	if len(result) > 0 {
		sort.Ints(result)
		fmt.Print(result[0])
		for i := 1; i < len(result); i++ {
			fmt.Printf(" %d", result[i])
		}
		fmt.Println()
	} else {
		fmt.Println("0 0")
	}
}

func findPairWithSum(arr []int, target int) []int {
	seen := make(map[int]bool)
	for _, num := range arr {
		if seen[target-num] {
			return []int{num, target - num}
		}
		seen[num] = true
	}
	return []int{}
}



