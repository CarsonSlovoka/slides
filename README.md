# Slides

主要透過[reveal.js](https://github.com/hakimel/reveal.js)來生成投影片

## install

```yaml
git clone https://github.com/CarsonSlovoka/slides.git
cd slides
git submodule update --init --recursive reveal.js
# go install -ldflags "-s -w"  # 執行檔GOPATH/bin目錄
go build -ldflags "-s -w"
```

## USAGE

拿到此執行檔後，可以將執行檔的位置加入環境變數

接著在你想要的工作目錄，建立三種內容

1. 📜slides.gohtml, 這是一個模板你可以抓取[預設的內容後修改](slides.gohtml)
2. 📂md: 這是一個目錄，裡面放所有你想要投影的md檔案
3. 📂assets: 如果你在md之中，有想要用img來加入本地的圖片，可以考慮將圖片路徑保存在此目錄。使用連結`/assets`會自動以此目錄為相對位置開始找尋

