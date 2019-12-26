package weather

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
)

func GetAllCity() []cityProfile {
	var city []cityProfile
	json.Unmarshal([]byte(cityJson),&city)

	return city
}

func GetWeatherByCity(townID string) interface{} {
	return get(townID)
}

func get(townId string) *weatherData {
	client := &http.Client{}
	req,err := http.NewRequest("GET","http://tj.nineton.cn/Heart/index/all?city="+townId, nil)
	if err != nil {

	}
	response,_ := client.Do(req)
	if response.StatusCode != 200 {
		fmt.Println(err.Error())
	}

	body, _ := ioutil.ReadAll(response.Body)

	var weatherResult weatherData
	json.Unmarshal(body, &weatherResult)
	fmt.Println(weatherResult)

	return &weatherResult
}