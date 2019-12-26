package weather

import (
	"go/types"
)

type weatherData struct {
	Status  string
	Weather []WeatherProfile
}

type WeatherProfile struct {
	City_name   string
	City_id     string
	Last_update string
	Now        nowProfile
	Today      todayProfile
	Future     []futureUniqueProfile
}

//当前
type nowProfile struct {
	Text           string
	Code           string
	Temperature    string
	Feels_like      string
	Wind_direction  string
	Wind_speed      string
	Wind_scale      string
	Humidity       string
	Visibility     string
	Pressure       string
	Pressure_rising string
	Air_quality     airQualityProfile
	Alarms         []interface{}
}

type airQualityProfile struct {
	Stations types.Nil
	City     airCity
}

/**
城市空气质量
 */
type airCity struct {
	Aqi        string
	Pm25       string
	Pm10       string
	So2        string
	No2        string
	Co         string
	O3         string
	Last_update string
	Quality    string
}

/**
今天天气详情
 */
type todayProfile struct {
	Sunrise    string
	Sunset     string
	Suggestion todayProfileSuggestionProfile
}

/**
今天天气建议
 */
type todayProfileSuggestionProfile struct {
	Dressing   suggestionDetailProfile
	Uv         suggestionDetailProfile
	Car_washing suggestionDetailProfile
	Travel     suggestionDetailProfile
	Flu        suggestionDetailProfile
	Sport      suggestionDetailProfile
}

/**
今天天气建议详情
 */
type suggestionDetailProfile struct {
	Brief   string
	Details string
}

/**
未来天气详情
 */
type futureUniqueProfile struct {
	Date  string
	High  string
	Low   string
	Day   string
	Text  string
	Code1 string
	Code2 string
	Cop   string
	Wind  string
}
