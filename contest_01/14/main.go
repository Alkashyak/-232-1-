package main
import "fmt"
import "strconv"
func main() {
var row, col int
fmt.Scan(&row)
fmt.Scan(&col)
 fmt.Print("    |")
 for i := 1; i <= col; i += 1 {
  if i < 10 {
   fmt.Print("   ", i)
  } else {
   fmt.Print("  ", i)
  }
 }
 fmt.Print("\n")
 fmt.Print("   --")
 for i := 1; i <= col; i += 1 {
  fmt.Print("----")
 }
 fmt.Print("\n")
 spacelen := 4
 for i := 1; i <= row; i += 1 {
  if i < 10 {
   fmt.Print("   ", i, "|")
  } else {
   fmt.Print("  ", i, "|")
  }
  for j := 1; j <= col; j += 1 {
   ij := (spacelen - len(strconv.Itoa(i*j)))
   for c := 1; c <= ij; c += 1 {
    fmt.Print(" ")
   }
   fmt.Print(i * j)
  }
  fmt.Print("\n")
 }
}
