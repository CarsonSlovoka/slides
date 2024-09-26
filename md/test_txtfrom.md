{{printf "<style>%s</style>" (txtFrom
  (printf "%s%s" .BaseURL "/assets/pre.css")
  (printf "%s%s" .BaseURL "/assets/fragment.css")
)
}}

```go [|2|3-5|7|8|7-8]
type VheaTable struct {
Version              uint32 // 0x00010000 or 0x00110000
Ascent               int16
Descent              int16
LineGap              int16 // Reserved; set to 0
AHMax                uint16

MinTSB               int16
MinBSB               int16
YMaxExtent           int16
CaretSlopeRise       int16
CaretSlopeRun        int16
CaretOffset          int16
reserved1            int16
reserved2            int16
reserved3            int16
reserved4            int16
MetricDataFormat     int16

NumMetrics           uint16
}

```

---

- This is an <span class="fragment grow">apple</span>
- This is a <span class="fragment grow">pen</span>

Note:

透過fragment.css達成grow時顏色修改 `:is(li, a, span).current-fragment {color: #ff5a00;}`

---

{{tmpl (join "" .BaseURL "/assets/end_slides.md") (dict "Img" "/assets/home2.webp" "Href" "/md/example.md?theme=sky#3/2")}}

Note:

透過這種方式可以將指定的樣板匯入
