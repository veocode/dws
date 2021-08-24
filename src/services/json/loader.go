package json

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Loader struct {
	httpClient *http.Client
	decoder    *json.Decoder
	lastErr    error
}

func NewLoader() *Loader {
	jsonLoader := new(Loader)
	jsonLoader.httpClient = &http.Client{}
	return jsonLoader
}

func (j *Loader) FromURL(url string) *Loader {
	j.lastErr = nil
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		j.lastErr = err
		return j
	}

	request.Header.Add("Accept", "application/json")

	response, err := j.httpClient.Do(request)
	if err != nil {
		j.lastErr = err
		return j
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad response status code: %d; url=%s", response.StatusCode, url)
		j.lastErr = err
		return j
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	j.decoder = json.NewDecoder(response.Body)
	return j
}

func (j *Loader) FromFile(filePath string) *Loader {
	j.lastErr = nil
	jsonFile, err := os.Open(filePath)
	if err != nil {
		j.lastErr = err
		return j
	}

	j.decoder = json.NewDecoder(jsonFile)
	return j
}

func (j *Loader) DecodeTo(result interface{}) error {
	if j.lastErr != nil {
		return j.lastErr
	}

	return j.decoder.Decode(result)
}
