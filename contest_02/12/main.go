package main

import "fmt"

func fill(input [][]int) {
    numR := len(input)
    numC := len(input[0])
    for row := 0; row < numR; row++ {
        for col := 0; col < numC; col++ {
            if input[row][col] == -1 {
                if col > 0 && input[row][col-1] != -1 {
                    input[row][col-1]++
                }
                if col < numC-1 && input[row][col+1] != -1 {
                    input[row][col+1]++
                }
                if row > 0 {
                    if input[row-1][col] != -1 {
                        input[row-1][col]++
                    }
                    if col > 0 && input[row-1][col-1] != -1 {
                        input[row-1][col-1]++
                    }
                    if col < numC-1 && input[row-1][col+1] != -1 {
                        input[row-1][col+1]++
                    }
                }
                if row < numR -1 {
                    if input[row+1][col] != -1 {
                        input[row+1][col]++
                    }
                    if col > 0 && input[row+1][col-1] != -1 {
                        input[row+1][col-1]++
                    }
                    if col < numC-1 && input[row+1][col+1] != -1 {
                        input[row+1][col+1]++
                    }
                }
            }
        }
    }
}

func main() {
    var rows, cols int
    fmt.Scanf("%d %d\n", &rows, &cols)
    
    // Создаём срез и заполняем его данными о расположении мин
    maze := make([][]int, rows, rows)
    for i := range maze {
        maze[i] = make([]int, cols, cols)
        for j := range maze[i] {
            fmt.Scanf("%d", &maze[i][j])
        }
    }
    
    // Заполняем игровое поле подсказками
    fill(maze)
    
    // Выводим на экран
    for _, row := range maze {
         for _, cell := range row {
             fmt.Printf("%3d", cell)
        }
        fmt.Println()
    }
}
