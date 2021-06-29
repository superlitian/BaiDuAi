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

//url

//通用文字识别，参数为本地图片
const basicGeneralurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"

//通用文字识别高精度版
const basicAccurateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"

//通用文字识别（含位置信息版）
const generalurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"

//调用通用文字识别（含位置高精度版）
const accurateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate"

//网络图片文字识别
const webImageurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage"

//身份证识别
const idcardurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard"

//身份证混贴识别
const multiidcardurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/multi_idcard"

//银行卡识别
const bankcardurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard"

//驾驶证识别
const drivingLicenseurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/driving_license"

//行驶证识别
const vehicleLicenseurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_license"

//车牌识别
const licensePlateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate"

//营业执照识别
const businessLicenseurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/business_license"

//名片识别
const businesscardurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/business_card"

//通用票据识别
const receipturl = "https://aip.baidubce.com/rest/2.0/ocr/v1/receipt"

//自定义模板文字识别
const customurl = ""

//表格文字识别同步接口
const formurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/form"

//表格文字识别
const tableRecognitionAsyncurl = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/request"

//表格识别结果
const getTableRecognitionResulturl = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/get_request_result"

//试卷分析与识别
const docAnalysisurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/doc_analysis"

//仪表盘读数识别
const meterurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/meter"

//网络图片文字识别（含位置版）
const webimageLocurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage_loc"

//增值税发票识别
const vatInvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice"

//增值税发票验真
const vatinvoiceverificationurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice_verification"

//定额发票识别
const quotainvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/quota_invoice"

//出租车票识别
const taxiReceipturl = "https://aip.baidubce.com/rest/2.0/ocr/v1/taxi_receipt"

//VIN码识别
const vinCodeurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vin_code"

//火车票识别
const trainTicketurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/train_ticket"

//数字识别
const numberurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/numbers"

//飞机行程单识别
const airTicketurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/air_ticket"

//二维码识别
const qrcodeurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/qrcode"

//手写文字识别
const handwritingurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/handwriting"

//护照识别
const passporturl = "https://aip.baidubce.com/rest/2.0/ocr/v1/passport"

//通用机打发票识别
const invoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/invoice"

//户口本识别
const householdRegisterurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/household_register"

//机动车销售发票识别
const vehicleInvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_invoice"

//车辆合格证识别
const vehicleCertificateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_certificate"

//港澳通行证识别
const HKMacauexitentrypermiturl = "https://aip.baidubce.com/rest/2.0/ocr/v1/HK_Macau_exitentrypermit"

//台湾通行证识别
const taiwanexitentrypermiturl = "https://aip.baidubce.com/rest/2.0/ocr/v1/taiwan_exitentrypermit"

//出生医学证明识别
const birthcertificateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/birth_certificate"

//多卡证类别检测
const multicardclassifyurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/multi_card_classify"

//汽车票识别
const busticketurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/bus_ticket"

//过来过桥费发票识别
const tollinvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/toll_invoice"

//船票识别
const ferryticketurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/ferry_ticket"

//网约车行程单识别
const onlinetaxiitineraryurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/online_taxi_itinerary"

//医疗发票识别
const medicalinvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/medical_invoice"

//医疗费用结算单识别
const medicalstatementurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/medical_statement"

//医疗费用明细识别
const medicaldetailurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/medical_detail"

//病案首页识别
const medicalrecordurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/medical_record"

//保单识别
const insurancedocumentsurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/insurance_documents"

//机动车登记证书识别
const vehicleregistrationcertificateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_registration_certificate"

//二手车销售发票识别
const usedvehicleinvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/used_vehicle_invoice"

//公式识别
const formulaurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/formula"

//印章识别
const sealurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/seal"

//彩票识别
const lotteryurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/lottery"

//门脸文字识别
const facadeurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/facade"

//只能结构化识别
const intelligentocrurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/intelligent_ocr"

//自定义模板通用版
const recogniseurl = "https://aip.baidubce.com/rest/2.0/solution/v1/iocr/recognise"

//自定义模板财会版
const financeurl = "https://aip.baidubce.com/rest/2.0/solution/v1/iocr/recognise/finance"

//办公文档识别
const docanalysisofficeurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/doc_analysis_office"

func NewAipOcr(_appId string, _apiKey string, _secretKey string) (*AipOcr, error) {
	res := Oauth(_apiKey, _secretKey)
	data := map[string]string{}
	json.Unmarshal([]byte(res), &data)
	keys := GetKeys(data)
	if !In("access_token", keys) {
		return nil, errors.New(res)
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
	requrl := GetUrlBuild(basicGeneralurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别, 图片参数为远程url图片
*/
func (client *AipOcr) BasicGeneralUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(basicGeneralurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别, 图片参数为PDF
*/
func (client *AipOcr) BasicGeneralPDF(pdfFile string, options map[string]string) string {
	requrl := GetUrlBuild(basicGeneralurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（含位置信息版），参数为本地图片
*/
func (client *AipOcr) General(image string, options map[string]string) string {
	requrl := GetUrlBuild(generalurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（含位置信息版），参数为url
*/
func (client *AipOcr) GeneralUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(generalurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（含位置信息版），参数为PDF
*/
func (client *AipOcr) GeneralPDF(pdfFile string, options map[string]string) string {
	requrl := GetUrlBuild(generalurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（高精度版）参数为本地图片
*/
func (client *AipOcr) BasicAccurate(image string, options map[string]string) string {
	requrl := GetUrlBuild(basicAccurateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（高精度版）参数为url
*/
func (client *AipOcr) BasicAccurateUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(basicAccurateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（高精度版）参数为pdf_file
*/
func (client *AipOcr) BasicAccuratePDF(pdfFile string, options map[string]string) string {
	requrl := GetUrlBuild(basicAccurateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（高精度含位置版）参数为本地图片
*/
func (client *AipOcr) Accurate(image string, options map[string]string) string {
	requrl := GetUrlBuild(accurateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（高精度含位置版）参数为url
*/
func (client *AipOcr) AccurateUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(accurateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用文字识别（高精度含位置版）参数为PDF
*/
func (client *AipOcr) AccuratePDF(pdfFile string, options map[string]string) string {
	requrl := GetUrlBuild(accurateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	办公文档识别 参数为image
*/

func (client *AipOcr) DocAnalysisOffice(image string, options map[string]string) string {
	requrl := GetUrlBuild(docanalysisofficeurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	办公文档识别 参数为url
*/

func (client *AipOcr) DocAnalysisOfficeUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(docanalysisofficeurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	办公文档识别 参数为PDF
*/

func (client *AipOcr) DocAnalysisOfficePDF(pdfFile string, options map[string]string) string {
	requrl := GetUrlBuild(docanalysisofficeurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	网络图片文字识别 参数为image
*/

func (client *AipOcr) WebImage(image string, options map[string]string) string {
	requrl := GetUrlBuild(webImageurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	网络图片文字识别 参数为url
*/

func (client *AipOcr) WebImageUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(webImageurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(webimageLocurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(webimageLocurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(handwritingurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(handwritingurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(numberurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(numberurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(formurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(formurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	表格文字识别
*/

func (client *AipOcr) TableRecognition(image string, options map[string]string) string {
	requrl := GetUrlBuild(tableRecognitionAsyncurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	h := NewHttpSend(requlr)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(qrcodeurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(qrcodeurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(idcardurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["id_card_side"] = idCardSide
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(idcardurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["id_card_side"] = idCardSide
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(multiidcardurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(bankcardurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(bankcardurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(businessLicenseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(businessLicenseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(businesscardurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(businesscardurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(passporturl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(passporturl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(HKMacauexitentrypermiturl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(HKMacauexitentrypermiturl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(taiwanexitentrypermiturl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(taiwanexitentrypermiturl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(householdRegisterurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(householdRegisterurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	出生医学证明识别
*/
func (client *AipOcr) BirthCertificate(image string) string {
	requrl := GetUrlBuild(birthcertificateurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(multicardclassifyurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(multicardclassifyurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vatInvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vatInvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vatInvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	增值税发票验真
*/
func (client *AipOcr) VatInvoiceVerification(invoiceCode string, invoiceNum string, invoiceDate string, invoiceType string, checkCode string, totalAmount string) string {
	requrl := GetUrlBuild(vatinvoiceverificationurl, client.accessToken)
	options := map[string]string{
		"invoice_code": invoiceCode,
		"invoice_num":  invoiceNum,
		"invoice_date": invoiceDate,
		"invoice_type": invoiceType,
		"check_code":   checkCode,
		"total_amount": totalAmount,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(quotainvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(quotainvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(quotainvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(invoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(invoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(invoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(trainTicketurl, client.accessToken)
	if optios == nil {
		optios = map[string]string{}
	}
	optios["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(optios)
	res, err := h.Post()
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
	requrl := GetUrlBuild(trainTicketurl, client.accessToken)
	if optios == nil {
		optios = map[string]string{}
	}
	optios["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(optios)
	res, err := h.Post()
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
	requrl := GetUrlBuild(trainTicketurl, client.accessToken)
	if optios == nil {
		optios = map[string]string{}
	}
	optios["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(optios)
	res, err := h.Post()
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
	requrl := GetUrlBuild(taxiReceipturl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(taxiReceipturl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(taxiReceipturl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(airTicketurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(airTicketurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(airTicketurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(busticketurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(busticketurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(tollinvoiceurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(tollinvoiceurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(ferryticketurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(ferryticketurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(receipturl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(receipturl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(onlinetaxiitineraryurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(onlinetaxiitineraryurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(medicalinvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(medicalinvoiceurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(medicalstatementurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(medicalstatementurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(medicaldetailurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(medicaldetailurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(medicalrecordurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(medicalrecordurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(insurancedocumentsurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(insurancedocumentsurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vehicleLicenseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vehicleLicenseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(drivingLicenseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(drivingLicenseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(licensePlateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(licensePlateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vinCodeurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vinCodeurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vehicleCertificateurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vehicleCertificateurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vehicleregistrationcertificateurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vehicleregistrationcertificateurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vehicleInvoiceurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(vehicleInvoiceurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(usedvehicleinvoiceurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(usedvehicleinvoiceurl, client.accessToken)
	options := map[string]string{
		"url": url,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(docAnalysisurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(docAnalysisurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(docAnalysisurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(formulaurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(formulaurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(meterurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(meterurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(sealurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(sealurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(sealurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(lotteryurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(lotteryurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	门脸文字识别
*/
func (client *AipOcr) Facade(image string) string {
	requrl := GetUrlBuild(facadeurl, client.accessToken)
	options := map[string]string{
		"image": image,
	}
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(intelligentocrurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(intelligentocrurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(recogniseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(recogniseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(recogniseurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(financeurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(financeurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
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
	requrl := GetUrlBuild(financeurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["pdf_file"] = pdfFile
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
