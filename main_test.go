package main

import (
	"fmt"
	"testing"
)

func TestSearchFunc(t *testing.T) {
	const code = `func SearchTarget() {
	println("Hello, World!")
}`

	src := []byte(fmt.Sprintf(`package main
%s
`, code))

	block, err := searchFunc("target.go", "SearchTarget", src)
	if err != nil {
		t.Fatalf("failed to search function: error: %s", err)
	}
	if string(block) != code {
		t.Fatalf("failed to search function:\n\nexpected:\n\n%s\n\nactual:\n\n%s", code, string(block))
	}
}
