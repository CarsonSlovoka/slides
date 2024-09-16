package flag

import "fmt"

type SliceFlags []string

// String is an implementation of the flag.Value interface
func (f *SliceFlags) String() string {
	return fmt.Sprintf("%v", *f)
}

// Set is an implementation of the flag.Value interface
func (f *SliceFlags) Set(value string) error {
	*f = append(*f, value)
	return nil
}
