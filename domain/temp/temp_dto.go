package temp

type WeatherReport struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		Temperature int `json:"temperature"`
	} `json:"current"`
}
