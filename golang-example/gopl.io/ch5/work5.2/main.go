package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 练习 5.2： 编写函数，记录在HTML树中出现的同名元素的次数。

// 思路: map 递归

var count = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse: %v\n", err)
		os.Exit(1)
	}
	visit(doc)

	for t, c := range count {
		fmt.Println(t, c)
	}
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
