package main
import (
    "bufio"
    "fmt"
    "os"
)

func isLucky (n string) bool {
    if (n[0] + n[1] + n[2]) == (n[3] + n[4] + n[5]) {
        return true
    }else {
        return false
    }
}

func main() {
    number, _  := bufio.NewReader(os.Stdin).ReadString('\n')
    if isLucky( number[:len(number)-1] ){  // Обрезаем \n в конце строки
        fmt.Println("YES")
    }else{
        fmt.Println("NO")
    }
}
