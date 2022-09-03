package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cmars/lsys"
)

var rules = map[string]struct {
	Seed      []byte
	Transform func(b byte) []byte
}{
	"algae":      {[]byte{'A'}, lsys.Algae},
	"tree":       {[]byte{'0'}, lsys.BinaryTree},
	"cantor":     {[]byte{'A'}, lsys.Cantor},
	"koch":       {[]byte{'F'}, lsys.Koch},
	"sierpinski": {[]byte{'F'}, lsys.Sierpinski},
}

var N = flag.Int("N", 3, "number of generations")
var rule = flag.String("rule", "algae", "L-system seed and rule")

func main() {
	flag.Parse()

	seedRule, ok := rules[*rule]
	if !ok {
		fmt.Println(*rule, "not found. available rules:")
		for ruleName := range rules {
			fmt.Println(ruleName)
		}
		os.Exit(1)
	}
	result := lsys.Seed(seedRule.Seed, seedRule.Transform)
	for i := 1; i < *N; i++ {
		result = lsys.Generate(result)
	}

	for result.Next() {
		fmt.Printf("%c", result.Value())
	}
	fmt.Println()
}
