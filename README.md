# Bing\-Wallpaper\-RESTful

The original repository: [TimothyYe/bing-wallpaper: A RESTful API to fetch daily wallpaper from Bing.com (github.com)](https://github.com/TimothyYe/bing-wallpaper)\.

## Build 

```bash
go build -o bw.exe
```

## Run

```bash
bw run
```

```bash
bw -h
```

## Docker Build 

```bash
docker build -t 4thrun/bing-wallpaper-restful .
```

## Docker Pull

```
docker pull 4thrun/bing-wallpaper-restful:latest
```

## Docker Run 

```bash
docker run -d -p 9002:9002 4thrun/bing-wallpaper-restful
```

## HTTP GET

### API

- Path: `/`
- Method: GET

### Parameters

- `resolution`: \[`1920`, `1366`, `3840`\]

  The resolution of wallpaper image. `1920` is the default value, you can also use `1366` and `3840`(4K resolution)\.

- `format`: \[`json`, `image`\]

  The response format, can be `json` or `image`. **If response format is set to `image`, you will be redirected to the wallpaper image directly**.

- `index`: [`0`, +∞)

  The index of wallpaper, starts from `0`\.

- `mkt`: \[`zh-CN`, `en-US`, `ja-JP`, `en-AU`, `en-UK`, `de-DE`, `en-NZ`, `en-CA`\]

  The region parameter, default is `zh-CN`, you can also use `en-US`, `ja-JP`, `en-AU`, `en-UK`, `de-DE`, `en-NZ`, `en-CA`.

### Example 

Request:

```
http://127.0.0.1:9002/?resolution=1920&format=json&index=0&mkt=zh-CN
```

Response:

```json
{"sdate":"20210830","edate":"20210831","url":"https://cn.bing.com/th?id=OHR.DjurdjevicaBridge_ZH-CN0284105882_1920x1080.jpg","copyright":"杜德维卡塔拉大桥，黑山 (© Hike The World/Shutterstock)","copyright_link":"https://www.bing.com/search?q=%E6%9D%9C%E5%BE%B7%E7%BB%B4%E5%8D%A1%E5%A1%94%E6%8B%89%E5%A4%A7%E6%A1%A5\u0026form=hpcapt\u0026mkt=zh-cn"}
```

