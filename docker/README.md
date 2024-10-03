powershell執行範例

```shell
docker run -p 8080:8080 -e PORT=8080 -e MD_DIR="docs" `
  -e FS_DIRS="pages,static,tmpl" `
  -v .\docs:/usr/local/bin/docs `
  -v .\static:/usr/local/bin/static `
  -v .\tmpl/:/usr/local/bin/tmpl `
  --name SlidesCmd slides-cmd:v0.2.0
```

- `.\docs`: 將當前工作目錄下的docs目錄，綁定掛載到`/usr/local/bin/docs`目錄之中。(在Bind Mounts可以看到掛載的內容)