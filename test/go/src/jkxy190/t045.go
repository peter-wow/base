package main

import (
    "fmt"
)


func main() {
    var a A = Work{3}    
    s := a.(Work)
    fmt.Println(s.ShowA())
    fmt.Println(s.ShowB())
}

type A interface {
    ShowA() int
}

type B interface {
    ShowB() int
}

type Work struct {
    i int
}

func (w Work) ShowA() int {
    return w.i + 10
}

func (w Work) ShowB() int {
    return w.i +20
}

