package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	src, err := os.ReadFile(os.Args[1])
	if err != nil {
		return err
	}
	f, err := parser.ParseFile(token.NewFileSet(), os.Args[1], src, 0)
	if err != nil {
		return err
	}
	for _, d := range f.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			if fn.Name.Name == os.Args[2] {
				start, end := fn.Pos()-f.FileStart, fn.End()-f.FileStart
				fmt.Printf("%s", src[start:end])
			}
		}
	}
	return nil
}
