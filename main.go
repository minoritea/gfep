package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
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
	var index int
	flag.IntVar(&index, "i", 1, "index to use when multiple functions with same name are present")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		return nil
	}
	var (
		src []byte
		err error
	)
	if args[0] == "-" {
		src, err = io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
	} else {
		src, err = os.ReadFile(args[0])
		if err != nil {
			return err
		}
	}
	index--
	if index < 0 {
		index = 0
	}
	funcBlock, err := searchFunc(args[0], args[1], src, index)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", funcBlock)
	return nil
}

func searchFunc(fileName, funcName string, src []byte, index int) ([]byte, error) {
	f, err := parser.ParseFile(token.NewFileSet(), fileName, src, 0)
	if err != nil {
		return nil, err
	}
	var i int
	for _, d := range f.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			if fn.Name.Name == funcName {
				if i == index {
					start, end := fn.Pos()-f.FileStart, fn.End()-f.FileStart
					return src[start:end], nil
				}
				i++
			}
		}
	}
	return nil, fmt.Errorf("function: %s not found in file: %s", funcName, fileName)
}
