/* Copyright 2021 CoderLee Inc. All Rights Reserved. */
/* aip_ocr.go - OCR模块 */
/* modification history
-----------------------
2021/6/30, by coderlee, 创建
*/
/*
DESCRIPTION
	OCR模块各接口的封装，包含OCR所有场景
*/

package BaiDu_Ai

import (
	"encoding/json"
	"errors"
	"strings"
)

type AipOcr struct {
	appId       string
	apiKey      string
	secretKey   string
	accessToken map[string]string
}

func NewAipOcr(_appId string, _apiKey string, _secretKey string) (*AipOcr, error) {
	resp := Oauth(_apiKey, _secretKey)
	data := map[string]string{}
	json.Unmarshal([]byte(resp), &data)
	keys := GetKeys(data)
	if !In("access_token", keys) {
		return nil, errors.New(resp)
	}
	_accessToken := data["access_token"]
	chartIndex := strings.Index(_accessToken, "-")
	tem := _accessToken[chartIndex+1:]
	if tem != _appId {
		return nil, errors.New("appId错误，请前往控制台核实")
	}
	t := map[string]string{
		"access_token": data["access_token"],
	}
	return &AipOcr{
		appId:       _appId,
		apiKey:      _apiKey,
		secretKey:   _secretKey,
		accessToken: t,
	}, nil
}

/*
	通用文字识别，图片参数为本地图片
*/
func (client *AipOcr) BasicGeneral(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(basicGeneralurl, client.accessToken)
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
	通用文字识别, 图片参数为远程url图片
*/
func (client *AipOcr) BasicGeneralUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(basicGeneralurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别, 图片参数为PDF
*/
func (client *AipOcr) BasicGeneralPDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(basicGeneralurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（含位置信息版），参数为本地图片
*/
func (client *AipOcr) General(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(generalurl, client.accessToken)
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
	通用文字识别（含位置信息版），参数为url
*/
func (client *AipOcr) GeneralUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(generalurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（含位置信息版），参数为PDF
*/
func (client *AipOcr) GeneralPDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(generalurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（高精度版）参数为本地图片
*/
func (client *AipOcr) BasicAccurate(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(basicAccurateurl, client.accessToken)
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
	通用文字识别（高精度版）参数为url
*/
func (client *AipOcr) BasicAccurateUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(basicAccurateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（高精度版）参数为pdf_file
*/
func (client *AipOcr) BasicAccuratePDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(basicAccurateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（高精度含位置版）参数为本地图片
*/
func (client *AipOcr) Accurate(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(accurateurl, client.accessToken)
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
	通用文字识别（高精度含位置版）参数为url
*/
func (client *AipOcr) AccurateUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(accurateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（高精度含位置版）参数为PDF
*/
func (client *AipOcr) AccuratePDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(accurateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	办公文档识别 参数为image
*/

func (client *AipOcr) DocAnalysisOffice(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(docanalysisofficeurl, client.accessToken)
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
	办公文档识别 参数为url
*/

func (client *AipOcr) DocAnalysisOfficeUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(docanalysisofficeurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	办公文档识别 参数为PDF
*/

func (client *AipOcr) DocAnalysisOfficePDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(docanalysisofficeurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	网络图片文字识别 参数为image
*/

func (client *AipOcr) WebImage(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(webImageurl, client.accessToken)
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
	网络图片文字识别 参数为url
*/

func (client *AipOcr) WebImageUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(webImageurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	网络图片文字识别（含位置版）
	参数为本地图片
*/

func (client *AipOcr) WebImageLoc(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(webimageLocurl, client.accessToken)
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
	网络图片文字识别（含位置版）
	参数为url
*/

func (client *AipOcr) WebImageLocUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(webimageLocurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	手写文字识别
	参数为本地图片
*/

func (client *AipOcr) HandWriting(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(handwritingurl, client.accessToken)
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
	手写文字识别
	参数为url
*/

func (client *AipOcr) HandWritingUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(handwritingurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	数字识别
	参数为本地图片
*/

func (client *AipOcr) Numbers(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(numberurl, client.accessToken)
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
	数字识别
	参数为url
*/

func (client *AipOcr) NumbersUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(numberurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	表格文字识别同步接口
	参数为本地图片
*/
func (client *AipOcr) Form(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(formurl, client.accessToken)
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
	表格文字识别同步接口
	参数为url
*/
func (client *AipOcr) FormUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(formurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	表格文字识别
*/

func (client *AipOcr) TableRecognition(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(tableRecognitionAsyncurl, client.accessToken)
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
	获取表格文字识别结果
*/

func (client *AipOcr) GetTableRecognitionResult(requestId string, options map[string]string) string {
	requlr := GetUrlBuild(getTableRecognitionResulturl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["request_id"] = requestId
	host = NewHttpSend(requlr)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	二维码识别
	参数为本地图片
*/
func (client *AipOcr) Qrcode(image string) string {
	reqUrl = GetUrlBuild(qrcodeurl, client.accessToken)
	options := map[string]string{
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
	二维码识别
	参数为url
*/
func (client *AipOcr) QrcodeUrl(url string) string {
	reqUrl = GetUrlBuild(qrcodeurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	身份证识别
	参数为本地图片
*/

func (client *AipOcr) IdCard(image string, idCardSide string, options map[string]string) string {
	reqUrl = GetUrlBuild(idcardurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["id_card_side"] = idCardSide
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	身份证识别
	参数为url
*/

func (client *AipOcr) IdCardUrl(url string, idCardSide string, options map[string]string) string {
	reqUrl = GetUrlBuild(idcardurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["id_card_side"] = idCardSide
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	身份证混贴识别
	参数为本地图片
*/
func (client *AipOcr) MultiIdcard(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(multiidcardurl, client.accessToken)
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
	银行卡识别
	参数为本地图片
*/
func (client *AipOcr) Bankcard(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(bankcardurl, client.accessToken)
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
	银行卡识别
	参数为url
*/
func (client *AipOcr) BankcardUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(bankcardurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	营业执照识别
	参数为本地图片
*/
func (client *AipOcr) BusinessLicense(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(businessLicenseurl, client.accessToken)
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
	营业执照识别
	参数为url
*/
func (client *AipOcr) BusinessLicenseUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(businessLicenseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	名片识别
	参数为本地图片
*/
func (client *AipOcr) BusinessCard(image string) string {
	reqUrl = GetUrlBuild(businesscardurl, client.accessToken)
	options := map[string]string{
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
	名片识别
	参数为url
*/
func (client *AipOcr) BusinessCardUrl(url string) string {
	reqUrl = GetUrlBuild(businesscardurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	护照识别
	参数为本地图片
*/
func (client *AipOcr) Passport(image string) string {
	reqUrl = GetUrlBuild(passporturl, client.accessToken)
	options := map[string]string{
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
	护照识别
	参数为url
*/
func (client *AipOcr) PassportUrl(url string) string {
	reqUrl = GetUrlBuild(passporturl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	港澳通行证识别
	参数为本地图片
*/
func (client *AipOcr) HKMacauExitentrypermit(image string) string {
	reqUrl = GetUrlBuild(HKMacauexitentrypermiturl, client.accessToken)
	options := map[string]string{
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
	港澳通行证识别
	参数为url
*/
func (client *AipOcr) HKMacauExitentrypermitUrl(url string) string {
	reqUrl = GetUrlBuild(HKMacauexitentrypermiturl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	台湾通行证识别
	参数为本地图片
*/
func (client *AipOcr) TaiWanExitentrypermit(image string) string {
	reqUrl = GetUrlBuild(taiwanexitentrypermiturl, client.accessToken)
	options := map[string]string{
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
	台湾通行证识别
	参数为url
*/
func (client *AipOcr) TaiWanExitentrypermitUrl(url string) string {
	reqUrl = GetUrlBuild(taiwanexitentrypermiturl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	户口本识别
	参数为本地图片
*/
func (client *AipOcr) HouseholdRegister(image string) string {
	reqUrl = GetUrlBuild(householdRegisterurl, client.accessToken)
	options := map[string]string{
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
	户口本识别
	参数为url
*/
func (client *AipOcr) HouseholdRegisterUrl(url string) string {
	reqUrl = GetUrlBuild(householdRegisterurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	出生医学证明识别
*/
func (client *AipOcr) BirthCertificate(image string) string {
	reqUrl = GetUrlBuild(birthcertificateurl, client.accessToken)
	options := map[string]string{
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
	多卡证类别检测
	参数为本地图片
*/
func (client *AipOcr) MultiCardClassify(image string) string {
	reqUrl = GetUrlBuild(multicardclassifyurl, client.accessToken)
	options := map[string]string{
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
	多卡证类别检测
	参数为url
*/
func (client *AipOcr) MultiCardClassifyUrl(url string) string {
	reqUrl = GetUrlBuild(multicardclassifyurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	增值税发票识别
	参数为本地图片
*/
func (client *AipOcr) VatInvoice(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(vatInvoiceurl, client.accessToken)
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
	增值税发票识别
	参数为url
*/
func (client *AipOcr) VatInvoiceUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(vatInvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	增值税发票识别
	参数为PDF
*/
func (client *AipOcr) VatInvoicePDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(vatInvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	增值税发票验真
*/
func (client *AipOcr) VatInvoiceVerification(invoiceCode string, invoiceNum string, invoiceDate string, invoiceType string, checkCode string, totalAmount string) string {
	reqUrl = GetUrlBuild(vatinvoiceverificationurl, client.accessToken)
	options := map[string]string{
		"invoice_code": invoiceCode,
		"invoice_num":  invoiceNum,
		"invoice_date": invoiceDate,
		"invoice_type": invoiceType,
		"check_code":   checkCode,
		"total_amount": totalAmount,
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
	定额发票识别
	参数为本地图片
*/
func (client *AipOcr) QuotaInvoice(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(quotainvoiceurl, client.accessToken)
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
	定额发票识别
	参数为url
*/
func (client *AipOcr) QuotaInvoiceUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(quotainvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	定额发票识别
	参数为pdf
*/
func (client *AipOcr) QuotaInvoicePDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(quotainvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用机打发票识别
	参数为本地图片
*/
func (client *AipOcr) Invoice(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(invoiceurl, client.accessToken)
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
	通用机打发票识别
	参数为url
*/
func (client *AipOcr) InvoiceUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(invoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用机打发票识别
	参数为PDF
*/
func (client *AipOcr) InvoicePDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(invoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	火车票识别
	参数为本地图片
*/
func (client *AipOcr) TrainTicket(image string, optios map[string]string) string {
	reqUrl = GetUrlBuild(trainTicketurl, client.accessToken)
	if optios == nil {
		optios = map[string]string{}
	}
	optios["image"] = image
	host = NewHttpSend(reqUrl)
	host.SetBody(optios)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	火车票识别
	参数为url
*/
func (client *AipOcr) TrainTicketUrl(url string, optios map[string]string) string {
	reqUrl = GetUrlBuild(trainTicketurl, client.accessToken)
	if optios == nil {
		optios = map[string]string{}
	}
	optios["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(optios)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	火车票识别
	参数为PDF
*/
func (client *AipOcr) TrainTicketPDF(pdfFile string, optios map[string]string) string {
	reqUrl = GetUrlBuild(trainTicketurl, client.accessToken)
	if optios == nil {
		optios = map[string]string{}
	}
	optios["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(optios)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	出租车票识别
	参数为本地图片
*/
func (client *AipOcr) TaxiReceipt(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(taxiReceipturl, client.accessToken)
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
	出租车票识别
	参数为url
*/
func (client *AipOcr) TaxiReceiptUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(taxiReceipturl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	出租车票识别
	参数为PDF
*/
func (client *AipOcr) TaxiReceiptPDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(taxiReceipturl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	飞机行程单识别
	参数为本地图片
*/
func (client *AipOcr) AirTicket(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(airTicketurl, client.accessToken)
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
	飞机行程单识别
	参数为url
*/
func (client *AipOcr) AirTicketUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(airTicketurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	飞机行程单识别
	参数为PDF
*/
func (client *AipOcr) AirTicketUrlPDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(airTicketurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	汽车票识别
	参数为本地图片
*/
func (client *AipOcr) BusTicket(image string) string {
	reqUrl = GetUrlBuild(busticketurl, client.accessToken)
	options := map[string]string{
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
	汽车票识别
	参数为url
*/
func (client *AipOcr) BusTicketUrl(url string) string {
	reqUrl = GetUrlBuild(busticketurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	过路过桥费发票识别
	参数为本地图片
*/
func (client *AipOcr) TollInvoice(image string) string {
	reqUrl = GetUrlBuild(tollinvoiceurl, client.accessToken)
	options := map[string]string{
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
	过路过桥费发票识别
	参数为url
*/
func (client *AipOcr) TollInvoiceUrl(url string) string {
	reqUrl = GetUrlBuild(tollinvoiceurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	船票识别
	参数为本地图片
*/
func (client *AipOcr) FerryTicket(image string) string {
	reqUrl = GetUrlBuild(ferryticketurl, client.accessToken)
	options := map[string]string{
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
	船票识别
	参数为url
*/
func (client *AipOcr) FerryTicketUrl(url string) string {
	reqUrl = GetUrlBuild(ferryticketurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	通用票据识别
	参数为本地图片
*/
func (client *AipOcr) Receipt(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(receipturl, client.accessToken)
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
	通用票据识别
	参数为url
*/
func (client *AipOcr) ReceiptUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(receipturl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	网约车行程单识别
	参数为本地图片
*/
func (client *AipOcr) OnlineTaxiItinerary(image string) string {
	reqUrl = GetUrlBuild(onlinetaxiitineraryurl, client.accessToken)
	options := map[string]string{
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
	网约车行程单识别
	参数为url
*/
func (client *AipOcr) OnlineTaxiItineraryUrl(url string) string {
	reqUrl = GetUrlBuild(onlinetaxiitineraryurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	医疗发票识别
	参数为本地图片
*/
func (client *AipOcr) MedicalInvoice(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(medicalinvoiceurl, client.accessToken)
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
	医疗发票识别
	参数为url
*/
func (client *AipOcr) MedicalInvoiceUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(medicalinvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	医疗费用结算单识别
	参数为本地图片
*/
func (client *AipOcr) MedicalStatement(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(medicalstatementurl, client.accessToken)
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
	医疗费用结算单识别
	参数为url
*/
func (client *AipOcr) MedicalStatementUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(medicalstatementurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	医疗费用明细识别
	参数为本地图片
*/
func (client *AipOcr) MedicalDetail(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(medicaldetailurl, client.accessToken)
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
	医疗费用明细识别
	参数为url
*/
func (client *AipOcr) MedicalDetailUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(medicaldetailurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	病案首页识别
	参数为本地图片
*/
func (client *AipOcr) MedicalRecord(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(medicalrecordurl, client.accessToken)
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
	病案首页识别
	参数为url
*/
func (client *AipOcr) MedicalRecordUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(medicalrecordurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	保单识别
	参数为本地图片
*/
func (client *AipOcr) InsuranceDocuments(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(insurancedocumentsurl, client.accessToken)
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
	保单识别
	参数为url
*/
func (client *AipOcr) InsuranceDocumentsUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(insurancedocumentsurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	行驶证识别
	参数为本地图片
*/
func (client *AipOcr) VehicleLicense(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(vehicleLicenseurl, client.accessToken)
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
	行驶证识别
	参数为url
*/
func (client *AipOcr) VehicleLicenseUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(vehicleLicenseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	驾驶证识别
	参数为本地图片
*/
func (client *AipOcr) DrivingLicense(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(drivingLicenseurl, client.accessToken)
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
	驾驶证识别
	参数为url
*/
func (client *AipOcr) DrivingLicenseUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(drivingLicenseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	车牌识别
	参数为本地图片
*/
func (client *AipOcr) LicensePlate(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(licensePlateurl, client.accessToken)
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
	车牌识别
	参数为url
*/
func (client *AipOcr) LicensePlateUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(licensePlateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	VIN码识别
	参数为本地图片
*/
func (client *AipOcr) VinCode(image string) string {
	reqUrl = GetUrlBuild(vinCodeurl, client.accessToken)
	options := map[string]string{
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
	VIN码识别
	参数为url
*/
func (client *AipOcr) VinCodeUrl(url string) string {
	reqUrl = GetUrlBuild(vinCodeurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	车辆合格证识别
	参数为本地图片
*/
func (client *AipOcr) VehicleCertificate(image string) string {
	reqUrl = GetUrlBuild(vehicleCertificateurl, client.accessToken)
	options := map[string]string{
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
	车辆合格证识别
	参数为url
*/
func (client *AipOcr) VehicleCertificateUrl(url string) string {
	reqUrl = GetUrlBuild(vehicleCertificateurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	机动车登记证书识别
	参数为本地图片
*/
func (client *AipOcr) VehicleRegistrationCertificate(image string) string {
	reqUrl = GetUrlBuild(vehicleregistrationcertificateurl, client.accessToken)
	options := map[string]string{
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
	机动车登记证书识别
	参数为url
*/
func (client *AipOcr) VehicleRegistrationCertificateUrl(url string) string {
	reqUrl = GetUrlBuild(vehicleregistrationcertificateurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	机动车销售发票识别
	参数为本地图片
*/
func (client *AipOcr) VehicleInvoice(image string) string {
	reqUrl = GetUrlBuild(vehicleInvoiceurl, client.accessToken)
	options := map[string]string{
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
	机动车销售发票识别
	参数为url
*/
func (client *AipOcr) VehicleInvoiceUrl(url string) string {
	reqUrl = GetUrlBuild(vehicleInvoiceurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	二手车销售发票识别
	参数为本地图片
*/
func (client *AipOcr) UsedVehicleInvoice(image string) string {
	reqUrl = GetUrlBuild(usedvehicleinvoiceurl, client.accessToken)
	options := map[string]string{
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
	二手车销售发票识别
	参数为url
*/
func (client *AipOcr) UsedVehicleInvoiceUrl(url string) string {
	reqUrl = GetUrlBuild(usedvehicleinvoiceurl, client.accessToken)
	options := map[string]string{
		"url": url,
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
	试卷分析与识别
	参数为本地图片
*/
func (client *AipOcr) DocAnalysis(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(docAnalysisurl, client.accessToken)
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
	试卷分析与识别
	参数为url
*/
func (client *AipOcr) DocAnalysisUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(docAnalysisurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	试卷分析与识别
	参数为pdf
*/
func (client *AipOcr) DocAnalysisPDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(docAnalysisurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	公式识别
	参数为本地图片
*/
func (client *AipOcr) Formula(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(formulaurl, client.accessToken)
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
	公式识别
	参数为url
*/
func (client *AipOcr) FormulaUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(formulaurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	仪器仪表盘读数识别
	参数为本地图片
*/
func (client *AipOcr) Meter(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(meterurl, client.accessToken)
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
	仪器仪表盘读数识别
	参数为url
*/
func (client *AipOcr) MeterUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(meterurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	印章识别
	参数为本地图片
*/
func (client *AipOcr) Seal(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(sealurl, client.accessToken)
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
	印章识别
	参数为url
*/
func (client *AipOcr) SealUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(sealurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	印章识别
	参数为pdf
*/
func (client *AipOcr) SealPDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(sealurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	彩票识别
	参数为本地图片
*/
func (client *AipOcr) Lottery(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(lotteryurl, client.accessToken)
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
	彩票识别
	参数为url
*/
func (client *AipOcr) LotteryUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(lotteryurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	门脸文字识别
*/
func (client *AipOcr) Facade(image string) string {
	reqUrl = GetUrlBuild(facadeurl, client.accessToken)
	options := map[string]string{
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
	智能结构化识别
	参数为本地图片
*/
func (client *AipOcr) IntelligentOcr(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(intelligentocrurl, client.accessToken)
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
	智能结构化识别
	参数为url
*/
func (client *AipOcr) IntelligentOcrUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(intelligentocrurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	iOCR通用版
	参数为本地图片
*/
func (client *AipOcr) Recognise(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(recogniseurl, client.accessToken)
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
	iOCR通用版
	参数为url
*/
func (client *AipOcr) RecogniseUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(recogniseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	iOCR通用版
	参数为pdf
*/
func (client *AipOcr) RecognisePDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(recogniseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	iOCR财会版
	参数为本地图片
*/
func (client *AipOcr) Finance(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(financeurl, client.accessToken)
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
	iOCR财会版
	参数为url
*/
func (client *AipOcr) FinanceUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(financeurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	iOCR财会版
	参数为pdf
*/
func (client *AipOcr) FinancePDF(pdfFile string, options map[string]string) string {
	reqUrl = GetUrlBuild(financeurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
