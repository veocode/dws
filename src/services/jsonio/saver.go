package jsonio

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Saver struct {
}

func NewSaver() *Saver {
	return new(Saver)
}

func (j *Saver) SaveToFile(data interface{}, filePath string) error {
	jsonText, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, jsonText, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
