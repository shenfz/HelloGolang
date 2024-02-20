package main

/**
 * @Author shenfz
 * @Date 2022/1/21 10:11
 * @Email 1328919715@qq.com
 * @Description: 通过 go vet 拓展自定义检查器 ，比如检查某些禁用文字
 **/

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
	"strings"
)

var analyzer = &analysis.Analyzer{
	Name: "checkWords",
	Doc:  "check forbidden words usage in our strings",
	Run:  run,
}

var (
	forbiddenWords = []string{
		"bird",
		"water",
		"candy",
	}
)

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.BasicLit:
				if x.Kind != token.STRING {
					return false
				}
				words := strings.Fields(x.Value)
				for _, word := range words {
					for _, forbiddenWord := range forbiddenWords {
						if word == forbiddenWord {
							pass.Reportf(x.Pos(), "Forbidden word used, please do not use the word =>[ %s ]<= in your strings", word)
							/*
								Test\example.go:13:14: Forbidden word used, please do not use the word =>[ candy ]<= in your strings
							*/
						}
					}
				}
				return false
			}
			return true
		})
	}
	return nil, nil
}

//go:generate go build -o checkwords.exe
//go:generate go vet -vettool=./checkwords.exe -checkWords ./Test
func main() {
	unitchecker.Main(
		analyzer,
	)
}
