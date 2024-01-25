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
	funcBlock, err := searchFunc(os.Args[1], os.Args[2], src)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", funcBlock)
	return nil
}

func searchFunc(fileName, funcName string, src []byte) ([]byte, error) {
	f, err := parser.ParseFile(token.NewFileSet(), fileName, src, 0)
	if err != nil {
		return nil, err
	}
	for _, d := range f.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			if fn.Name.Name == funcName {
				start, end := fn.Pos()-f.FileStart, fn.End()-f.FileStart
				return src[start:end], nil
			}
		}
	}
	return nil, fmt.Errorf("function: %s not found in file: %s", funcName, fileName)
}
