package main

import (
	"fmt"
	"github.com/mvity/go-sdk/gaode"
)

func main() {

	gaode.InitGaodeService("ceed620e8e30a3bdd19bdcfe516e8ff0")
	if result, err := gaode.GeocodeGeo(&gaode.GeocodeGeoParam{Address: "郑东中心", City: "100"}); result != nil {
		fmt.Printf("Result: %v\n", result)
		fmt.Println(result.InfoCode)
	} else {
		fmt.Printf("Error: %v\n", err)
	}

}
