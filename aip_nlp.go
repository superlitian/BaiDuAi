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

// AipNlp 自然语言类
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
		- &aipNlp: 自然语言实例
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

/*
	Lexer 词法分析通用版
	PARAMS:
		- text: 待分析文本，长度不超过20000字节
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipNlp) Lexer(text string) string {
	reqUrl = GetUrlBuild(lexerUrl, client.accessToken)
	var options = map[string]string{
		"text": text,
	}
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	LexerCustom 词法分析定制版
	PARAMS:
		- text: 待分析文本，长度不超过20000字节
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipNlp) LexerCustom(text string) string {
	reqUrl = GetUrlBuild(lexerCustomUrl, client.accessToken)
	var options = map[string]string{
		"text": text,
	}
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	WordEmbVec 词向量表示
	PARAMS:
		- word: 文本内容，最大64字节
		- options: 可选参数
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipNlp) WordEmbVec(word string, options map[string]string) string {
	reqUrl = GetUrlBuild(wordEmbVecUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["word"] = word
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	WordEmbSim 词义相似度
	PARAMS:
		- word1: 词1，最大64字节
		- word2: 词2，最大64字节
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipNlp) WordEmbSim(word1 string, word2 string) string {
	reqUrl = GetUrlBuild(wordEmbSimUrl, client.accessToken)
	var options = map[string]string{
		"word_1": word1,
		"word_2": word2,
	}
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	DnnlmCn DNN语言模型
	PARAMS:
		- text: 文本内容，最大256字节，不需要切词
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipNlp) DnnlmCn(text string) string {
	reqUrl = GetUrlBuild(dnnlmCnUrl, client.accessToken)
	var options = map[string]string{
		"text": text,
	}
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	Depparser 依存句法分析
	PARAMS:
		- text: 待分析文本，长度不超过128字符
	RETURNS:
		- string, 接口返回的消息
*/
func (client *AipNlp) Depparser(text string) string {
	reqUrl = GetUrlBuild(depparserUrl, client.accessToken)
	var options = map[string]string{
		"text": text,
	}
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	Simnet 短文本相似度
	PARAMS:
		- text1: 待比较文本1，最大512字节
		- text2: 待分析文本，待比较文本2，最大512字节
		- options: 可选参数
	RETURNS:
		- string, 接口返回的消息
*/
func (client *AipNlp) Simnet(text1 string, text2 string, options map[string]string) string {
	reqUrl = GetUrlBuild(simnetUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["text_1"] = text1
	options["text_2"] = text2
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	Keyword 文章标签
	PARAMS:
		- title: 文章标题，最大80字节
		- content: 文章内容，最大65535字节
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipNlp) Keyword(title string, content string) string {
	reqUrl = GetUrlBuild(keywordUrl, client.accessToken)
	var options = map[string]string{
		"title":   title,
		"content": content,
	}
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	Topic 文章分类
	PARAMS:
		- title: 文章标题，最大80字节
		- content: 文章内容，最大65535字节
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipNlp) Topic(title string, content string) string {
	reqUrl = GetUrlBuild(topicUrl, client.accessToken)
	var options = map[string]string{
		"title":   title,
		"content": content,
	}
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	Ecnet 文本纠错
	PARAMS:
		- text: 待纠错文本，输入限制511字节
	RETURNS:
		- string,接口返回消息
*/
func (client *AipNlp) Ecnet(text string) string {
	reqUrl = GetUrlBuild(ecnetUrl, client.accessToken)
	var options = map[string]string{
		"text": text,
	}
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	NewsSummary 新闻摘要
	PARAMS:
		- content: 字符串（限3000字符数以内）
字符串仅支持GBK编码，长度需小于3000字符数（即6000字节），请输入前确认字符数没有超限，若字符数超长会返回错误。正文中如果包含段落信息，请使用"\n"分隔，段落信息算法中有重要的作用，请尽量保留
		- maxSummaryLen: 此数值将作为摘要结果的最大长度。例如：原文长度1000字，本参数设置为150，则摘要结果的最大长度是150字；推荐最优区间：200-500字
		- options: 可选参数， : title
	RETURNS:
		- string,接口返回消息
*/
func (client *AipNlp) NewsSummary(content string, maxSummaryLen int, options map[string]interface{}) string {
	reqUrl = GetUrlBuild(newsSummaryUrl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["content"] = content
	options["max_summary_len"] = maxSummaryLen
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	CommentTag 评论观点抽取通用版
	PARAMS:
		- text:评论内容，最大10240字节
		- options: 可选参数：
						- type: int,评论行业类型，默认为4（餐饮美食）
	RETURNS:
		- string,接口返回消息
*/
func (client *AipNlp) CommentTag(text string, options map[string]interface{}) string {
	reqUrl = GetUrlBuild(commentTagUrl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["text"] = text
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	CommentTag 评论观点抽取定制版
	PARAMS:
		- text:评论内容，最大10240字节
		- options: 可选参数：
						- type: int,评论行业类型，默认为4（餐饮美食）
	RETURNS:
		- string,接口返回消息
*/
func (client *AipNlp) CommentTagCustom(text string, options map[string]interface{}) string {
	reqUrl = GetUrlBuild(commentTagCustomUrl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["text"] = text
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	SentimentClassify 情感倾向分析通用版
	PARAMS:
		- text: 文本内容，最大2048字节
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipNlp) SentimentClassify(text string) string {
	reqUrl = GetUrlBuild(sentimentClassifyUrl, client.accessToken)
	var options = map[string]string{
		"text": text,
	}
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	SentimentClassifyCustom 情感倾向分析定制版
	PARAMS:
		- text: 文本内容，最大2048字节
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipNlp) SentimentClassifyCustom(text string) string {
	reqUrl = GetUrlBuild(sentimentClassifyCustomUrl, client.accessToken)
	var options = map[string]string{
		"text": text,
	}
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	Emotion 对话情绪识别
	PARAMS:
		- text: 待识别情感文本，输入限制512字节
		- options: 可选参数
				- scene: string,default（默认项-不区分场景），talk（闲聊对话-如度秘聊天等），task（任务型对话-如导航对话等），customer_service（客服对话-如电信/银行客服等）
	RETURNS:
		- string, 接口返回消息
*/
func (client *AipNlp) Emotion(text string, options map[string]string) string {
	reqUrl = GetUrlBuild(emotionUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["text"] = text
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	Address 地址识别
	PARAMS:
		- text: 待识别的文本内容，不超过1000字节
		- options: 可选参数
				- confidence: int,取值100-0，不设置时默认为50。该字段用于触发补充解析策略，对置信度在配置值以下的结果，进行补充解析，以提高结果精度。该字段配置会增加服务耗时。经评测，在保证准确率提升效果的前提下，当取值=50时，服务平响增长相对较小。也可根据业务数据评测，决定取值。
	RETURNS:
		- string,接口返回消息
*/
func (client *AipNlp) Address(text string, options map[string]interface{}) string {
	reqUrl = GetUrlBuild(addressUrl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["text"] = text
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
