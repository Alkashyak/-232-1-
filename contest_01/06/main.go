package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var n5000, n1000, n500, n200, n100 int = 0, 0, 0, 0, 0

	for N > 0 {
		if N >= 5000 {
			n5000++
			N -= 5000
		} else if N >= 1000 {
			n1000++
			N -= 1000
		} else if N >= 500 {
			n500++
			N -= 500
		} else if N >= 200 {
			n200++
			N -= 200
		} else if N >= 100 {
			n100++
			N -= 100
		} else {
			break
		}
	}

	fmt.Println(n5000, n1000, n500, n200, n100)
}
