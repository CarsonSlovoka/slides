<style>
  .current-fragment a {
    color: greenyellow;
  }
</style>

{{$themeSlice := list "black" "white" "league" "beige" "night" "serif" "simple" "solarized" "moon" "dracula" "sky" "blood"}}

{{$discuss := list
  (list "./example.md" "" "Example")
  (list "./layout.md?theme=sky" "" "Layout")
  (list "./media.md?theme=league" "white" "Media")
  (list "./slides-plugin.md?theme=beige" "" "Slides Plugin")
  (list "./test_raw.md?theme=beige&raw" "" "no gohtml")
}}

# Overview

<ul>
{{range $i, $item := $discuss}}
  {{$href := index $item 0}}
  {{$theme := index $item 1}}
  {{if eq (len $theme) 0}}
    {{$theme = index $themeSlice (mod $i (len $themeSlice))}}
  {{end}}
  {{$textContent := index $item 2}}
  {{printf "<li class='fragment'><a href='%s?theme=%s' target='_blank'>%s</a></li>" $href $theme $textContent }}
{{end}}
</ul>
