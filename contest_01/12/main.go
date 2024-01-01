package main
import "fmt"
func main() {
    var x, cn int 
    fmt.Scan(&x)
    for x != 1 {
        if x % 2 == 0 {
            x = x/2
            cn++
        }else {
            x = 3 * x + 1
            cn++}
    }
    fmt.Println(cn)
}
