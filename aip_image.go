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
	return &AipImage{
		appId:       _appId,
		apiKey:      _apiKey,
		secretKey:   _secretKey,
		accessToken: t,
	}, nil
}

/*
	组合接口API
	参数为image
*/
func (client *AipImage) Combination(image string, scenes []string, options map[string]interface{}) string {
	reqUrl = GetUrlBuild(combinationurl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["image"] = image
	options["scenes"] = scenes
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
	组合接口API
	参数为imageurl
*/
func (client *AipImage) CombinationUrl(imgUrl string, scenes []string, options map[string]interface{}) string {
	reqUrl = GetUrlBuild(combinationurl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["imgUrl"] = imgUrl
	options["scenes"] = scenes
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
	通用物体和场景识别高级版
	参数为本地图片
*/
func (client *AipImage) AdvancedGeneral(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(advancedgeneralurl, client.accessToken)
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
	通用物体和场景识别高级版
	参数为url
*/
func (client *AipImage) AdvancedGeneralUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(advancedgeneralurl, client.accessToken)
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
	图像单主体检测
*/
func (client *AipImage) ObjectDetect(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(objectdetecturl, client.accessToken)
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
	动物识别
	参数为本地图片
*/
func (client *AipImage) Animal(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(animalurl, client.accessToken)
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
	动物识别
	参数为url
*/
func (client *AipImage) AnimalUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(animalurl, client.accessToken)
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
	植物识别
	参数为本地图片
*/
func (client *AipImage) Plant(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(planturl, client.accessToken)
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
	植物识别
	参数为url
*/
func (client *AipImage) PlantUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(planturl, client.accessToken)
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
	logo识别
	参数为本地图片
*/
func (client *AipImage) Logo(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(logourl, client.accessToken)
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
	logo识别
	参数为url
*/
func (client *AipImage) LogoUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(logourl, client.accessToken)
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
	果蔬识别
	参数为本地图片
*/
func (client *AipImage) Ingredient(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(ingredienturl, client.accessToken)
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
	果蔬识别
	参数为url
*/
func (client *AipImage) IngredientUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(ingredienturl, client.accessToken)
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
	自定义菜品-入库
	参数为本地图片
*/
func (client *AipImage) DishAdd(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(dishaddurl, client.accessToken)
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
	自定义菜品-入库
	参数为url
*/
func (client *AipImage) DishAddUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(dishaddurl, client.accessToken)
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
	自定义菜品-检索
	参数为本地图片
*/
func (client *AipImage) DishSearch(image string) string {
	reqUrl = GetUrlBuild(dishsearchurl, client.accessToken)
	options := map[string]string{}
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
	自定义菜品-检索
	参数为url
*/
func (client *AipImage) DishSearchUrl(url string) string {
	reqUrl = GetUrlBuild(dishsearchurl, client.accessToken)
	options := map[string]string{}
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
	自定义菜品-删除
	参数为image
*/
func (client *AipImage) DishDelete(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(dishdeleteurl, client.accessToken)
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
	自定义菜品-删除
	参数为url
*/
func (client *AipImage) DishDeleteUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(dishdeleteurl, client.accessToken)
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
	菜品识别
	参数为本地图片
*/
func (client *AipImage) Dish(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(dishurl, client.accessToken)
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
	菜品识别
	参数为url
*/
func (client *AipImage) DishUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(dishurl, client.accessToken)
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
	红酒识别
	参数为本地图片
*/
func (client *AipImage) Redwine(image string) string {
	reqUrl = GetUrlBuild(redwineurl, client.accessToken)
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
	红酒识别
	参数为url
*/
func (client *AipImage) RedwineUrl(url string) string {
	reqUrl = GetUrlBuild(redwineurl, client.accessToken)
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
	货币识别
	参数为本地图片
*/
func (client *AipImage) Currency(image string) string {
	reqUrl = GetUrlBuild(currencyurl, client.accessToken)
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
	货币识别
	参数为url
*/
func (client *AipImage) CurrencyUrl(url string) string {
	reqUrl = GetUrlBuild(currencyurl, client.accessToken)
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
	地标识别
	参数为本地图片
*/
func (client *AipImage) Landmark(image string) string {
	reqUrl = GetUrlBuild(landmarkurl, client.accessToken)
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
	地标识别
	参数为url
*/
func (client *AipImage) LandmarkUrl(url string) string {
	reqUrl = GetUrlBuild(landmarkurl, client.accessToken)
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
	图像多主体检测
	参数为本地图片
*/
func (client *AipImage) MultiObjectDetect(image string) string {
	reqUrl = GetUrlBuild(multiobjectdetecturl, client.accessToken)
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
	图像多主体检测
	参数为url
*/
func (client *AipImage) MultiObjectDetectUrl(url string) string {
	reqUrl = GetUrlBuild(multiobjectdetecturl, client.accessToken)
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
	自定义红酒-入库
	参数为本地图片
*/
func (client *AipImage) RedwineAdd(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(redwineaddurl, client.accessToken)
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
	自定义红酒-入库
	参数为url
*/
func (client *AipImage) RedwineAddUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(redwineaddurl, client.accessToken)
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
	自定义红酒-检索
	参数为本地图片
*/
func (client *AipImage) RedwineSearch(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(redwinesearchurl, client.accessToken)
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
	自定义红酒-检索
	参数为url
*/
func (client *AipImage) RedwineSearchUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(redwinesearchurl, client.accessToken)
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
	自定义红酒-删除
	参数为本地图片
*/
func (client *AipImage) RedwineDelete(image string) string {
	reqUrl = GetUrlBuild(redwinedeleteurl, client.accessToken)
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
	自定义红酒-删除
	参数为cont_sign_list
*/
func (client *AipImage) RedwineDeleteContSignList(contSignList string) string {
	reqUrl = GetUrlBuild(redwinedeleteurl, client.accessToken)
	options := map[string]string{
		"cont_sign_list": contSignList,
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
	自定义红酒—更新
	参数为本地图片
*/
func (client *AipImage) RedwineUpdate(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(redwineupdateurl, client.accessToken)
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
	自定义红酒—更新
	参数为url
*/
func (client *AipImage) RedwineUpdateUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(redwineupdateurl, client.accessToken)
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
	黑白图像上色
	参数为本地图片
*/
func (client *AipImage) Colourize(image string) string {
	reqUrl = GetUrlBuild(colourizeurl, client.accessToken)
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
	黑白图像上色
	参数为url
*/
func (client *AipImage) ColourizeUrl(url string) string {
	reqUrl = GetUrlBuild(colourizeurl, client.accessToken)
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
	图像风格转换
	参数为本地图片
*/
func (client *AipImage) StyleTrans(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(styletransurl, client.accessToken)
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
	图像风格转换
	参数为url
*/
func (client *AipImage) StyleTransUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(styletransurl, client.accessToken)
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
	人像动漫化
	参数为本地图片
*/
func (client *AipImage) SelfieAnime(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(selfieanimeurl, client.accessToken)
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
	人像动漫化
	参数为url
*/
func (client *AipImage) SelfieAnimeUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(selfieanimeurl, client.accessToken)
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
	天空分割
	参数为本地图片
*/
func (client *AipImage) SkySeg(image string) string {
	reqUrl = GetUrlBuild(skysegurl, client.accessToken)
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
	天空分割
	参数为url
*/
func (client *AipImage) SkySegUrl(url string) string {
	reqUrl = GetUrlBuild(skysegurl, client.accessToken)
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
	图像去雾
	参数为本地图片
*/
func (client *AipImage) Dehaze(image string) string {
	reqUrl = GetUrlBuild(dehazeurl, client.accessToken)
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
	图像去雾
	参数为url
*/
func (client *AipImage) DehazeUrl(url string) string {
	reqUrl = GetUrlBuild(dehazeurl, client.accessToken)
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
	图像对比度增强
	参数为本地图片
*/
func (client *AipImage) ContrastEnhance(image string) string {
	reqUrl = GetUrlBuild(contrastenhanceurl, client.accessToken)
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
	图像对比度增强
	参数为url
*/
func (client *AipImage) ContrastEnhanceUrl(url string) string {
	reqUrl = GetUrlBuild(contrastenhanceurl, client.accessToken)
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
	图像无损放大
	参数为本地图片
*/
func (client *AipImage) ImageQualityEnhance(image string) string {
	reqUrl = GetUrlBuild(imagequalityenhanceurl, client.accessToken)
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
	图像无损放大
	参数为url
*/
func (client *AipImage) ImageQualityEnhanceUrl(url string) string {
	reqUrl = GetUrlBuild(imagequalityenhanceurl, client.accessToken)
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
	拉伸图像恢复
	参数为本地图片
*/
func (client *AipImage) StretchRestore(image string) string {
	reqUrl = GetUrlBuild(stretchrestoreurl, client.accessToken)
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
	拉伸图像恢复
	参数为url
*/
func (client *AipImage) StretchRestoreUrl(url string) string {
	reqUrl = GetUrlBuild(stretchrestoreurl, client.accessToken)
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
	图像修复
	参数为本地图片
*/
func (client *AipImage) Inpainting(rectangle []map[string]string, image string) string {
	reqUrl = GetUrlBuild(inpaintingurl, client.accessToken)
	options := map[string]interface{}{
		"rectangle": rectangle,
		"image":     image,
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
	图像修复
	参数为url
*/
func (client *AipImage) InpaintingUrl(rectangle []map[string]string, url string) string {
	reqUrl = GetUrlBuild(inpaintingurl, client.accessToken)
	options := map[string]interface{}{
		"rectangle": rectangle,
		"url":       url,
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
	图像清晰度增强
	参数为本地图片
*/
func (client *AipImage) ImageDefinitionEnhance(image string) string {
	reqUrl = GetUrlBuild(imagedefinitionenhanceurl, client.accessToken)
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
	图像清晰度增强
	参数为url
*/
func (client *AipImage) ImageDefinitionEnhanceUrl(url string) string {
	reqUrl = GetUrlBuild(imagedefinitionenhanceurl, client.accessToken)
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
	图像色彩增强
	参数为本地图片
*/
func (client *AipImage) ColorEnhance(image string) string {
	reqUrl = GetUrlBuild(colorenhanceurl, client.accessToken)
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
	图像色彩增强
	参数为url
*/
func (client *AipImage) ColorEnhanceUrl(url string) string {
	reqUrl = GetUrlBuild(colorenhanceurl, client.accessToken)
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
	相似图片搜索—入库
	参数为本地图片
*/
func (client *AipImage) SimilarAdd(image string, brief string, options map[string]string) string {
	reqUrl = GetUrlBuild(similaraddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["brief"] = brief
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(similaraddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["brief"] = brief
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(similarsearchurl, client.accessToken)
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
	相似图片搜索—检索
	参数为url
*/
func (client *AipImage) SimilarSearchUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(similarsearchurl, client.accessToken)
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
	相似图片搜索—删除
	参数为本地图片
*/
func (client *AipImage) SimilarDelete(image string) string {
	reqUrl = GetUrlBuild(similardeleteurl, client.accessToken)
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
	相似图片搜索—删除
	参数为url
*/
func (client *AipImage) SimilarDeleteUrl(url string) string {
	reqUrl = GetUrlBuild(similardeleteurl, client.accessToken)
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
	相似图片搜索—删除
	cont_sign
*/
func (client *AipImage) SimilarDeleteContSign(contSign string) string {
	reqUrl = GetUrlBuild(similardeleteurl, client.accessToken)
	options := map[string]string{
		"cont_sign": contSign,
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
	相似图片搜索—更新
	参数为本地图片
*/
func (client *AipImage) SimilarUpdate(image string) string {
	reqUrl = GetUrlBuild(similarupdateurl, client.accessToken)
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
	相似图片搜索—更新
	参数为url
*/
func (client *AipImage) SimilarUpdateUrl(url string) string {
	reqUrl = GetUrlBuild(similarupdateurl, client.accessToken)
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
	相似图片搜索—更新
	cont_sign
*/
func (client *AipImage) SimilarUpdateContSign(contSign string) string {
	reqUrl = GetUrlBuild(similarupdateurl, client.accessToken)
	options := map[string]string{
		"cont_sign": contSign,
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
	相同图片搜索—入库
	参数为本地图片
*/
func (client *AipImage) SameHqAdd(image string, brief string, options map[string]string) string {
	reqUrl = GetUrlBuild(samehqaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["brief"] = brief
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(samehqaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["brief"] = brief
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(samehqsearchurl, client.accessToken)
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
	相同图片搜索—检索
	参数为url
*/
func (client *AipImage) SameHqSearchUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(samehqsearchurl, client.accessToken)
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
	相同图片搜索—删除
	参数为本地图片
*/
func (client *AipImage) SameHqDelete(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(samehqdeleteurl, client.accessToken)
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
	相同图片搜索—删除
	参数为url
*/
func (client *AipImage) SameHqDeleteUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(samehqdeleteurl, client.accessToken)
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
	相同图片搜索—删除
	参数为cont_sign
*/
func (client *AipImage) SameHqDeleteContSign(contSign string, options map[string]string) string {
	reqUrl = GetUrlBuild(samehqdeleteurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["cont_sign"] = contSign
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(samehqupdateurl, client.accessToken)
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
	相同图片搜索—更新
	参数为url
*/
func (client *AipImage) SameHqUpdateUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(samehqupdateurl, client.accessToken)
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
	相同图片搜索—更新
	参数为cont_sign
*/
func (client *AipImage) SameHqUpdateContSign(contSign string, options map[string]string) string {
	reqUrl = GetUrlBuild(samehqupdateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["cont_sign"] = contSign
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(productaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["brief"] = brief
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(productaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["brief"] = brief
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(productsearch, client.accessToken)
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
	商品图片搜索—检索
	参数为url
*/
func (client *AipImage) ProductSearchUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(productsearch, client.accessToken)
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
	商品图片搜索—删除
	参数为本地图片
*/
func (client *AipImage) ProductDelete(image string) string {
	reqUrl = GetUrlBuild(productdeleteurl, client.accessToken)
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
	商品图片搜索—删除
	参数为url
*/
func (client *AipImage) ProductDeleteUrl(url string) string {
	reqUrl = GetUrlBuild(productdeleteurl, client.accessToken)
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
	商品图片搜索—删除
	参数为cont_sign
*/
func (client *AipImage) ProductDeleteContSign(contSign string) string {
	reqUrl = GetUrlBuild(productdeleteurl, client.accessToken)
	options := map[string]string{
		"cont_sign": contSign,
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
	商品图片搜索—更新
	参数为本地图片
*/
func (client *AipImage) ProductUpdate(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(productupdateurl, client.accessToken)
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
	商品图片搜索—更新
	参数为url
*/
func (client *AipImage) ProductUpdateUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(productupdateurl, client.accessToken)
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
	商品图片搜索—更新
	参数为cont_sign
*/
func (client *AipImage) ProductUpdateContSign(cont_sign string, options map[string]string) string {
	reqUrl = GetUrlBuild(productupdateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["cont_sign"] = cont_sign
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(picturebookaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["brief"] = brief
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(picturebookaddurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	options["brief"] = brief
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(picturebooksearchurl, client.accessToken)
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
	绘本图片搜索—检索
	参数为url
*/
func (client *AipImage) PicturebookSearchUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(picturebooksearchurl, client.accessToken)
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
	绘本图片搜索—删除
	参数为本地图片
*/
func (client *AipImage) PicturebookDelete(image string) string {
	reqUrl = GetUrlBuild(picturebookdeleteurl, client.accessToken)
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
	绘本图片搜索—删除
	参数为url
*/
func (client *AipImage) PicturebookDeleteUrl(url string) string {
	reqUrl = GetUrlBuild(picturebookdeleteurl, client.accessToken)
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
	绘本图片搜索—删除
	参数为cont_sign
*/
func (client *AipImage) PicturebookDeleteContSign(contSign string) string {
	reqUrl = GetUrlBuild(picturebookdeleteurl, client.accessToken)
	options := map[string]string{
		"cont_sign": contSign,
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
	绘本图片搜索—更新
*/
func (client *AipImage) PicturebookUpdate(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(picturebookupdateurl, client.accessToken)
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
	绘本图片搜索—更新
	参数为url
*/
func (client *AipImage) PicturebookUpdateUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(picturebookupdateurl, client.accessToken)
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
	绘本图片搜索—更新
	参数为cont_sign
*/
func (client *AipImage) PicturebookUpdateUrlContSign(contSign string, options map[string]string) string {
	reqUrl = GetUrlBuild(picturebookupdateurl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["cont_sign"] = contSign
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(carurl, client.accessToken)
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
	车型识别
	参数为url
*/
func (client *AipImage) CarUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(carurl, client.accessToken)
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
	车辆检测
	参数为本地图片
*/
func (client *AipImage) VehicleDetect(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(vehicledetecturl, client.accessToken)
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
	车辆检测
	参数为url
*/
func (client *AipImage) VehicleDetectUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(vehicledetecturl, client.accessToken)
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
	车辆外观损伤识别
	参数为本地图片
*/
func (client *AipImage) VehicleDamage(image string) string {
	reqUrl = GetUrlBuild(vehicledamageurl, client.accessToken)
	options := map[string]string{}
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
	车辆外观损伤识别
	参数为url
*/
func (client *AipImage) VehicleDamageUrl(url string) string {
	reqUrl = GetUrlBuild(vehicledamageurl, client.accessToken)
	options := map[string]string{}
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
	车流统计
	参数为本地图片
*/
func (client *AipImage) TrafficFlow(image string, caseId int, caseInit string, area string, options map[string]interface{}) string {
	reqUrl = GetUrlBuild(trafficflowurl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["image"] = image
	options["case_id"] = caseId
	options["case_init"] = caseInit
	options["area"] = area
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(trafficflowurl, client.accessToken)
	if options == nil {
		options = map[string]interface{}{}
	}
	options["url"] = url
	options["case_id"] = caseId
	options["case_init"] = caseInit
	options["area"] = area
	host = NewHttpSend(reqUrl)
	host.SetBody(options)
	res, err = host.Post()
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
	reqUrl = GetUrlBuild(vehicleattrurl, client.accessToken)
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
	车辆属性识别
	参数为url
*/
func (client *AipImage) VehicleAttrUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(vehicleattrurl, client.accessToken)
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
	车辆分割
	参数为本地图片
*/
func (client *AipImage) VehicleSeg(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(vehiclesegurl, client.accessToken)
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
	车辆分割
	参数为url
*/
func (client *AipImage) VehicleSegUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(vehiclesegurl, client.accessToken)
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
	车辆检测-高空版
	参数为本地图片
*/
func (client *AipImage) VehicleDetectHigh(image string, options map[string]string) string {
	reqUrl = GetUrlBuild(vehicledetecthighurl, client.accessToken)
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
	车辆检测-高空版
	参数为url
*/
func (client *AipImage) VehicleDetectHighUrl(url string, options map[string]string) string {
	reqUrl = GetUrlBuild(vehicledetecthighurl, client.accessToken)
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
