package face

import (
	"../auth"
	comutil "../comutil"
	"../httpUtil"
	"encoding/json"
	"errors"
	"strings"
)

type AipFace struct {
	appId       string
	apiKey      string
	secretKey   string
	accessToken map[string]string
}

//url
//人脸检测
const detectUrl = "https://aip.baidubce.com/rest/2.0/face/v3/detect"

//人脸对比
const matchUrl = "https://aip.baidubce.com/rest/2.0/face/v3/match"

//人脸搜索
const searchUrl = "https://aip.baidubce.com/rest/2.0/face/v3/search"

//人脸搜索 M:N
const multiSearchUrl = "https://aip.baidubce.com/rest/2.0/face/v3/multi-search"

//人脸库
//人脸注册
const addUserUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add"

//人脸更新
const updateUserUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/update"

//人脸删除
const faceDeleteUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/face/delete"

//用户信息查询
const getUserUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/get"

//获取用户人脸列表
const faceGetlistUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/face/getlist"

//获取用户列表
const getGroupUsersUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/getusers"

//复制用户
const userCopyUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/copy"

//删除用户
const deleteUserUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/delete"

//创建用户组
const groupAddUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/add"

//删除用户组
const groupDeleteUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/delete"

//组列表查询
const getGroupListUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/getlist"

//人脸实名认证
const personVerifyUrl = "https://aip.baidubce.com/rest/2.0/face/v3/person/verify"

//随机校验码接口
const sessioncodeUrl = "https://aip.baidubce.com/rest/2.0/face/v1/faceliveness/sessioncode"

//在线图片活体检测
const faceverifyUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceverify"

//视频活体检测
const videoverifyUrl = "https://aip.baidubce.com/rest/2.0/face/v1/faceliveness/verify"

func NewAipFace(_appId string, _apiKey string, _secretKey string) (*AipFace, error) {
	res := auth.Oauth(_apiKey, _secretKey)
	data := map[string]string{}
	json.Unmarshal([]byte(res), &data)
	keys := comutil.GetKeys(data)
	if !comutil.In("access_token", keys) {
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
	return &AipFace{
		appId:       _appId,
		apiKey:      _apiKey,
		secretKey:   _secretKey,
		accessToken: t,
	}, nil
}

func (client *AipFace) Detect(image string, imageType string, options map[string]string) string {
	require := httpUtil.GetUrlBuild(detectUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = imageType
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) Match(images []map[string]string) string {
	require := httpUtil.GetUrlBuild(matchUrl, client.accessToken)
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(images)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) Search(image string, imageType string, groupIdList string, options map[string]string) string {
	require := httpUtil.GetUrlBuild(searchUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = imageType
	options["group_id_list"] = groupIdList
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) MultiSearch(image string, imageType string, groupIdList string, options map[string]string) string {
	require := httpUtil.GetUrlBuild(multiSearchUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = imageType
	options["group_id_list"] = groupIdList
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) AddUser(image string, imageType string, groupId string, userId string, options map[string]string) string {
	require := httpUtil.GetUrlBuild(addUserUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = imageType
	options["group_id"] = groupId
	options["user_id"] = userId
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) UpdateUser(image string, imageType string, groupId string, userId string, options map[string]string) string {
	require := httpUtil.GetUrlBuild(updateUserUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = imageType
	options["group_id"] = groupId
	options["user_id"] = userId
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
func (client *AipFace) FaceDelete(user_id string, group_id string, face_token string) string {
	require := httpUtil.GetUrlBuild(faceDeleteUrl, client.accessToken)
	options := map[string]string{
		"user_id":    user_id,
		"group_id":   group_id,
		"face_token": face_token,
	}
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) GetUser(user_id string, group_id string) string {
	require := httpUtil.GetUrlBuild(getUserUrl, client.accessToken)
	options := map[string]string{
		"user_id":  user_id,
		"group_id": group_id,
	}
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) FaceGetlist(user_id string, group_id string) string {
	require := httpUtil.GetUrlBuild(faceGetlistUrl, client.accessToken)
	options := map[string]string{
		"user_id":  user_id,
		"group_id": group_id,
	}
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) GetGroupUsers(group_id string, options map[string]string) string {
	require := httpUtil.GetUrlBuild(getGroupUsersUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["group_id"] = group_id
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
func (client *AipFace) UserCopy(user_id string, options map[string]string) string {
	require := httpUtil.GetUrlBuild(userCopyUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["user_id"] = user_id
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) DeleteUser(group_id string, user_id string) string {
	require := httpUtil.GetUrlBuild(deleteUserUrl, client.accessToken)
	options := map[string]string{
		"group_id": group_id,
		"user_id":  user_id,
	}
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) GroupAdd(group_id string) string {
	require := httpUtil.GetUrlBuild(groupAddUrl, client.accessToken)
	options := map[string]string{
		"group_id": group_id,
	}
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) GroupDelete(group_id string) string {
	require := httpUtil.GetUrlBuild(groupDeleteUrl, client.accessToken)
	options := map[string]string{
		"group_id": group_id,
	}
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) GetGroupList(options map[string]string) string {
	require := httpUtil.GetUrlBuild(getGroupListUrl, client.accessToken)
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) PersonVerify(image string, image_type string, id_card_number string, name string, options map[string]string) string {
	require := httpUtil.GetUrlBuild(personVerifyUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = image_type
	options["id_card_number"] = id_card_number
	options["name"] = name
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) Sessioncode(options map[string]string) string {
	require := httpUtil.GetUrlBuild(sessioncodeUrl, client.accessToken)
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) Faceverify(image string, image_type string, options map[string]string) string {
	require := httpUtil.GetUrlBuild(faceverifyUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = image_type
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
func (client *AipFace) VideoVerify(video_base64 string, options map[string]string) string {
	require := httpUtil.GetUrlBuild(videoverifyUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["video_base64"] = video_base64
	h := httpUtil.NewHttpSend(require)
	h.SetSendType("JSON")
	h.SetBody(options)
	res, err := h.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
