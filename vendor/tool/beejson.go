package tool

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
)

// 请求json
type RequestJson struct {
	Action   string `json:"action"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Data     string `json:"data"`
}

// 返回Json
type RespondJson struct {
	Status bool        `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

//* 请求结构体解析 *//
type UserToken struct {
	Token string `json:"token"`
}

// 用户登录
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 设置云控账户
type SetCouldJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}

// 删除设备
type DelDeviceJson struct {
	Uuid string `json:"uuid"`
}

type RequestType int

const (
	Setcould RequestType = iota
	AddDevice
	Deldevice
	Editdevice
	Userlogin
	DownRes
)

var (
	requestTypeMap = map[string]RequestType{
		"setcould":   Setcould,   //设置Could
		"device":     AddDevice,  //设备上传
		"deldevice":  Deldevice,  //删除设备
		"editdevice": Editdevice, //编辑设备
		"userlogin":  Userlogin,  //用户登录
		"downres":    DownRes,    //下载资源
	}
)

func NewRequestJson(data string) (RequestJson, error) {
	var request RequestJson
	err := json.Unmarshal([]byte(data), &request)
	if err != nil {
		beego.Error("parse json is", err.Error())
	}
	return request, err
}

func GetRequestType(t string) RequestType {
	return requestTypeMap[t]
}

func (r *RequestJson) Beejson() (interface{}, error) {
	switch requestTypeMap[r.Action] {
	case Setcould:
		var request SetCouldJson
		err := json.Unmarshal([]byte(r.Data), &request)
		if err != nil {
			beego.Error("parse json is", err.Error())
		}
		return request, err
		break
	case AddDevice:
		return nil, nil
		break
	case Deldevice:
		var request DelDeviceJson
		err := json.Unmarshal([]byte(r.Data), &request)
		if err != nil {
			beego.Error("parse json is", err.Error())
		}
		return request, err
		break
	case Editdevice:
		return nil, nil
		break
	default:
	}
	return nil, errors.New("can't find the requesttype type")
}
