//go:build tmpl

package main

import _ "embed"

//go:embed slides.gohtml
var slidesTmplBytes []byte
