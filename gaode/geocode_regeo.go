package gaode

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mvity/go-box/x"
)

func GeocodeRegeo(param *GeocodeRegeoParam) (*GeocodeRegeoResult, error) {
	if x.StringIsAllBlank(param.Key, gaodeKey) {
		return nil, errors.New("未知的高德Key，请通过gaode.InitGaodeService(key)或param.Key进行设置")
	}
	uri := fmt.Sprintf("https://restapi.amap.com/v3/geocode/geo?key=%s",
		x.StringDefaultIfBlank(param.Key, gaodeKey))
	if ok, resp, _ := x.HttpGet(uri); ok {
		result := &GeocodeRegeoResult{}
		if err := json.Unmarshal([]byte(resp), result); err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return nil, errors.New(resp)
	}

}

type GeocodeRegeoParam struct {
	Key        string // 高德Key
	Location   string
	PoiType    string
	Radius     int
	Extensions string
	RoadLevel  string
	HomeOrCorp string
}

type GeocodeRegeoResult struct {
}
