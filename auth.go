/* Copyright 2021 CoderLee Inc. All Rights Reserved. */
/* auth.go - 鉴权模块 */
/* modification history
-----------------------
2021/6/30, by coderlee, 创建
*/
/*
DESCRIPTION
	验证ai应用的鉴权信息
*/

package BaiDu_Ai

func Oauth(apiKey string, secretKey string) string {
	param := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     apiKey,
		"client_secret": secretKey,
	}
	h := NewHttpSend(GetUrlBuild(authurl, param))
	res, err = h.Get()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
