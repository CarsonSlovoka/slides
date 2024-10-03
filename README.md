# Slides

主要透過[reveal.js](https://github.com/hakimel/reveal.js)來生成投影片

## Download

您可以自己[Build](#build)，如果不想要建立，可以到[release的頁面下](https://github.com/CarsonSlovoka/slides/releases)下載對應平台的zip檔案，裡面有build完成的執行檔

## Build

```yaml
git clone https://github.com/CarsonSlovoka/slides.git
cd slides
git submodule update --init --recursive reveal.js
git submodule update --init --recursive plugin/guess
git submodule update --init --recursive plugin/marker
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

    注意! slides.gohtml裡面對於md的分隔符號是`\n`，如果你想要改成`\r\n`可以[調整為](https://github.com/CarsonSlovoka/slides/blob/84bcf2f776ffc87d4d96f051ad3a2da856b43123/slides.gohtml#L16-L17)

    ```html
    data-separator="^\r\n---\r\n"
    data-separator-vertical="^\r\n----\r\n"
    ```


3. (optional ) 📂 assets: 如果你在md之中，有想要用img來加入本地的圖片，可以考慮將圖片路徑保存在此目錄。使用連結`/assets`會自動以此目錄為相對位置開始找尋

啟動上可以直接
```yaml
slides.exe
slides.exe -port=12345 # port預設用8080, 如果被佔用或者不喜歡可以改成其他，或者指定為0會自動抓取

# 如果你不喜歡把幻燈片都保存在md的目錄，你可以自定義您的目錄名稱，啟動上使用-md來告知
slides.exe -md="mySlides"
```

或者用https

```
slides.exe -tls
```

## Docker

提供兩種方法

1. 使用[carsonTseng/slides](https://hub.docker.com/layers/carsontseng/slides/v0.2.0/images/sha256-99243e1e2dbc0ec22ad8e1b849cbdc6340ca51f80b7dbe2d4444f3a70693fc3d): [把slides當作image](https://github.com/CarsonSlovoka/slides/blob/b5166bdd88e1af3a89651e399a744f6c807dc2f9/docker/Dockerfile.example#L10), 所以你還要自己額外寫Dockerfile才可以看到結果，可以參考[Dockerfile.example](docker/Dockerfile.example)
2. 使用[carsonTseng/slides-cmd](https://hub.docker.com/repository/docker/carsontseng/slides-cmd/general): 好處是你不需要再寫Dockerfile，可以利用綁定掛載(Bind mounts)的方式,來讀取到要被渲染的內容，請參考以下範例

    ```yaml
    # 使用powershell
    docker run -p 8080:8080 -e PORT=8080 `
      -e MD_DIR="docs" ` # 放到md的所在目錄
      -v .\docs:/usr/local/bin/docs ` # 後面的docs需與MD_DIR所指定的目錄相同
      -v .\slides.gohtml/:/usr/local/bin/slides.gohtml ` # 可選，如果你需要自定義slides.gohtml也可以告知
      -v .\assets:/usr/local/bin/assets ` # 可選(assets目錄，預設就會被當成fs，不需要再加到FS_DIRS去)
      -e FS_DIRS="pages,static,tmpl" ` # 即 -fs pages -f static -fs tmpl
      -v .\pages:/usr/local/bin/pages ` # 這些內容對應於 FS_DIRS # 其中的.\docs應該於您目前的工作目錄中能找到
      -v .\static:/usr/local/bin/static ` # 同上
      -v .\tmpl/:/usr/local/bin/tmpl ` # 同上
      --name my-slides-example carsontseng/slides-cmd:v0.2.0
    ```


## Plugin 自定義插件

可以參考[此commit](https://github.com/CarsonSlovoka/slides/commit/b239af8f9b9ffcf27bbb8b00e46e9f2fb516cf47)

裡面製作了一個[guessEx的插件](https://github.com/CarsonSlovoka/slides/blob/b239af8f9b9ffcf27bbb8b00e46e9f2fb516cf47/plugin/guessEx/guessEx.mjs)

當你製作好plugin之後，可以在工作目錄下建立一個plugin的目錄，將其內容放入[即可被程式抓到](https://github.com/CarsonSlovoka/slides/blob/68d94531130c50deb653260bcb094f11cf03071b/main.go#L212)

之後[建立屬於您自己的slides.gohtml，導入該plugin](https://github.com/CarsonSlovoka/slides/blob/68d94531130c50deb653260bcb094f11cf03071b/slides.gohtml#L47)即可

### [3-rd plugins](https://github.com/hakimel/reveal.js/wiki/Plugins,-Tools-and-Hardware)

