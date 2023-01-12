package methods

import (
	"encoding/json"
)

func (service *Service) JsonParser(result Result) []byte {

	resultJson, err := json.Marshal(result)

	if err != nil {
		panic(err)
	}

	return resultJson
}
