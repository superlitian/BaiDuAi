# BauDuAiHelp

百度人工智能模块go语言SDK，持续更新中

# 使用说明

- 使用示例（人脸检测接口）

```go
package main

import (
	"encoding/base64"
	"fmt"
	"github.com/iscoderLee/BaiDuAiHelp/face"
	"io/ioutil"
)

func main() {
  //此处为百度AI应用的鉴权信息
	appid,apikey,secretkey:="","",""
	client,err := face.NewAipFace(appid,apikey,secretkey)
	if err != nil{
		fmt.Println(err.Error())
	}
	ff,_ :=ioutil.ReadFile("本地图片路径")
	image := base64.StdEncoding.EncodeToString(ff)
	res:=client.Detect(image,"BASE64",nil)
	fmt.Println(res)
}
```

