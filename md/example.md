# Slides

2024/09/22

Carson Tseng

---

## import image

relative path

![slides_icon.png](img/slides_icon.png)

----

## import image

from root

![go.png](/assets/go.png)

---

# Markdown

## heading

### h3

#### h4

##### h5

###### h6

----

## utils

- **bold**
- *italic*
- `code`

> blockquote

----

## Table

| Name | Desc  |
| ---- |:-----:|
go | ★★★★★ |
python |  ★★★  |

----

## CodeBlock

```go
package main

import "fmt"

func Example_Println() {
  fmt.Println("Hello World")

  // Output:
  // Hello World
}
```

---

# reveal.js

---

## animate


Note:

[Note](https://revealjs.com/speaker-view/#adding-the-speaker-notes-plugin)

> 動畫下一步是用 `↓`因此接`data-separator-vertical`會有問題

- note項目符號測試
- note2

---

### data-auto-animate

---

<section data-auto-animate>
  <h1>data-auto-animate</h1>
</section>

<section data-auto-animate>
  <h1 style="margin-top: 100px; color: red;">data-auto-animate</h1>
</section>

---

#### heading

<section data-auto-animate>
  <h1>Implicit</h1>
</section>

<section data-auto-animate>
  <h1>Implicit</h1>
  <h1>Animation</h1>
</section>

---

#### li

<section data-auto-animate>
  <ul>
    <li>Mercury</li>
    <li>Jupiter</li>
    <li>Mars</li>
  </ul>
</section>

<section data-auto-animate>
  <ul>
    <li>Mercury</li>
    <li>Earth</li>
    <li>Jupiter</li>
    <li>Saturn</li>
    <li>Mars</li>
  </ul>
</section>

---

### code-animation

---

<section data-auto-animate>
  <pre data-id="code-animation">
    <code data-trim data-line-numbers>

      let planets = [
        { name: 'mars', diameter: 6779 },
      ]
  </code>
  </pre>
</section>

<section data-auto-animate>
  <pre data-id="code-animation"><code data-trim data-line-numbers>

    let planets = [
      { name: 'mars', diameter: 6779 },
      { name: 'earth', diameter: 12742 },
      { name: 'jupiter', diameter: 139820 }
    ]
  </code>
  </pre>
</section>
<section data-auto-animate>
  <pre data-id="code-animation">
    <code data-trim data-line-numbers>

        let circumferenceReducer = ( c, planet ) => {
          return c + planet.diameter * Math.PI;
        }

        let planets = [
          { name: 'mars', diameter: 6779 },
          { name: 'earth', diameter: 12742 },
          { name: 'jupiter', diameter: 139820 }
        ]
        let c = planets.reduce( circumferenceReducer, 0 )

   </code>
  </pre>
</section>

---

## Thanks
