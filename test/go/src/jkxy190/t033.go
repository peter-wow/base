package main

import (
    "fmt"
)

func main() {
       t := Teacher{} 
       t.ShowA()
}


type People struct {}


func (p *People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}

func (p *People) ShowB() {
    fmt.Println("showB")
}

type Teacher struct {
    People
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}
