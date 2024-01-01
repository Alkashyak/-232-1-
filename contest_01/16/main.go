package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	huyna := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&huyna[i])
	}
	sredhui := make([]float64, n)
    sredhui[0] = float64(huyna[0])
    sredhui[n-1] = float64(huyna[n-1])
	for i := 1; i < n-1; i++ {
		sredhui[i] = (float64(huyna[i-1]) + float64(huyna[i]) + float64(huyna[i+1])) / 3.0
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%.10f", sredhui[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
