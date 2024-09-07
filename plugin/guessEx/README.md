## 原理

利用css的filter功能來達到模糊的效果

```css
.demo {
  filter: blur(5px) grayscale(var(--grayscale,10%))
}
```

## Usage:

對所有guessEx選擇器中的含有`blur-elem`的元素進行模糊處理

```html
<div data-guess-ex data-bk-img="https://picsum.photos/id/1015/600/400" style="width:100%; height:200px;">
  <div class="blur-elem"  <!-- 第一個被模糊的對象 -->
     style="
    --blur: 20px; <!-- 霧化的初始值 -->
    --blurStep: 2; <!-- 每次點擊滑鼠左鍵blur會減少的數值 -->
    --grayscale: 80%;
    --grayscaleStep: 10; <!-- 每次點擊滑鼠左鍵--grayscale會減少的數值 -->
    clip-path: polygon(50% 32%, 0 0, 100% 100%);
  "></div>

  <div class="blur-elem" style=" <!-- 第二個被模糊的對象 -->
    --blur: 20px;
    --blurStep: 2;
    --grayscale: 80%;
    --grayscaleStep: 10;
    clip-path: polygon(60% 10%, 100% 38%, 82% 100%, 85% 42%, 70% 48%);
  "></div>
</div>
```

## Clip-path

利用檢查元素，去查圖片的大小，再到[clip-path](https://bennettfeely.com/clippy/)的網站

輸入好寬高，右邊可以選擇`Custom Polygon`就可以任意添加多個點去框出自己想要的區域
