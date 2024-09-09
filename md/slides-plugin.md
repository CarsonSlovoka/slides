# Slides-Plugin

---

## Guess

<section data-auto-animate>

採用預設值:
```html [|2]
<div
  data-guess
>Hello World</div>
```

<div data-guess>Hello World</div>
</section>

<section data-auto-animate>

指定初始模糊程度:

```html [|2|3]
<div data-guess
     style="
     --val:20;
">Hello World</div>`
```

<div data-guess style="--val:20">Hello World</div>
</section>

<section data-auto-animate>

指定模糊程度與step:
```html [1: 2-3|4|2-4]
<div data-guess
     style="
     --val:20;
     --step:5;
">Hello World</div>
```
<div data-guess style="--val:20;--step:5">Hello World</div>
</section>

---

## Thanks
