package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONLoader struct {
	url        string
	httpClient *http.Client
}

func NewJSONLoader() *JSONLoader {
	jsonLoader := new(JSONLoader)
	jsonLoader.httpClient = &http.Client{}
	return jsonLoader
}

func (j *JSONLoader) Load(url string) *JSONLoader {
	j.url = url
	return j
}

func (j *JSONLoader) ToStruct(result interface{}) error {
	if j.url == "" {
		return fmt.Errorf("jsonloader url not specified")
	}

	request, err := http.NewRequest(http.MethodGet, j.url, nil)
	if err != nil {
		return err
	}

	request.Header.Add("Accept", "application/vnd.github.v3+json")

	response, err := j.httpClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad response status code: %d; url=%s", response.StatusCode, j.url)
		return err
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	return json.NewDecoder(response.Body).Decode(result)
}
