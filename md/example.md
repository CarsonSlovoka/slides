# Slides

2024/09/05

Carson Tseng

---

## import image

relative path

![slides_icon.png](img/slides_icon.png)

----

## import image

from root

![go.png](/assets/go.png)


----

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

func init() {
  fmt.Println("Hello World")

  // Output:
  // Hello World
}
```

---

# reveal.js

---

## Content

----

### Code

<pre><code data-line-numbers="3-5|8-10|13-15">
<table>
  <tr>
    <td>Apples</td>
    <td>$1</td>
    <td>7</td>
  </tr>
  <tr>
    <td>Oranges</td>
    <td>$2</td>
    <td>18</td>
  </tr>
  <tr>
    <td>Kiwi</td>
    <td>$3</td>
    <td>1</td>
  </tr>
</table>
</code></pre>

Note:

~~有scrollbar的時候會怪怪的~~，請此請確定相關的css, font有被載入

----

<pre><code data-line-numbers data-ln-start-from="7">
<tr>
  <td>Oranges</td>
  <td>$2</td>
  <td>18</td>
</tr>
</code></pre>

---

## Backgrounds


> ❗ 使用`data-background-color`不要用於vertical

---

`data-background-color="rgb(70, 70, 255)"`

<section data-background-color="rgb(70, 70, 255)">
  <h2>🍰</h2>
</section>

---

`data-background-gradient="linear-gradient(to bottom, #283b95, #17b2c3)"`

<section data-background-gradient="linear-gradient(to bottom, #283b95, #17b2c3)">
  <h2>🐟</h2>
</section>

---

`data-background-gradient="radial-gradient(#283b95, #17b2c3)"`

<section data-background-gradient="radial-gradient(#283b95, #17b2c3)">
  <h2>🐳</h2>
</section>

---

## Features

----

### [Jump to Slide](https://revealjs.com/jump-to-slide/)

<span id="jump-to"></span>

Press <key>G</key>, type a slide `number` or `id`

| Input   | Result |
|---------| ---- |
| 5       | Navigate to slide number 5
| 11/2    | Navigate to horizontal slide 11, vertical slide 2
| 11.2    | same as above
| the-end | Navigate to a slide with this id (for example: `<section id="the-end">`, `<p id="the-end">`)


Note:

使用id的時候，不需要打完全名也可以搜尋

----

## [Auto-Slide](https://revealjs.com/auto-slide/)

<p class="fragment" data-autoslide="20000">

It's also possible to pause or resume sliding by pressing <kbd>A</kbd>

</p>

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

# Review

- <a href="#/slides">Slides</a>
- <a href="#/markdown">Markdown</a>
- <a href="#/revealjs">reveal.js</a>
  - <a href="#/jump-to">Jump To Slide</a>

Note:

internal links

> 注意markdown, heading生成的id，都是小寫的, `.`會被移除，
>
> 例如: <a href="#/revealjs">reveal.js</a>

---

<h2 id="the-end">Thanks</h2>
