package Controllers

import (
	"net/http"

	"io/ioutil"

	"github.com/gin-gonic/gin"

	"encoding/json"
)

type Alert struct {
	sender_name string
	event       string
	start       int
	end         int
	description string
}

type Current struct {
	Dt         int       `json:"dt"`
	Sunrise    int       `json:"sunrise"`
	Sunset     int       `json:"sunset"`
	Temp       float32   `json:"temp"`
	FeelsLike  float32   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float32   `json:"dew_point"`
	Uvi        int       `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float32   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	Weather    []Weather `json:"weather"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherData struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        Current `json:"current"`
	Alerts         []Alert `json:"alerts"`
}

func GetHome(c *gin.Context) {
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/onecall?lat=40.75058192493313&lon=-73.99723493108772&exclude=minutely,hourly,daily&appid=023adc6acf7345d1d3850f5d2d9ee489")
	if err != nil {
		println("An error occurred!")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		var weatherData WeatherData
		json.Unmarshal([]byte(string(body)), &weatherData)
		c.JSON(http.StatusOK, weatherData.Current.Weather[0].Main)
	}
}
