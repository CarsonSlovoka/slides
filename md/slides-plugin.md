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

## GuessEx

- [clip-path](https://bennettfeely.com/clippy/)

<div data-guess-ex data-bk-img="https://picsum.photos/id/1015/600/400" style="width:100%; height:200px;">
  <div class="blur-elem"
     style="
    --blur: 20px;
    --blurStep: 2;
    --grayscale: 80%;
    --grayscaleStep: 10;
    clip-path: polygon(50% 32%, 0 0, 100% 100%);
  "></div>

  <div class="blur-elem" style="
    --blur: 20px;
    --blurStep: 2;
    --grayscale: 80%;
    --grayscaleStep: 10;
    clip-path: polygon(60% 10%, 100% 38%, 82% 100%, 85% 42%, 70% 48%);
  "></div>
</div>

Note:

利用檢查元素，去查圖片的大小，再到[clip-path](https://bennettfeely.com/clippy/)的網站

輸入好寬高，右邊可以選擇`Custom Polygon`就可以任意添加多個點去框出自己想要的區域

---

## [Thanks](./index.md)
