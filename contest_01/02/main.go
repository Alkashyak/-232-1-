package main
import "fmt"
import "math"
func main() {
    chel := 0.5 * 365
    fmt.Println(chel, int(math.Ceil(chel/32)), int(math.Ceil(chel/20)))

} 
