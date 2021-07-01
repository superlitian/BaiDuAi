/* Copyright 2021 CoderLee Inc. All Rights Reserved. */
/* aip_body_analysis.go 人体分析模块 */
/* modification history
-----------------------
2021/7/1, by coderlee, 创建
*/
/*
DESCRIPTION
	人体模块的封装，包含人体分析所有API
*/

package BaiDu_Ai

import (
	"encoding/json"
	"errors"
	"strings"
)

// Constants

//Typedefs

// AipBodyAnalysis 人体分析类
type AipBodyAnalysis struct {
	appId       string
	apiKey      string
	secretKey   string
	accessToken map[string]string
}

//functions

/*
	NewAipBodyAnalysis 初始化人体分析模块
	PARAMS:
		- _appId: Ai应用的appid
		- _apiKey: Ai应用的api key
		- _secretKey Ai应用的secret key
	RETURNS:
		- error: 鉴权发生错误
		- &aipBodyAnalysis: 人体分析实例
*/
func NewAipBodyAnalysis(_appId string, _apiKey string, _secretKey string) (*AipBodyAnalysis, error) {
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
	return &AipBodyAnalysis{
		appId:       _appId,
		apiKey:      _apiKey,
		secretKey:   _secretKey,
		accessToken: t,
	}, nil
}

/*
	BodyAnalysis 人体关键点识别
	PARAMS:
		- image: 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M。图片的base64编码是不包含图片头的，如(data:image/jpg;base64,)，支持图片格式：jpg、bmp、png，最短边至少50px，最长边最大4096px
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipBodyAnalysis) BodyAnalysis(image string) string {
	reqUrl = GetUrlBuild(bodyAnalysisUrl, client.accessToken)
	var options = map[string]string{
		"image": image,
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
	BodyAttr 人体检测和属性识别
	PARAMS:
		- image: 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M。图片的base64编码是不包含图片头的，如(data:image/jpg;base64,)，支持图片格式：jpg、bmp、png，最短边至少50px，最长边最大4096px
		- options: 可选参数
	RETURNS:
		- string, 接口返消息
*/
func (client *AipBodyAnalysis) BodyAttr(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(bodyAttrUrl, client.accessToken)
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
	BodyNum 人流量统计
	PARAMS:
		- image: 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M。图片的base64编码是不包含图片头的，如(data:image/jpg;base64,)，支持图片格式：jpg、bmp、png，最短边至少50px，最长边最大4096px
		- options: 可选参数
	RETURNS:
		- string, 接口返消息
*/
func (client *AipBodyAnalysis) BodyNum(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(bodyNumUrl, client.accessToken)
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
	Gesture 手势识别
	PARAMS:
		- image: 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M。图片的base64编码是不包含图片头的，如(data:image/jpg;base64,)，支持图片格式：jpg、bmp、png，最短边至少50px，最长边最大4096px
	RETURNS:
		- string, 接口返消息
*/
func (client *AipBodyAnalysis) Gesture(image string) string {
	reqUrl = GetUrlBuild(gestureUrl, client.accessToken)
	var options = map[string]string{
		"image": image,
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
	BodySeg 人像分割
	PARAMS:
		- image: 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M。图片的base64编码是不包含图片头的，如(data:image/jpg;base64,)，支持图片格式：jpg、bmp、png，最短边至少50px，最长边最大4096px
		- options: 可选参数
	RETURNS:
		- string, 接口返消息
*/
func (client *AipBodyAnalysis) BodySeg(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(bodySegUrl, client.accessToken)
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
	DriverBehavior 驾驶行为分析
	PARAMS:
		- image: 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M。图片的base64编码是不包含图片头的，如(data:image/jpg;base64,)，支持图片格式：jpg、bmp、png，最短边至少50px，最长边最大4096px
		- options: 可选参数
	RETURNS:
		- string, 接口返消息
*/
func (client *AipBodyAnalysis) DriverBehavior(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(driverBehaviorUrl, client.accessToken)
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
	BodyTracking 人流量统计（动态版）
	PARAMS:
		- dynamic: true,false; true：动态人流量统计，返回总人数、跟踪ID、区域进出人数； false：静态人数统计，返回总人数
		- image: 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M。图片的base64编码是不包含图片头的，如(data:image/jpg;base64,)，支持图片格式：jpg、bmp、png，最短边至少50px，最长边最大4096px
		- options: 可选参数
	RETURNS:
		- string, 接口返消息
*/
func (client *AipBodyAnalysis) BodyTracking(dynamic string, image string, options map[string]string) string {
	reqUrl = GetUrlBuild(bodyTrackingUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["dynamic"] = dynamic
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	HandAnalysis 手部关键点识别
	PARAMS:
		- image: 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M。图片的base64编码是不包含图片头的，如(data:image/jpg;base64,)，支持图片格式：jpg、bmp、png，最短边至少50px，最长边最大4096px
	RETURNS:
		- string, 接口返消息
*/
func (client *AipBodyAnalysis) HandAnalysis(image string) string {
	reqUrl = GetUrlBuild(handAnalysisUrl, client.accessToken)
	var options = map[string]string{
		"image": image,
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
	BodyDanger 危险行为识别
	PARAMS:
		- data: 视频数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M，支持mp4、 mov格式，5s以内的监控视频片段
	RETURNS:
		- string, 接口返消息
*/
func (client *AipBodyAnalysis) BodyDanger(data string) string {
	reqUrl = GetUrlBuild(bodyDangerUrl, client.accessToken)
	var options = map[string]string{
		"data": data,
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
	Fingertip 指尖检测
	PARAMS:
		- image: 图像数据，base64编码后进行urlencode，要求base64编码和urlencode后大小不超过4M。图片的base64编码是不包含图片头的，如(data:image/jpg;base64,)，支持图片格式：jpg、bmp、png，最短边至少50px，最长边最大4096px。

	RETURNS:
		- string, 接口返消息
*/
func (client *AipBodyAnalysis) Fingertip(image string) string {
	reqUrl = GetUrlBuild(fingertipUrl, client.accessToken)
	var options = map[string]string{
		"image": image,
	}
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
