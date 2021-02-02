package dao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type OilpriceNodes struct {
	OilpriceNodes []Oilprice `json:"oilprices"`
}

type Oilprice struct {
	Date  string  `json: "date"`
	Price float64 `json: "price"`
}

func GetOPData() OilpriceNodes {

	file, err := ioutil.ReadFile("sample/opdata.json")

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("file: ", file)

	data := OilpriceNodes{}

	_ = json.Unmarshal([]byte(file), &data)

	return data
}
