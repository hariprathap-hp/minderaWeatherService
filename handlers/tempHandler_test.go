package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyKey(t *testing.T) {
	access_key := ""
	result := validate(access_key, "")
	assert.NotNil(t, result)
}

func TestValidateEmptyCity(t *testing.T) {
	city := ""
	result := validate("access_key", city)
	assert.NotNil(t, result)
}

func TestValidateCityAsInteger(t *testing.T) {
	city := "1234"
	result := validate("access_key", city)
	assert.NotNil(t, result)
}

func TestValidateValidCityAndKey(t *testing.T) {
	access_key := "test_access_key"
	city := "Melbourne"
	result := validate(access_key, city)
	assert.Nil(t, result)
}
