package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

// 练习 5.7： 完善startElement和endElement函数，
// 使其成为通用的HTML输出器。
// 要求：输出注释结点，
// 文本结点以及每个元素的属性（< a href='...'>）。
// 使用简略格式输出没有孩子结点的元素
// （即用<img/>代替<img></img>）。
// 编写测试，验证程序输出的格式正确。（详见11章）

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s=%q", a.Key, a.Val)
		}
		if n.Data == "img" && n.FirstChild == nil {
			fmt.Println("/>")
		} else {
			fmt.Println(">")
		}
		depth++
	case html.CommentNode:
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.Data != "img" || n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
