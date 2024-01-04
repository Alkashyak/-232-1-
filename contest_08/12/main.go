package main

import (
    "fmt"
)

type Stack interface{
    empty() bool
    push(r rune)
    top() rune
    pop() rune
}

type SimpleStack []rune
func NewSimpleStack() *SimpleStack {
	return &SimpleStack{}
}

func (stk *SimpleStack) push(r rune) {
	*stk = append(*stk, r)
}
func (stk *SimpleStack) pop() rune {
	r := stk.top()
	*stk = (*stk)[:len(*stk)-1]
	return r
}
func (stk *SimpleStack) empty() bool {
	return len(*stk) == 0
}
func (stk *SimpleStack) top() rune {
	return (*stk)[len(*stk)-1]
}


func bracketValidator(stack Stack, str string) bool{
    mapping := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, char := range str {
        if char == '(' || char == '[' || char == '{' {
            stack.push(char)
        } else if char == ')' || char == ']' || char == '}' {
            if stack.empty() || stack.top() != mapping[char] {
                return false
            }
            stack.pop()
        }
    }

    return stack.empty()
}

func main() {
    var input string
    fmt.Scan(&input)
    fmt.Println(bracketValidator(NewSimpleStack(), input))
}
