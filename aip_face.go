/* Copyright 2021 CoderLee Inc. All Rights Reserved. */
/* aip_face.go - 人脸模块 */
/* modification history
-----------------------
2021/6/30, by coderlee, 创建
*/
/*
DESCRIPTION
	人脸模块的封装，包含人脸所有API
*/

package BaiDu_Ai

import (
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

func NewAipFace(_appId string, _apiKey string, _secretKey string) (*AipFace, error) {
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
	return &AipFace{
		appId:       _appId,
		apiKey:      _apiKey,
		secretKey:   _secretKey,
		accessToken: t,
	}, nil
}

func (client *AipFace) Detect(image string, imageType string, options map[string]string) string {
	reqUrl = GetUrlBuild(detectUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = imageType
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) Match(images []map[string]string) string {
	reqUrl = GetUrlBuild(matchUrl, client.accessToken)
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(images)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) Search(image string, imageType string, groupIdList string, options map[string]string) string {
	reqUrl = GetUrlBuild(searchUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = imageType
	options["group_id_list"] = groupIdList
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) MultiSearch(image string, imageType string, groupIdList string, options map[string]string) string {
	reqUrl = GetUrlBuild(multiSearchUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = imageType
	options["group_id_list"] = groupIdList
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) AddUser(image string, imageType string, groupId string, userId string, options map[string]string) string {
	reqUrl = GetUrlBuild(addUserUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = imageType
	options["group_id"] = groupId
	options["user_id"] = userId
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) UpdateUser(image string, imageType string, groupId string, userId string, options map[string]string) string {
	reqUrl = GetUrlBuild(updateUserUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = imageType
	options["group_id"] = groupId
	options["user_id"] = userId
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
func (client *AipFace) FaceDelete(user_id string, group_id string, face_token string) string {
	reqUrl = GetUrlBuild(faceDeleteUrl, client.accessToken)
	options := map[string]string{
		"user_id":    user_id,
		"group_id":   group_id,
		"face_token": face_token,
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

func (client *AipFace) GetUser(user_id string, group_id string) string {
	reqUrl = GetUrlBuild(getUserUrl, client.accessToken)
	options := map[string]string{
		"user_id":  user_id,
		"group_id": group_id,
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

func (client *AipFace) FaceGetlist(user_id string, group_id string) string {
	reqUrl = GetUrlBuild(faceGetlistUrl, client.accessToken)
	options := map[string]string{
		"user_id":  user_id,
		"group_id": group_id,
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

func (client *AipFace) GetGroupUsers(group_id string, options map[string]string) string {
	reqUrl = GetUrlBuild(getGroupUsersUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["group_id"] = group_id
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
func (client *AipFace) UserCopy(user_id string, options map[string]string) string {
	reqUrl = GetUrlBuild(userCopyUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["user_id"] = user_id
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) DeleteUser(group_id string, user_id string) string {
	reqUrl = GetUrlBuild(deleteUserUrl, client.accessToken)
	options := map[string]string{
		"group_id": group_id,
		"user_id":  user_id,
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

func (client *AipFace) GroupAdd(group_id string) string {
	reqUrl = GetUrlBuild(groupAddUrl, client.accessToken)
	options := map[string]string{
		"group_id": group_id,
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

func (client *AipFace) GroupDelete(group_id string) string {
	reqUrl = GetUrlBuild(groupDeleteUrl, client.accessToken)
	options := map[string]string{
		"group_id": group_id,
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

func (client *AipFace) GetGroupList(options map[string]string) string {
	reqUrl = GetUrlBuild(getGroupListUrl, client.accessToken)
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) PersonVerify(image string, image_type string, id_card_number string, name string, options map[string]string) string {
	reqUrl = GetUrlBuild(personVerifyUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = image_type
	options["id_card_number"] = id_card_number
	options["name"] = name
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) Sessioncode(options map[string]string) string {
	reqUrl = GetUrlBuild(sessioncodeUrl, client.accessToken)
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func (client *AipFace) Faceverify(image string, image_type string, options map[string]string) string {
	reqUrl = GetUrlBuild(faceverifyUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["image"] = image
	options["image_type"] = image_type
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
func (client *AipFace) VideoVerify(video_base64 string, options map[string]string) string {
	reqUrl = GetUrlBuild(videoverifyUrl, client.accessToken)
	if options == nil {
		options = map[string]string{}
	}
	options["video_base64"] = video_base64
	host = NewHttpSend(reqUrl)
	host.SetSendType("JSON")
	host.SetBody(options)
	res, err = host.Post()
	if err != nil {
		return err.Error()
	}
	return string(res)
}
