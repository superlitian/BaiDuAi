# BauDuAi

百度人工智能模块go语言SDK，持续更新中

# 安装

> go get github.com/iscoderLee/BaiDuAi

# 使用说明

- 使用示例（人脸检测接口）

```go
package main

import (
   "encoding/base64"
   "fmt"
   "github.com/iscoderLee/BaiDuAi"
   "io/ioutil"
)

func main(){
   //百度智能云控制台获取的AI应用鉴权信息
   appid:=""
   apikey:=""
   secretkey:=""
   client,err:=BaiDu_Ai.NewAipFace(appid,apikey,secretkey)
   if err!=nil{
      fmt.Println(err.Error())
   }
   ff,_:=ioutil.ReadFile("本地图片路径")
   image := base64.StdEncoding.EncodeToString(ff)
   res:=client.Detect(image,"BASE64",nil)
   fmt.Println(res)
}
```

