package main
import "strings"
import "sort"
import "bufio"
import "fmt"
import "os"
func main(){
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
strok1 := scanner.Text()
her := strings.Split(strok1," ")
var lasther []string
for idx1,element1 := range her {
    if element1 == "end"{
         break
        }
        var count int = 0
        for idx2,element2 := range her {
         if element2 == "end"{
          break
         }
            if element2 == element1 && idx2 >= idx1 {
                count += 1
                if count >= 2{
                    var exist bool = false
                    for _,finalelement := range lasther{
                        if finalelement == element1{
                            exist = true
                        }
                    }
                    if exist == false{
                        lasther = append(lasther,element1)
                    }
                    break
                }
            }
        }
    }
    sort.Strings(lasther)
for _,element := range lasther {
    fmt.Print(element, " ")
}
}
