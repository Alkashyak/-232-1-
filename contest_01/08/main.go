package main

import "fmt"

func main() {
    var a, b, c, maximussuka int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    if a > b {
        maximussuka = a
    }else {
        maximussuka = b
    }
    if c > maximussuka {
        fmt.Println(c)
    }else{
        fmt.Println(maximussuka)
    }
}
