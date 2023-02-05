package gaode

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/mvity/go-box/x"
	gosdk "github.com/mvity/go-sdk"
)

var gaodeKey = ""

func InitGaodeService(key string) {
	gaodeKey = key
}

// 高德接口请求参数
type gaodeParam struct {
	Key string `label:"高德KEY" validate:"omitempty,len=32"`
	Sig string `label:"数字签名" validate:"omitempty,len=32"`
}

func (param *gaodeParam) valid() (bool, string) {
	if x.StringIsAllBlank(param.Key, gaodeKey) {
		return false, "未知的高德Key，请通过gaode.InitGaodeService(key)或param.Key进行设置"
	}
	if err := gosdk.Validate.Struct(param); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, verr := range errs.Translate(gosdk.Translator) {
				return false, verr
			}
		}
	}
	return true, ""
}

// 高德接口通用响应
type gaodeResult struct {
	Status   string `json:"status"`   // 返回结果状态值，返回值为 0 或 1，0 表示请求失败；1 表示请求成功。
	Info     string `json:"info"`     // 返回状态说明，当 status 为 0 时，info 会返回具体错误原因，否则返回“OK”。详情可以参阅 info状态表 [https://lbs.amap.com/api/webservice/guide/tools/info]
	InfoCode string `json:"infocode"` // 返回状态代码，详情可以参阅 info状态表 [https://lbs.amap.com/api/webservice/guide/tools/info]
	Count    string `json:"count"`    // 返回结果数目，返回结果的个数。
}

// IsSuccess 是否成功
func (result *gaodeResult) IsSuccess() bool {
	return result.Status == "1"
}

// parse 转换请求结果
func (result *gaodeResult) parse(jsonStr string) error {
	if err := json.Unmarshal([]byte(jsonStr), result); err != nil {
		return err
	}
	return nil
}
