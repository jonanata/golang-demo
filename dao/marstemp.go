package dao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type MarstempNodes struct {
	MarstempNodes []Marstemp `json:"marstemps"`
}

type Marstemp struct {
	Date   string  `json: "date"`
	Degree float64 `json: "degree"`
}

func GetMTData() MarstempNodes {

	file, err := ioutil.ReadFile("sample/mtdata.json")

	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("file: ", file)

	data := MarstempNodes{}

	_ = json.Unmarshal([]byte(file), &data)

	/* for i := 0; i < len(data.MarstempNodes); i++ {
		fmt.Println("Date: ", data.MarstempNodes[i].Date)
		fmt.Println("Degree: ", data.MarstempNodes[i].Degree)
	} */

	return data
}
