package reader

import (
	"io/ioutil"

	"encoding/json"
)

type Json struct{}

func (Json) Open(filepath string) (output FileContent, err error) {
	var rawFile []byte
	rawFile, err = ioutil.ReadFile(filepath)
	err = json.Unmarshal(rawFile, &output)
	return
}