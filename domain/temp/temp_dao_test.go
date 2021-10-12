package temp

import (
	"errors"
	"hariprathap-hp/minderaWeatherService/api"
	"hariprathap-hp/minderaWeatherService/clients/rest_client"
	"io/ioutil"

	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	rest_client.StartMockups()
	os.Exit(m.Run())
}

func TestWeatherStackAPIFail(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockups(rest_client.Mock{
		URL:        api.GetWeatherStackURL(""),
		HTTPMethod: http.MethodGet,
		Err:        errors.New("error fetching weather info from both weather stack and open weather apis"),
	})
	result, err := GetWeatherReport("")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error fetching weather info from both weather stack and open weather apis", err.Error)
	rest_client.FlushMockups()
}

func TestOpenWeatherAPIFail(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockups(rest_client.Mock{
		URL:        api.GetOpenWeatherURL(""),
		HTTPMethod: http.MethodGet,
		Err:        errors.New("error fetching weather info from both weather stack and open weather apis"),
	})
	result, err := getOpenWeatherReport("")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Status, http.StatusInternalServerError)
	assert.EqualValues(t, err.Error, "error fetching weather info from both weather stack and open weather apis")
	rest_client.FlushMockups()
}

func TestWeatherStackAPISuccess(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockups(rest_client.Mock{
		URL:        api.GetWeatherStackURL(""),
		HTTPMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"current":{"temperature":17,"wind_speed":43}}`)),
		},
	})
	result, err := GetWeatherReport("")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 17, result.Temperature)
	assert.EqualValues(t, 43, result.WindSpeed)
}

func TestOpenWeatherAPISuccess(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockups(rest_client.Mock{
		URL:        api.GetOpenWeatherURL(""),
		HTTPMethod: http.MethodGet,
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"main":{"temp":373.15},"wind":{"speed":10}}`)),
		},
	})
	result, err := getOpenWeatherReport("")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 100, result.Temperature)
	assert.EqualValues(t, 36, result.WindSpeed)
}
