package BaiDu_Ai

const requrl = "https://aip.baidubce.com/oauth/2.0/token"

func Oauth(apiKey string, secretKey string) string {
	param := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     apiKey,
		"client_secret": secretKey,
	}
	h := NewHttpSend(GetUrlBuild(requrl, param))
	res, err := h.Get()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
