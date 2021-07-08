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

func main() {
  //百度智能云控制台获取的AI应用鉴权信息
  appid := ""
  apikey := ""
  secretkey := ""
  client, err := BaiDu_Ai.NewAipFace(appid, apikey, secretkey)
  if err != nil {
    fmt.Println(err.Error())
  }
  ff, _ := ioutil.ReadFile("本地图片路径")
  image := base64.StdEncoding.EncodeToString(ff)
  res := client.Detect(image, "BASE64", nil)
  fmt.Println(res)

  // 带可选参数调用
  ff, _ := ioutil.ReadFile("本地图片路径")
  image := base64.StdEncoding.EncodeToString(ff)
  options := map[string]string{
    "face_field": "age,beauty",
  }
  res := client.Detect(image, "BASE64", options)
  fmt.Println(res)
}
```

- 语音示例

  - 语音识别

  > ```go
  > package main
  > 
  > import (
  > 	BaiDu_Ai "BaiDu.Ai"
  > 	"fmt"
  > 	"io/ioutil"
  > )
  > 
  > func main() {
  > 	appid := "24494480"
  > 	apikey := "ACpt2CbFtbH7AKqtxGgpirMl"
  > 	secretkey := "yYjrV7H156RkAzyNIGEgpae1wFXOSDpf"
  > 	client, err := BaiDu_Ai.NewAipSpeech(appid, apikey, secretkey)
  > 	if err!=nil {
  > 		fmt.Println(err)
  > 	}
  > 	speech := ioutil.ReadFile("本地音频文件路径")
  > 
  > 	res:= client.Asr("pcm","cuid",speech,nil)
  > 
  > 	fmt.Println(res)
  >   
  >   //带可选参数调用
  >   options := map[string]interface{}{
  >     "dev_pid":"1837", //四川话模型
  >   }
  >   res:= client.Asr("pcm","cuid",speech,options)
  >   fmt.Println(res)
  > }
  > 
  > ```

  - 语音合成

  >
  >
  >```go
  >package main
  >
  >import (
  >	BaiDu_Ai "BaiDu.Ai"
  >	"fmt"
  >)
  >
  >func main() {
  >	appid := ""
  >	apikey := ""
  >	secretkey := ""
  >	client, err := BaiDu_Ai.NewAipSpeech(appid, apikey, secretkey)
  >	if err!=nil {
  >		fmt.Println(err)
  >	}
  >	var options = map[string]string{
  >		"per":"103",
  >		"aue":"3",
  >	}
  >	
  >  //要保存的文件路径
  >  filePath = "file/result.mp3"
  >  
  >  //待合成文本
  >  
  >  text = "百度你好"
  >  
  >	res:= client.Synthesis(text,filePath,"cuid",options)
  >
  >	fmt.Println(res)
  >}
  >```

>
>
