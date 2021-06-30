package main

import (
	BaiDu_Ai "BaiDu.Ai"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	appId := "23842866"
	apiKey := "mTdZaKel1W1pKkaIwalAUFxt"
	secretKey := "MuMc9lBsrlaO00wzinTkoGvARVLhLKbA"
	client, err := BaiDu_Ai.NewAipImage(appId, apiKey, secretKey)
	if err != nil {
		fmt.Println(err.Error())
	}
	ff, _ := ioutil.ReadFile("img/1.jpeg")
	image := base64.StdEncoding.EncodeToString(ff)
	image = "https://aip.bdstatic.com/portal-pc-node/dist/1591263471100/images/technology/imagerecognition/general/1.jpg"
	res := client.AdvancedGeneralUrl(image, nil)
	fmt.Println(res)
}
