package temp

type WeatherReport struct {
	Current struct {
		Temperature int `json:"temperature"`
		WindSpeed   int `json:"wind_speed"`
	} `json:"current"`
}

type OpenWeatherReport struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

type WeatherInfo struct {
	WindSpeed   float64 `json:"wind_speed"`
	Temperature float64 `json:"temperature_degrees"`
}

func (wi *WeatherInfo) FTOC() {
	wi.WindSpeed = wi.WindSpeed * 10
	wi.Temperature = wi.Temperature - 273.15
}
