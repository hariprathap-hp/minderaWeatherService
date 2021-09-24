package rest_client

import (
	"errors"
	"net/http"
)

var (
	MockupsEnabled = false
	mocks          = make(map[string]*Mock)
)

type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	Err        error
}

func StartMockups() {
	MockupsEnabled = true
}

func StopMockups() {
	MockupsEnabled = false
}

func AddMockups(mock Mock) {
	mocks[mock.URL] = &mock
}

func FlushMockups() {
	mocks = make(map[string]*Mock)
}

func GetMocks() map[string]*Mock {
	return mocks
}
func Get(url string) (*http.Response, error) {
	if MockupsEnabled {
		mock := mocks[url]
		if mock == nil {
			return nil, errors.New("no mockup found for given request")
		}
		return mock.Response, mock.Err
	}
	client := http.Client{}
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	return client.Do(request)
}
