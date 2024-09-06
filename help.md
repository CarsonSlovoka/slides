## 查看你的幻燈片清單

> 👉 [/md](/md)

## USAGE

1. 在執行檔下的目錄放置[slides.gohtml](https://github.com/CarsonSlovoka/slides/blob/master/slides.gohtml)，你可以適當的修改
2. [md](https://github.com/CarsonSlovoka/slides/tree/master/md)目錄之中可以放想要投影的投影片內容
3. 訪問連結 `/demo/<位於md資料夾的md檔案名稱>`
    - example: [demo](/md/example.md)

> 有關於example.md的內容，可以[至此下載](https://github.com/CarsonSlovoka/slides/blob/master/md/example.md)後存放到md的目錄之中

## 可選項:

所有的可選項透過GET.query的參數設定，你可以將多個可選項用`&`組合起來

> demo: [md/example.md?theme=sky&autoSlide=5000](/md/example.md?theme=sky&autoSlide=5000)

### 下載成pdf

請於網址最後面補上`?print-pdf`，在使用列印(<kbd>Ctrl+P</kbd>)即可

```
?print-pdf
```

> demo: [md/example.md?print-pdf](/md/example.md?print-pdf)

### View

> demo: [md/example.md?view=scroll](/md/example.md?view=scroll)

### 變更主題顏色

theme=[?](https://github.com/hakimel/reveal.js/tree/472535065c7525abf0cc9df51c66f19fd2d2204f/dist/theme)

> demo: [md/example.md?theme=sky](/md/example.md?theme=sky)

### autoSlide

`?autoSlide=<毫秒>`

按下快捷鍵<kbd>A</kbd>可以Toggle

> demo: [md/example.md?autoSlide=5000](/md/example.md?autoSlide=5000)
