package main
import "fmt"

type Human struct{
    name string
    age int
}

func printTypes(data []interface{}) {
    var str string;
    for _, val := range data {
        switch val.(type) {
            case int:
            str = "int"
            case string:
            str = "string"
            case Human:
            str = "Human"
            case *Human:
            str = "*Human"
            default:
            str = "interface"
        }
        fmt.Print(str, " ")
    }
    fmt.Println()
}



func main() {
    var count int
    fmt.Scan(&count)
    
    var data []interface{} = make([]interface{}, count)
    for i:=0; i<count; i++{
        var t string
        fmt.Scan(&t)
        switch t {
            case "int": 
                var val int = 0
                data[i] = interface{}(val)
            case "string":
                var val string = "0"
                data[i] = interface{}(val)
            case "Human":
                var val Human = Human{}
                data[i] = interface{}(val)
            case "*Human":
                var val *Human = &Human{}
                data[i] = interface{}(val)
        }
    }
    
    printTypes(data)
}
