# Slides

主要透過[reveal.js](https://github.com/hakimel/reveal.js)來生成投影片

## install

```yaml
git clone https://github.com/CarsonSlovoka/slides.git
cd slides
git submodule update --init --recursive reveal.js
# go install -ldflags "-s -w" -tags tmpl  # 執行檔GOPATH/bin目錄
go build -ldflags "-s -w" -tags tmpl
```

> 建議打包的時候可以包含`-tags tmpl`，他可以將預設的[slides.gohtml](slides.gohtml)[嵌入](embed_tmpl.go)
>
> 使得當您不想要再對樣板修改時，也不在需要提供樣板

## USAGE

拿到此執行檔後，可以將執行檔的位置加入環境變數

接著在你想要的工作目錄，建立三種內容

1. 📂 md: 這是一個目錄，裡面放所有你想要投影的md檔案
2. (optional ) 📜 slides.gohtml, 這是一個模板你可以抓取[預設的內容後修改](slides.gohtml), 如果你不想準備此模板，請用`go build -tags tmpl`去生成執行檔
3. (optional ) 📂 assets: 如果你在md之中，有想要用img來加入本地的圖片，可以考慮將圖片路徑保存在此目錄。使用連結`/assets`會自動以此目錄為相對位置開始找尋

