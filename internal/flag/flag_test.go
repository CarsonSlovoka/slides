package flag_test

import (
	"flag"
	"fmt"
	flag2 "github.com/CarsonSlovoka/slides/internal/flag"
)

func ExampleSliceFlags_String() {
	flagSlice := flag2.SliceFlags{"a", "b"}
	fmt.Println(flagSlice.String())
	// Output:
	// [a b]
}

func ExampleSliceFlags_Set() {
	var flagSlice flag2.SliceFlags
	flag.Var(&flagSlice, "list", "-list 'b' -list 'b'")
	if err := flag.CommandLine.Parse([]string{
		"-list", "a",
		"-list", "b",
	}); err != nil {
		return
	}
	fmt.Println(flagSlice)
	// Output:
	// [a b]
}
