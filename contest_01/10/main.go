package main
import "fmt"
func main() {
    var ot1, ot2, ot3 string
    net := "Нет"
    da := "Да"
    fmt.Scan(&ot1)
    fmt.Scan(&ot2)
    fmt.Scan(&ot3)
    if ot1 == net{
        if ot2 == net{
            if ot3 == net {
                fmt.Println("Кот")
            }else if ot3 == da {
                fmt.Println("Жираф")
            }
        }else if ot2 == da{
            if ot3 == net{
                fmt.Println("Курица")
            }else if ot3 == da{
                fmt.Println("Страус")
            }
        }
    }else if ot1 == da{
        if ot2 == net{
            if ot3 == net {
                fmt.Println("Дельфин")
            }else if ot3 == da {
                fmt.Println("Плезиозавры")
            }
        }else if ot2 == da{
            if ot3 == net{
                fmt.Println("Пингвин")
            }else if ot3 == da{
                fmt.Println("Утка")
            }
        }
    }
}
