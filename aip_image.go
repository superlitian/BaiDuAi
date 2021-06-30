/* Copyright 2021 CoderLee Inc. All Rights Reserved. */
/* aip_image.go - 图像模块 */
/* modification history
-----------------------
2021/6/30, by coderlee, 创建
*/
/*
DESCRIPTION
	图像模块的封装，包含：图像搜索，图像特效与增强，图像识别，车辆分析
*/

package BaiDu_Ai

import (
	"encoding/json"
	"errors"
	"strings"
)

type AipImage struct {
	appId       string
	apiKey      string
	secretKey   string
	accessToken map[string]string
}

func NewAipImage(_appId string, _apiKey string, _secretKey string) (*AipImage, error) {
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
	return &AipImage{
		appId:       _appId,
		apiKey:      _apiKey,
		secretKey:   _secretKey,
		accessToken: t,
	}, nil
}

//组合接口api
const combinationurl = "https://aip.baidubce.com/api/v1/solution/direct/imagerecognition/combination"

//通用物体和场景识别高级版
const advancedgeneralurl = "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general"

//图像单主体检测
const objectdetecturl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/object_detect"

//动物识别
const animalurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/animal"

//植物识别
const planturl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/plant"

//logo识别
const logourl = "https://aip.baidubce.com/rest/2.0/image-classify/v2/logo"

//果蔬识别
const ingredienturl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/classify/ingredient"

//自定义菜品识别-入库
const dishaddurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/dish/add"

//自定义菜品-检索
const dishsearchurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/dish/search"

//自定义菜品-删除
const dishdeleteurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/dish/delete"

//菜品识别
const dishurl = "https://aip.baidubce.com/rest/2.0/image-classify/v2/dish"

//红酒识别
const redwineurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/redwine"

//货币识别
const currencyurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/currency"

//地标识别
const landmarkurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/landmark"

//图像多主体检测
const multiobjectdetecturl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/multi_object_detect"

//自定义红酒-入库
const redwineaddurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/redwine/add"

//自定义红酒-检索
const redwinesearchurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/redwine/search"

//自定义红酒-删除
const redwinedeleteurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/redwine/delete"

//自定义红酒—更新
const redwineupdateurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/redwine/update"

//黑白图像上色
const colourizeurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/colourize"

//图像风格转换
const styletransurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/style_trans"

//人像动漫化
const selfieanimeurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/selfie_anime"

//天空分割
const skysegurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/sky_seg"

//图像去雾
const dehazeurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/dehaze"

//图像对比度增强
const contrastenhanceurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/contrast_enhance"

//图像无损放大
const imagequalityenhanceurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/image_quality_enhance"

//拉伸图像恢复
const stretchrestoreurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/stretch_restore"

//图像修复
const inpaintingurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/inpainting"

//图像清晰度增强
const imagedefinitionenhanceurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/image_definition_enhance"

//图像色彩增强
const colorenhanceurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/color_enhance"

//相似图片搜索—入库
const similaraddurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/add"

//相似图片搜索—检索
const similarsearchurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/search"

//相似图片搜索—删除
const similardeleteurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/delete"

//相似图片搜索—更新
const similarupdateurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/update"

//相同图片搜索—入库
const samehqaddurl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/add"

//相同图片搜索—检索
const samehqsearchurl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/search"

//相同图片搜索—删除
const samehqdeleteurl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/delete"

//相同图片搜索—更新
const samehqupdateurl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/update"

//商品图片搜索—入库
const productaddurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/add"

//商品图片搜索—检索
const productsearch = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/search"

//商品图片搜索—删除
const productdeleteurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/delete"

//商品图片搜索—更新
const productupdateurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/update"

//绘本图片搜索—入库
const picturebookaddurl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/add"

//绘本图片搜索—检索
const picturebooksearchurl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/search"

//绘本图片搜索—删除
const picturebookdeleteurl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/delete"

//绘本图片搜索—更新
const picturebookupdateurl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/update"

//车型识别
const carurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/car"

//车辆检测
const vehicledetecturl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/vehicle_detect"

//车辆外观损伤识别
const vehicledamageurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/vehicle_damage"

//车流统计
const trafficflowurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/traffic_flow"

//车辆属性识别
const vehicleattrurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/vehicle_attr"

//车辆分割
const vehiclesegurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/vehicle_seg"

//车辆检测-高空版
const vehicledetecthighurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/vehicle_detect_high"

/*
	组合接口API
	参数为image
*/
func (client *AipImage) Combination(image string, scenes []string, options map[string]interface{}) string {
	requrl := GetUrlBuild(combinationurl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["image"] = image
	options["scenes"] = scenes
	h := NewHttpSend(requrl)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	组合接口API
	参数为imageurl
*/
func (client *AipImage) CombinationUrl(imgUrl string, scenes []string, options map[string]interface{}) string {
	requrl := GetUrlBuild(combinationurl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["imgUrl"] = imgUrl
	options["scenes"] = scenes
	h := NewHttpSend(requrl)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	通用物体和场景识别高级版
	参数为本地图片
*/
func (client *AipImage) AdvancedGeneral(image string, options map[string]string) string {
	requrl := GetUrlBuild(advancedgeneralurl, client.accessToken)
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
	通用物体和场景识别高级版
	参数为url
*/
func (client *AipImage) AdvancedGeneralUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(advancedgeneralurl, client.accessToken)
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
	图像单主体检测
*/
func (client *AipImage) ObjectDetect(image string, options map[string]string) string {
	requrl := GetUrlBuild(objectdetecturl, client.accessToken)
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
	动物识别
	参数为本地图片
*/
func (client *AipImage) Animal(image string, options map[string]string) string {
	requrl := GetUrlBuild(animalurl, client.accessToken)
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
	动物识别
	参数为url
*/
func (client *AipImage) AnimalUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(animalurl, client.accessToken)
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
	植物识别
	参数为本地图片
*/
func (client *AipImage) Plant(image string, options map[string]string) string {
	requrl := GetUrlBuild(planturl, client.accessToken)
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
	植物识别
	参数为url
*/
func (client *AipImage) PlantUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(planturl, client.accessToken)
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
	logo识别
	参数为本地图片
*/
func (client *AipImage) Logo(image string, options map[string]string) string {
	requrl := GetUrlBuild(logourl, client.accessToken)
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
	logo识别
	参数为url
*/
func (client *AipImage) LogoUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(logourl, client.accessToken)
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
	果蔬识别
	参数为本地图片
*/
func (client *AipImage) Ingredient(image string, options map[string]string) string {
	requrl := GetUrlBuild(ingredienturl, client.accessToken)
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
	果蔬识别
	参数为url
*/
func (client *AipImage) IngredientUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(ingredienturl, client.accessToken)
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
	自定义菜品-入库
	参数为本地图片
*/
func (client *AipImage) DishAdd(image string, options map[string]string) string {
	requrl := GetUrlBuild(dishaddurl, client.accessToken)
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
	自定义菜品-入库
	参数为url
*/
func (client *AipImage) DishAddUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(dishaddurl, client.accessToken)
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
	自定义菜品-检索
	参数为本地图片
*/
func (client *AipImage) DishSearch(image string) string {
	requrl := GetUrlBuild(dishsearchurl, client.accessToken)
	options := map[string]string{}
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
	自定义菜品-检索
	参数为url
*/
func (client *AipImage) DishSearchUrl(url string) string {
	requrl := GetUrlBuild(dishsearchurl, client.accessToken)
	options := map[string]string{}
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
	自定义菜品-删除
	参数为image
*/
func (client *AipImage) DishDelete(image string, options map[string]string) string {
	requrl := GetUrlBuild(dishdeleteurl, client.accessToken)
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
	自定义菜品-删除
	参数为url
*/
func (client *AipImage) DishDeleteUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(dishdeleteurl, client.accessToken)
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
	菜品识别
	参数为本地图片
*/
func (client *AipImage) Dish(image string, options map[string]string) string {
	requrl := GetUrlBuild(dishurl, client.accessToken)
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
	菜品识别
	参数为url
*/
func (client *AipImage) DishUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(dishurl, client.accessToken)
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
	红酒识别
	参数为本地图片
*/
func (client *AipImage) Redwine(image string) string {
	requrl := GetUrlBuild(redwineurl, client.accessToken)
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
	红酒识别
	参数为url
*/
func (client *AipImage) RedwineUrl(url string) string {
	requrl := GetUrlBuild(redwineurl, client.accessToken)
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
	货币识别
	参数为本地图片
*/
func (client *AipImage) Currency(image string) string {
	requrl := GetUrlBuild(currencyurl, client.accessToken)
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
	货币识别
	参数为url
*/
func (client *AipImage) CurrencyUrl(url string) string {
	requrl := GetUrlBuild(currencyurl, client.accessToken)
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
	地标识别
	参数为本地图片
*/
func (client *AipImage) Landmark(image string) string {
	requrl := GetUrlBuild(landmarkurl, client.accessToken)
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
	地标识别
	参数为url
*/
func (client *AipImage) LandmarkUrl(url string) string {
	requrl := GetUrlBuild(landmarkurl, client.accessToken)
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
	图像多主体检测
	参数为本地图片
*/
func (client *AipImage) MultiObjectDetect(image string) string {
	requrl := GetUrlBuild(multiobjectdetecturl, client.accessToken)
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
	图像多主体检测
	参数为url
*/
func (client *AipImage) MultiObjectDetectUrl(url string) string {
	requrl := GetUrlBuild(multiobjectdetecturl, client.accessToken)
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
	自定义红酒-入库
	参数为本地图片
*/
func (client *AipImage) RedwineAdd(image string, options map[string]string) string {
	requrl := GetUrlBuild(redwineaddurl, client.accessToken)
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
	自定义红酒-入库
	参数为url
*/
func (client *AipImage) RedwineAddUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(redwineaddurl, client.accessToken)
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
	自定义红酒-检索
	参数为本地图片
*/
func (client *AipImage) RedwineSearch(image string, options map[string]string) string {
	requrl := GetUrlBuild(redwinesearchurl, client.accessToken)
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
	自定义红酒-检索
	参数为url
*/
func (client *AipImage) RedwineSearchUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(redwinesearchurl, client.accessToken)
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
	自定义红酒-删除
	参数为本地图片
*/
func (client *AipImage) RedwineDelete(image string) string {
	requrl := GetUrlBuild(redwinedeleteurl, client.accessToken)
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
	自定义红酒-删除
	参数为cont_sign_list
*/
func (client *AipImage) RedwineDeleteContSignList(contSignList string) string {
	requrl := GetUrlBuild(redwinedeleteurl, client.accessToken)
	options := map[string]string{
		"cont_sign_list": contSignList,
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
	自定义红酒—更新
	参数为本地图片
*/
func (client *AipImage) RedwineUpdate(image string, options map[string]string) string {
	requrl := GetUrlBuild(redwineupdateurl, client.accessToken)
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
	自定义红酒—更新
	参数为url
*/
func (client *AipImage) RedwineUpdateUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(redwineupdateurl, client.accessToken)
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
	黑白图像上色
	参数为本地图片
*/
func (client *AipImage) Colourize(image string) string {
	requrl := GetUrlBuild(colourizeurl, client.accessToken)
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
	黑白图像上色
	参数为url
*/
func (client *AipImage) ColourizeUrl(url string) string {
	requrl := GetUrlBuild(colourizeurl, client.accessToken)
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
	图像风格转换
	参数为本地图片
*/
func (client *AipImage) StyleTrans(image string, options map[string]string) string {
	requrl := GetUrlBuild(styletransurl, client.accessToken)
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
	图像风格转换
	参数为url
*/
func (client *AipImage) StyleTransUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(styletransurl, client.accessToken)
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
	人像动漫化
	参数为本地图片
*/
func (client *AipImage) SelfieAnime(image string, options map[string]string) string {
	requrl := GetUrlBuild(selfieanimeurl, client.accessToken)
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
	人像动漫化
	参数为url
*/
func (client *AipImage) SelfieAnimeUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(selfieanimeurl, client.accessToken)
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
	天空分割
	参数为本地图片
*/
func (client *AipImage) SkySeg(image string) string {
	requrl := GetUrlBuild(skysegurl, client.accessToken)
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
	天空分割
	参数为url
*/
func (client *AipImage) SkySegUrl(url string) string {
	requrl := GetUrlBuild(skysegurl, client.accessToken)
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
	图像去雾
	参数为本地图片
*/
func (client *AipImage) Dehaze(image string) string {
	requrl := GetUrlBuild(dehazeurl, client.accessToken)
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
	图像去雾
	参数为url
*/
func (client *AipImage) DehazeUrl(url string) string {
	requrl := GetUrlBuild(dehazeurl, client.accessToken)
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
	图像对比度增强
	参数为本地图片
*/
func (client *AipImage) ContrastEnhance(image string) string {
	requrl := GetUrlBuild(contrastenhanceurl, client.accessToken)
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
	图像对比度增强
	参数为url
*/
func (client *AipImage) ContrastEnhanceUrl(url string) string {
	requrl := GetUrlBuild(contrastenhanceurl, client.accessToken)
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
	图像无损放大
	参数为本地图片
*/
func (client *AipImage) ImageQualityEnhance(image string) string {
	requrl := GetUrlBuild(imagequalityenhanceurl, client.accessToken)
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
	图像无损放大
	参数为url
*/
func (client *AipImage) ImageQualityEnhanceUrl(url string) string {
	requrl := GetUrlBuild(imagequalityenhanceurl, client.accessToken)
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
	拉伸图像恢复
	参数为本地图片
*/
func (client *AipImage) StretchRestore(image string) string {
	requrl := GetUrlBuild(stretchrestoreurl, client.accessToken)
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
	拉伸图像恢复
	参数为url
*/
func (client *AipImage) StretchRestoreUrl(url string) string {
	requrl := GetUrlBuild(stretchrestoreurl, client.accessToken)
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
	图像修复
	参数为本地图片
*/
func (client *AipImage) Inpainting(rectangle []map[string]string, image string) string {
	requrl := GetUrlBuild(inpaintingurl, client.accessToken)
	options := map[string]interface{}{
		"rectangle": rectangle,
		"image":     image,
	}
	h := NewHttpSend(requrl)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	图像修复
	参数为url
*/
func (client *AipImage) InpaintingUrl(rectangle []map[string]string, url string) string {
	requrl := GetUrlBuild(inpaintingurl, client.accessToken)
	options := map[string]interface{}{
		"rectangle": rectangle,
		"url":       url,
	}
	h := NewHttpSend(requrl)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	图像清晰度增强
	参数为本地图片
*/
func (client *AipImage) ImageDefinitionEnhance(image string) string {
	requrl := GetUrlBuild(imagedefinitionenhanceurl, client.accessToken)
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
	图像清晰度增强
	参数为url
*/
func (client *AipImage) ImageDefinitionEnhanceUrl(url string) string {
	requrl := GetUrlBuild(imagedefinitionenhanceurl, client.accessToken)
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
	图像色彩增强
	参数为本地图片
*/
func (client *AipImage) ColorEnhance(image string) string {
	requrl := GetUrlBuild(colorenhanceurl, client.accessToken)
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
	图像色彩增强
	参数为url
*/
func (client *AipImage) ColorEnhanceUrl(url string) string {
	requrl := GetUrlBuild(colorenhanceurl, client.accessToken)
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
	相似图片搜索—入库
	参数为本地图片
*/
func (client *AipImage) SimilarAdd(image string, brief string, options map[string]string) string {
	requrl := GetUrlBuild(similaraddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["brief"] = brief
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	相似图片搜索—入库
	参数为url
*/
func (client *AipImage) SimilarAddUrl(url string, brief string, options map[string]string) string {
	requrl := GetUrlBuild(similaraddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["brief"] = brief
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	相似图片搜索—检索
	参数为本地图片
*/
func (client *AipImage) SimilarSearch(image string, options map[string]string) string {
	requrl := GetUrlBuild(similarsearchurl, client.accessToken)
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
	相似图片搜索—检索
	参数为url
*/
func (client *AipImage) SimilarSearchUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(similarsearchurl, client.accessToken)
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
	相似图片搜索—删除
	参数为本地图片
*/
func (client *AipImage) SimilarDelete(image string) string {
	requrl := GetUrlBuild(similardeleteurl, client.accessToken)
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
	相似图片搜索—删除
	参数为url
*/
func (client *AipImage) SimilarDeleteUrl(url string) string {
	requrl := GetUrlBuild(similardeleteurl, client.accessToken)
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
	相似图片搜索—删除
	cont_sign
*/
func (client *AipImage) SimilarDeleteContSign(contSign string) string {
	requrl := GetUrlBuild(similardeleteurl, client.accessToken)
	options := map[string]string{
		"cont_sign": contSign,
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
	相似图片搜索—更新
	参数为本地图片
*/
func (client *AipImage) SimilarUpdate(image string) string {
	requrl := GetUrlBuild(similarupdateurl, client.accessToken)
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
	相似图片搜索—更新
	参数为url
*/
func (client *AipImage) SimilarUpdateUrl(url string) string {
	requrl := GetUrlBuild(similarupdateurl, client.accessToken)
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
	相似图片搜索—更新
	cont_sign
*/
func (client *AipImage) SimilarUpdateContSign(contSign string) string {
	requrl := GetUrlBuild(similarupdateurl, client.accessToken)
	options := map[string]string{
		"cont_sign": contSign,
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
	相同图片搜索—入库
	参数为本地图片
*/
func (client *AipImage) SameHqAdd(image string, brief string, options map[string]string) string {
	requrl := GetUrlBuild(samehqaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["brief"] = brief
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	相同图片搜索—入库
	参数为url
*/
func (client *AipImage) SameHqAddUrl(url string, brief string, options map[string]string) string {
	requrl := GetUrlBuild(samehqaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["brief"] = brief
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	相同图片搜索—检索
	参数为本地图片
*/
func (client *AipImage) SameHqSearch(image string, options map[string]string) string {
	requrl := GetUrlBuild(samehqsearchurl, client.accessToken)
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
	相同图片搜索—检索
	参数为url
*/
func (client *AipImage) SameHqSearchUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(samehqsearchurl, client.accessToken)
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
	相同图片搜索—删除
	参数为本地图片
*/
func (client *AipImage) SameHqDelete(image string, options map[string]string) string {
	requrl := GetUrlBuild(samehqdeleteurl, client.accessToken)
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
	相同图片搜索—删除
	参数为url
*/
func (client *AipImage) SameHqDeleteUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(samehqdeleteurl, client.accessToken)
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
	相同图片搜索—删除
	参数为cont_sign
*/
func (client *AipImage) SameHqDeleteContSign(contSign string, options map[string]string) string {
	requrl := GetUrlBuild(samehqdeleteurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["cont_sign"] = contSign
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	相同图片搜索—更新
	参数为本地图片
*/
func (client *AipImage) SameHqUpdate(image string, options map[string]string) string {
	requrl := GetUrlBuild(samehqupdateurl, client.accessToken)
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
	相同图片搜索—更新
	参数为url
*/
func (client *AipImage) SameHqUpdateUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(samehqupdateurl, client.accessToken)
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
	相同图片搜索—更新
	参数为cont_sign
*/
func (client *AipImage) SameHqUpdateContSign(contSign string, options map[string]string) string {
	requrl := GetUrlBuild(samehqupdateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["cont_sign"] = contSign
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	商品图片搜索—入库
	参数为本地图片
*/
func (client *AipImage) ProductAdd(image string, brief string, options map[string]string) string {
	requrl := GetUrlBuild(productaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["brief"] = brief
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	商品图片搜索—入库
	参数为url
*/
func (client *AipImage) ProductAddUrl(url string, brief string, options map[string]string) string {
	requrl := GetUrlBuild(productaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["brief"] = brief
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	商品图片搜索—检索
	参数为本地图片
*/
func (client *AipImage) ProductSearch(image string, options map[string]string) string {
	requrl := GetUrlBuild(productsearch, client.accessToken)
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
	商品图片搜索—检索
	参数为url
*/
func (client *AipImage) ProductSearchUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(productsearch, client.accessToken)
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
	商品图片搜索—删除
	参数为本地图片
*/
func (client *AipImage) ProductDelete(image string) string {
	requrl := GetUrlBuild(productdeleteurl, client.accessToken)
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
	商品图片搜索—删除
	参数为url
*/
func (client *AipImage) ProductDeleteUrl(url string) string {
	requrl := GetUrlBuild(productdeleteurl, client.accessToken)
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
	商品图片搜索—删除
	参数为cont_sign
*/
func (client *AipImage) ProductDeleteContSign(contSign string) string {
	requrl := GetUrlBuild(productdeleteurl, client.accessToken)
	options := map[string]string{
		"cont_sign": contSign,
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
	商品图片搜索—更新
	参数为本地图片
*/
func (client *AipImage) ProductUpdate(image string, options map[string]string) string {
	requrl := GetUrlBuild(productupdateurl, client.accessToken)
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
	商品图片搜索—更新
	参数为url
*/
func (client *AipImage) ProductUpdateUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(productupdateurl, client.accessToken)
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
	商品图片搜索—更新
	参数为cont_sign
*/
func (client *AipImage) ProductUpdateContSign(cont_sign string, options map[string]string) string {
	requrl := GetUrlBuild(productupdateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["cont_sign"] = cont_sign
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	绘本图片搜索—入库
	参数为本地图片
*/
func (client *AipImage) PicturebookAdd(image string, brief string, options map[string]string) string {
	requrl := GetUrlBuild(picturebookaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["brief"] = brief
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	绘本图片搜索—入库
	参数为url
*/
func (client *AipImage) PicturebookAddUrl(url string, brief string, options map[string]string) string {
	requrl := GetUrlBuild(picturebookaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["brief"] = brief
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	绘本图片搜索—检索
	参数为本地图片
*/
func (client *AipImage) PicturebookSearch(image string, options map[string]string) string {
	requrl := GetUrlBuild(picturebooksearchurl, client.accessToken)
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
	绘本图片搜索—检索
	参数为url
*/
func (client *AipImage) PicturebookSearchUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(picturebooksearchurl, client.accessToken)
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
	绘本图片搜索—删除
	参数为本地图片
*/
func (client *AipImage) PicturebookDelete(image string) string {
	requrl := GetUrlBuild(picturebookdeleteurl, client.accessToken)
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
	绘本图片搜索—删除
	参数为url
*/
func (client *AipImage) PicturebookDeleteUrl(url string) string {
	requrl := GetUrlBuild(picturebookdeleteurl, client.accessToken)
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
	绘本图片搜索—删除
	参数为cont_sign
*/
func (client *AipImage) PicturebookDeleteContSign(contSign string) string {
	requrl := GetUrlBuild(picturebookdeleteurl, client.accessToken)
	options := map[string]string{
		"cont_sign": contSign,
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
	绘本图片搜索—更新
*/
func (client *AipImage) PicturebookUpdate(image string, options map[string]string) string {
	requrl := GetUrlBuild(picturebookupdateurl, client.accessToken)
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
	绘本图片搜索—更新
	参数为url
*/
func (client *AipImage) PicturebookUpdateUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(picturebookupdateurl, client.accessToken)
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
	绘本图片搜索—更新
	参数为cont_sign
*/
func (client *AipImage) PicturebookUpdateUrlContSign(contSign string, options map[string]string) string {
	requrl := GetUrlBuild(picturebookupdateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["cont_sign"] = contSign
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	车型识别
	参数为本地图片
*/
func (client *AipImage) Car(image string, options map[string]string) string {
	requrl := GetUrlBuild(carurl, client.accessToken)
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
	车型识别
	参数为url
*/
func (client *AipImage) CarUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(carurl, client.accessToken)
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
	车辆检测
	参数为本地图片
*/
func (client *AipImage) VehicleDetect(image string, options map[string]string) string {
	requrl := GetUrlBuild(vehicledetecturl, client.accessToken)
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
	车辆检测
	参数为url
*/
func (client *AipImage) VehicleDetectUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(vehicledetecturl, client.accessToken)
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
	车辆外观损伤识别
	参数为本地图片
*/
func (client *AipImage) VehicleDamage(image string) string {
	requrl := GetUrlBuild(vehicledamageurl, client.accessToken)
	options := map[string]string{}
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
	车辆外观损伤识别
	参数为url
*/
func (client *AipImage) VehicleDamageUrl(url string) string {
	requrl := GetUrlBuild(vehicledamageurl, client.accessToken)
	options := map[string]string{}
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
	车流统计
	参数为本地图片
*/
func (client *AipImage) TrafficFlow(image string, caseId int, caseInit string, area string, options map[string]interface{}) string {
	requrl := GetUrlBuild(trafficflowurl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["image"] = image
	options["case_id"] = caseId
	options["case_init"] = caseInit
	options["area"] = area
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	车流统计
	参数为url
*/
func (client *AipImage) TrafficFlowUrl(url string, caseId int, caseInit string, area string, options map[string]interface{}) string {
	requrl := GetUrlBuild(trafficflowurl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["url"] = url
	options["case_id"] = caseId
	options["case_init"] = caseInit
	options["area"] = area
	h := NewHttpSend(requrl)
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

/*
	车辆属性识别
	参数为本地图片
*/
func (client *AipImage) VehicleAttr(image string, options map[string]string) string {
	requrl := GetUrlBuild(vehicleattrurl, client.accessToken)
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
	车辆属性识别
	参数为url
*/
func (client *AipImage) VehicleAttrUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(vehicleattrurl, client.accessToken)
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
	车辆分割
	参数为本地图片
*/
func (client *AipImage) VehicleSeg(image string, options map[string]string) string {
	requrl := GetUrlBuild(vehiclesegurl, client.accessToken)
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
	车辆分割
	参数为url
*/
func (client *AipImage) VehicleSegUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(vehiclesegurl, client.accessToken)
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
	车辆检测-高空版
	参数为本地图片
*/
func (client *AipImage) VehicleDetectHigh(image string, options map[string]string) string {
	requrl := GetUrlBuild(vehicledetecthighurl, client.accessToken)
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
	车辆检测-高空版
	参数为url
*/
func (client *AipImage) VehicleDetectHighUrl(url string, options map[string]string) string {
	requrl := GetUrlBuild(vehicledetecthighurl, client.accessToken)
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
