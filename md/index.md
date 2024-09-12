{{$discuss := list
  (list "./example.md" "Example")
  (list "./layout.md?theme=sky" "Layout")
  (list "./media.md?theme=league" "Media")
  (list "./slides-plugin.md?theme=beige" "Slides Plugin")
  (list "./test_raw.md?theme=beige&raw" "no gohtml")
}}

# Overview

<ul>
{{range $item := $discuss}}
  {{$href := index $item 0}}
  {{$textContent := index $item 1}}
  {{unsafeHTML ( printf "<li class='fragment'><a href=%q target='_blank'>%s</a></li>" $href $textContent )}}
{{end}}
</ul>

<style>
  .current-fragment a {
    color: greenyellow;
  }
</style>
