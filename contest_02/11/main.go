package main
import "fmt"

const rows int = 9
const cols int = 9

func isValidSudoku(input [9][9]int) bool {
    numR := len(input)
    numC := len(input[0])
    fact := true

    for row := 0; row < numR; row++ {
        for col := 0; col < numC; col++ {
            val := input[row][col]
            for innerCol := col + 1; innerCol < numC; innerCol++ {
                if val == input[row][innerCol] {
                    fact = false
                    break
                }
            }
        }
    }

    for col := 0; col < numC; col++ {
        for row := 0; row < numR; row++ {
            val := input[row][col]
            for innerRow := row + 1; innerRow < numR; innerRow++ {
                if val == input[innerRow][col] {
                    fact = false
                    break
                }
            }
        }
    }

    return fact
}




func main() {
    var bord [rows][cols]int
    
    for row:=0; row < rows; row++ {
        for col:=0; col < cols; col++ {
            fmt.Scanf("%c", &bord[row][col])  // Считываем один символ
            bord[row][col] -= '0'  // Чтобы из ASCII кода символа получить цифру
        }
        fmt.Scanf("\n")
    }
    
    if isValidSudoku(bord){
        fmt.Println("YES")
    }else{
        fmt.Println("NO")
    }
}
