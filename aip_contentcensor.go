/* Copyright 2021 CoderLee Inc. All Rights Reserved. */
/* aip_contentcensor.go 内容审核模块 */
/* modification history
-----------------------
2021/7/1, by coderlee, 创建
*/
/*
DESCRIPTION
	内容审核模块的封装，包含内容审核所有API
*/

package BaiDu_Ai

import (
	"encoding/json"
	"errors"
	"strings"
)

//Typedefs

// AipContentCensor 内容审核类
type AipContentCensor struct {
	appId       string
	apiKey      string
	secretKey   string
	accessToken map[string]string
}

//functions

/*
	NewAipContentCensor 初始化内容审核模块
	PARAMS:
		- _appId: Ai应用的appid
		- _apiKey: Ai应用的api key
		- _secretKey Ai应用的secret key
	RETURNS:
		- error: 鉴权发生错误
		- &aipContentCensor: 内容审核实例
*/
func NewAipContentCensor(_appId string, _apiKey string, _secretKey string) (*AipContentCensor, error) {
	resp := Oauth(_apiKey, _secretKey)
	var data map[string]string
	json.Unmarshal([]byte(resp), &data)
	var keys = GetKeys(data)
	if !In("access_token", keys) {
		return nil, errors.New(resp)
	}
	var _accessToken = data["access_token"]
	var chartIndex = strings.Index(_accessToken, "-")
	var tem = _accessToken[chartIndex+1:]
	if tem != _appId {
		return nil, errors.New("appId错误，请前往控制台核实")
	}
	var t = map[string]string{
		"access_token": data["access_token"],
	}
	return &AipContentCensor{
		appId:       _appId,
		apiKey:      _apiKey,
		secretKey:   _secretKey,
		accessToken: t,
	}, nil
}

/*
	ImgCensor 图像审核参数为本地图片
	PARAMS:
		- image: 待审核图像Base64编码字符串，以图像文件形式请求时必填，图像要求base64后大于等于5kb，小于等于4M，最短边大于等于128像素，小于等于4096像素，支持的图片格式：PNG、JPG、JPEG、BMP、GIF（仅对首帧进行审核）、Webp、TIFF
		- options: 可选参数
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipContentCensor) ImgCensor(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(imgCensorUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	ImgCensorUrl 图像审核参数为url
	PARAMS:
		- imgUrl: 图像URL地址，以URL形式请求，图像Url需要做UrlEncode，图像要求base64后大于等于5kb，小于等于4M，最短边大于等于128像素，小于等于4096像素支持的图片格式：PNG、JPG、JPEG、BMP、GIF（仅对首帧进行审核）、Webp、TIFF
		- options: 可选参数
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipContentCensor) ImgCensorUrl(imgUrl string, options map[string]string) string {
	reqUrl = GetUrlBuild(imgCensorUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["imgUrl"] = imgUrl
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	TextCensor 文本审核
	PARAMS:
		- text: 待审核文本字符串
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipContentCensor) TextCensor(text string) string {
	reqUrl = GetUrlBuild(textCensorUrl, client.accessToken)
	var options = map[string]string{
		"text": text,
	}
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	VideoCensor 短视频审核
	PARAMS:
		- name: 视频名称
		- videoUrl: 视频主URL地址，若主Url无效或抓取失败，则依次抓取备用地址videoUrl2、videoUrl3、videoUrl4，若全部抓取失败则审核失败
		- extId: 视频在用户平台的唯一ID，方便人工审核结束时数据推送，用户利用此ID唯一锁定一条平台资源，若无可填写视频Url
		- options: 可选参数
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipContentCensor) VideoCensor(name string, videoUrl string, extId string, options map[string]string) string {
	reqUrl = GetUrlBuild(videoCensorUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["name"] = name
	options["videoUrl"] = videoUrl
	options["extId"] = extId
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	VoiceCensor 语音审核参数为base64
	PARAMS:
		- base64: 语音文件的base64编码，与url二选一，若都有按base64调用
		- videoUrl: 语音文件格式 语音文件的格式，pcm、wav、amr、m4a。不区分大小写，推荐pcm格式
		- options: 可选参数
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipContentCensor) VoiceCensor(base64 string, fmt string, options map[string]string) string {
	reqUrl = GetUrlBuild(voiceCensorUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["base64"] = base64
	options["fmt"] = fmt
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	VoiceCensorUrl 语音审核参数为url
	PARAMS:
		- url: 语音文件的url地址，示例：www.asd.com/asd.acc
		- videoUrl: 语音文件格式 语音文件的格式，pcm、wav、amr、m4a。不区分大小写，推荐pcm格式
		- options: 可选参数
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipContentCensor) VoiceCensorUrl(url string, fmt string, options map[string]string) string {
	reqUrl = GetUrlBuild(voiceCensorUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["fmt"] = fmt
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
