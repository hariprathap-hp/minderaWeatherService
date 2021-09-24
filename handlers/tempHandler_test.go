package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyCity(t *testing.T) {
	city := ""
	result := validate(city)
	assert.NotNil(t, result)
}

func TestValidateCityAsInteger(t *testing.T) {
	city := "1234"
	result := validate(city)
	assert.NotNil(t, result)
}

func TestValidateValidCityAndKey(t *testing.T) {
	city := "Melbourne"
	result := validate(city)
	assert.Nil(t, result)
}
