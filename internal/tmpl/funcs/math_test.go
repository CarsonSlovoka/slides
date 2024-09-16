package funcs_test

import (
	"github.com/CarsonSlovoka/slides/internal/tmpl/funcs"
	"os"
	"text/template"
)

func ExampleAdd() {
	body := `
{{printf "%d" (add 1 3 -2)}}
{{printf "%0.1f" (add 1.0 3.5 -2.0)}}
`
	tmpl, err := template.New("").Funcs(map[string]any{
		"add": funcs.Add,
	}).Parse(body)
	if err != nil {
		return
	}
	if err = tmpl.Execute(os.Stdout, nil); err != nil {
		return
	}
	// Output:
	// 2
	// 2.5
}

func ExampleMod() {
	body := `
{{ mod (len (list 1 3 5)) 3  }}
{{ mod (len (list 1 3 5 -1 2)) 3  }}
`
	tmpl, err := template.New("").Funcs(map[string]any{
		"list": funcs.List,
		"mod":  funcs.Mod,
	}).Parse(body)
	if err != nil {
		return
	}
	if err = tmpl.Execute(os.Stdout, nil); err != nil {
		return
	}
	// Output:
	// 0
	// 2
}

func ExampleMul() {
	body := `
{{ mul 3 2  }}
{{ mul 3 2 8  }}
`
	tmpl, err := template.New("").Funcs(map[string]any{
		"mul": funcs.Mul,
	}).Parse(body)
	if err != nil {
		return
	}
	if err = tmpl.Execute(os.Stdout, nil); err != nil {
		return
	}
	// Output:
	// 6
	// 48
}

func ExamplePow() {
	body := `
{{ pow 3 4 }}
`
	tmpl, err := template.New("").Funcs(map[string]any{
		"pow": funcs.Pow,
	}).Parse(body)
	if err != nil {
		return
	}
	if err = tmpl.Execute(os.Stdout, nil); err != nil {
		return
	}
	// Output:
	// 81
}

func ExampleMax() {
	body := `
{{ max 3 4 9 1 }}
{{ max 100.1 4.0 9.0 1.0 }}
`
	tmpl, err := template.New("").Funcs(map[string]any{
		"max": funcs.Max,
	}).Parse(body)
	if err != nil {
		return
	}
	if err = tmpl.Execute(os.Stdout, nil); err != nil {
		return
	}
	// Output:
	// 9
	// 100.1
}

func ExampleMin() {
	body := `
{{ min 3 4 9 1 }}
{{ printf "%0.2f" (min 100.1 4.0 9.0 2.0) }}
`
	tmpl, err := template.New("").Funcs(map[string]any{
		"min": funcs.Min,
	}).Parse(body)
	if err != nil {
		return
	}
	if err = tmpl.Execute(os.Stdout, nil); err != nil {
		return
	}
	// Output:
	// 1
	// 2.00
}
