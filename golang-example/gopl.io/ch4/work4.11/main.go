package main

import (
	"example/gopl.io/ch4/work4.11/issue"
	"flag"
	"fmt"
	"os"
)

// 练习 4.11： 编写一个工具，允许用户在命令行创建、读取、更新和关闭GitHub上的issue，
// 当必要的时候自动打开用户默认的编辑器用于输入文本信息。

var (
	create = flag.Bool("c", false, "")
	list   = flag.Bool("l", false, "")
	read   = flag.Bool("r", false, "")
	edit   = flag.Bool("e", false, "")

	owner  = flag.String("owner", "", "")
	repo   = flag.String("repo", "", "")
	number = flag.String("number", "", "")
	token  = flag.String("token", "", "")

	title = flag.String("title", "", "")
	body  = flag.String("body", "", "")
)

func main() {
	flag.Parse()
	switch {
	case *list:
		p := issue.Params{Owner: *owner, Repo: *repo}
		issues, err := p.GetIssues()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		for _, i := range issues {
			fmt.Printf("%s\t%s\n", i.Title, i.Body)
		}
	case *read:
		p := issue.Params{Owner: *owner, Repo: *repo, Number: *number}
		i, err := p.GetIssue()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Printf("%s\t%s\n", i.Title, i.Body)
	case *create:
		p := issue.Params{Owner: *owner, Repo: *repo, Token: *token, Issue: issue.Issue{Title: *title, Body: *body}}
		if !p.CreateIssue() {
			fmt.Fprintf(os.Stderr, "create issue fail")
		}
	case *edit:
		p := issue.Params{Owner: *owner, Repo: *repo, Token: *token, Number: *number, Issue: issue.Issue{Title: *title, Body: *body}}
		if !p.EditIssue() {
			fmt.Fprintf(os.Stderr, "edit issue fail")
		}
	}
}
