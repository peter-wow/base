package main

import (
    "fmt"
)

func main() {
    slice := make([]int, 5, 5)    
    slice[0] = 1
    slice[1] = 2
    change(slice...)
    
    fmt.Println(len(slice), cap(slice))
    
    fmt.Println(slice)
    change(slice[0:2]...)
    fmt.Println(slice)
}

func change(s ...int) {
    s = append(s, 3)
}

