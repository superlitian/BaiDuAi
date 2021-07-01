/* Copyright 2021 CoderLee Inc. All Rights Reserved. */
/* aip_nlp.go 自然语言模块 */
/* modification history
-----------------------
2021/7/1, by coderlee, 创建
*/
/*
DESCRIPTION
	自然语言模块的封装，包含自然语言所有API
*/

package BaiDu_Ai

import (
	"encoding/json"
	"errors"
	"strings"
)

//Typedefs

// AipNlp 内容审核类
type AipNlp struct {
	appId       string
	apiKey      string
	secretKey   string
	accessToken map[string]string
}

/*
	AipNlp 初始化nlp模块
	PARAMS:
		- _appId: Ai应用的appid
		- _apiKey: Ai应用的api key
		- _secretKey Ai应用的secret key
	RETURNS:
		- error: 鉴权发生错误
		- &aipNlp: 内容审核实例
*/
func NewAipNlp(_appId string, _apiKey string, _secretKey string) (*AipNlp, error) {
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
	return &AipNlp{
		appId:       _appId,
		apiKey:      _apiKey,
		secretKey:   _secretKey,
		accessToken: t,
	}, nil
}
