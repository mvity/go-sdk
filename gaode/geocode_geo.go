package gaode

import (
	"errors"
	"fmt"
	"github.com/mvity/go-box/x"
	gosdk "github.com/mvity/go-sdk"
)

// GeocodeGeo 地理编码 https://lbs.amap.com/api/webservice/guide/api/georegeo#geo
func GeocodeGeo(param *GeocodeGeoParam) (*GeocodeGeoResult, error) {
	if valid, cause := param.valid(); !valid {
		return nil, errors.New(cause)
	}

	uri := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s&address=%s&city=%s",
		x.StringDefaultIfBlank(param.Key, gaodeKey), param.Address, param.City)
	if ok, resp, _ := x.HttpGet(uri); ok {
		gosdk.WARN(resp)
		result := &GeocodeGeoResult{}
		if err := result.parse(resp); err != nil {
			return nil, err
		}
		if !result.IsSuccess() {
			return result, nil
		}
		if node, err := x.JsonFromStringE(resp); err != nil {
			return nil, err
		} else {
			for i := 0; i < node.Name("geocodes").Size(); i++ {
				sub := node.Name("geocodes").Index(i)
				
			}
		}

		return result, nil
	} else {
		return nil, errors.New(resp)
	}
}

// GeocodeGeoParam 地理编码，请求参数
type GeocodeGeoParam struct {
	gaodeParam
	Address string `label:"结构化地址信息" validate:"required,min=1,max=256"`
	City    string `label:"指定查询的城市" validate:"omitempty,min=1,max=256"`
}

// GeocodeGeoResult  地理编码，响应结果
type GeocodeGeoResult struct {
	gaodeResult
	Geocodes []geocode `json:"geocodes"` // 地理编码信息列表
}

// 地理编码信息
type geocode struct {
	Country  string `json:"country"`  // 国家，国内地址默认返回中国
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市
	District string `json:"district"` // 县区
	Street   string `json:"street"`   // 街道
	Number   string `json:"number"`   // 门牌
	CityCode string `json:"citycode"` // 城市编码
	AdCode   string `json:"adcode"`   // 行政区划码
	Location string `json:"location"` // 坐标点，经度,纬度
	Level    string `json:"level"`    // 匹配级别
}
