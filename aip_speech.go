/* Copyright 2021 CoderLee Inc. All Rights Reserved. */
/* aip_speech.go 语音模块 */
/* modification history
-----------------------
2021/7/8, by coderlee, 创建
*/
/*
DESCRIPTION
	语音模块的封装，包含语音识别，语音合成API
*/

package BaiDu_Ai

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

//Typedefs

// AipSpeech 语音类
type AipSpeech struct {
	appId       string
	apiKey      string
	secretKey   string
	accessToken map[string]string
}

//functions

/*
	NewAipSpeech 初始化人体分析模块
	PARAMS:
		- _appId: Ai应用的appid
		- _apiKey: Ai应用的api key
		- _secretKey Ai应用的secret key
	RETURNS:
		- error: 鉴权发生错误
		- &aipBodyAnalysis: 人体分析实例
*/
func NewAipSpeech(_appId string, _apiKey string, _secretKey string) (*AipSpeech, error) {
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
	return &AipSpeech{
		appId:       _appId,
		apiKey:      _apiKey,
		secretKey:   _secretKey,
		accessToken: t,
	}, nil
}

/*
	Asr 短语音识别标准版
	PARAMS:
		- format: 语音文件的格式，pcm/wav/amr/m4a。不区分大小写。推荐pcm文件
		- cuid: 用户唯一标识，用来区分用户，计算UV值。建议填写能区分用户的机器 MAC 地址或 IMEI 码，长度为60字符以内。
		- speech: 本地语音文件的二进制语音数据
		- options: 可选参数
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipSpeech) Asr(format string, cuid string, speech []byte, options map[string]interface{}) string {
	if options == nil {
		options = map[string]interface{}{}
	}
	var speechLen = len(speech)
	var speechB64 = base64.StdEncoding.EncodeToString(speech)
	options["format"] = format
	options["rate"] = 16000
	options["channel"] = 1
	options["cuid"] = cuid
	options["token"] = client.accessToken["access_token"]
	options["speech"] = speechB64
	options["len"] = speechLen

	host = NewHttpSend(asrUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	Synthesis 语音合成
	PARAMS:
		- tex: 要合成的文本
		- savePath: 要写入的文件名称，可加前缀路径 例如：file/fileName.pcm
		- cuid: 用户标识
		- options: 可选参数
	RETURNS:
		- string, 写入成功或接口失败消息
*/
func (client *AipSpeech) Synthesis(tex string, savePath string, cuid string, options map[string]string) string {
	if options == nil {
		options = map[string]string{}
	}
	options["tex"] = tex
	options["cuid"] = cuid
	options["tok"] = client.accessToken["access_token"]
	options["ctp"] = "1"
	options["lan"] = "zh"
	host = NewHttpSend(synthesisUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	if strings.Index("err_no", string(res)) == -1 {
		return string(res)
	}
	err = ioutil.WriteFile(savePath, res, 0777)
	if err != nil {
		return err.Error()
	}
	return "success"
}
