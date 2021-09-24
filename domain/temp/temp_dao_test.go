package temp

import (
	"io/ioutil"
	"minderaWeatherService/clients/rest_client"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest_client.StartMockups()
	os.Exit(m.Run())
}

/*func TestWeatherStackAPIFail(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockups(rest_client.Mock{
		URL:        "http://api.weatherstack.com/current?access_key=key&query=city",
		HTTPMethod: http.MethodGet,
		Err:        errors.New(""),
	})
	result, err := GetWeatherReport("", "")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	rest_client.FlushMockups()
}

func TestOpenWeatherAPIFail(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockups(rest_client.Mock{
		URL:        "http://api.openweathermap.org/data/2.5/weather?q=melbourne,AU&appid=uid",
		HTTPMethod: http.MethodGet,
		Err:        errors.New("error fetching weather info from both weather stack and open weather apis"),
	})
	result, err := GetWeatherReport("", "")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Status, http.StatusInternalServerError)
	assert.EqualValues(t, err.Error, "error fetching weather info from both weather stack and open weather apis")
	rest_client.FlushMockups()
}

func TestWeatherStackAPISuccess(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockups(rest_client.Mock{
		URL:        "http://api.weatherstack.com/current?access_key1=key&query=city1",
		HTTPMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"wind_speed": 22,"temperature_degrees": 11}`)),
		},
	})
	_, err := GetWeatherReport("", "")
	assert.Nil(t, err)
	//assert.NotNil(t, result)
	//assert.EqualValues(t, 11, result.Temperature)
	//assert.EqualValues(t, 22, result.WindSpeed)
	rest_client.FlushMockups()
}*/

func TestOpenWeatherAPISuccess(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockups(rest_client.Mock{
		URL:        "http://api.openweathermap.org/data/2.5/weather?q=melbourne1,AU&appid=uid1",
		HTTPMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"wind_speed": 22,"temperature_degrees": 11}`)),
		},
	})
	result, err := GetWeatherReport("")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	//assert.EqualValues(t, 11, result.Temperature)
	//assert.EqualValues(t, 22, result.WindSpeed)
	rest_client.FlushMockups()
}
