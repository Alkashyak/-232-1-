package main
import (
    "bufio"
    "fmt"
    "os"
)

import "strings"
func isPalindrome(booba string) bool {
	booba = strings.ReplaceAll(booba, "~", "")
	booba = strings.ReplaceAll(booba, "!", "")
	booba = strings.ReplaceAll(booba, "@", "")
	booba = strings.ReplaceAll(booba, "#", "")
	booba = strings.ReplaceAll(booba, "%", "")
	booba = strings.ReplaceAll(booba, "^", "")
	booba = strings.ReplaceAll(booba, "&", "")
	booba = strings.ReplaceAll(booba, "*", "")
	booba = strings.ReplaceAll(booba, "(", "")
	booba = strings.ReplaceAll(booba, ")", "")
	booba = strings.ReplaceAll(booba, " ", "")
	booba = strings.ToLower(booba)
	booba = strings.ReplaceAll(booba, "\n", "")
	n := len(booba)
	for i := 0; i < n/2; i++ {
		if booba[i] != booba[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
    line, _  := bufio.NewReader(os.Stdin).ReadString('\n')
    if isPalindrome( line[:len(line)-1] ){  // Обрезаем \n в конце строки
        fmt.Println("YES")
    }else{
        fmt.Println("NO")
    }
}
