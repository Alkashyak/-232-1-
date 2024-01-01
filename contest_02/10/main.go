package main
import "fmt"

func shift(input []int, step int) {
    shiftres := make([]int, 10)
    for i := 0; i < 10; i++ {
        newi := (i + step + 1000000000 ) % 10
        shiftres[newi] = input[i]
    } 
    for i := 0; i < 10; i++ {
        input[i] = shiftres[i]
    } 
}

func main(){
    var steps int
    fmt.Scan(&steps)

    var data [10]int
    for index := range data{
        fmt.Scan(&data[index])
    }

    shift(data[:], steps);
    for _, value := range data{
        fmt.Printf("%d ", value)
    }
}
