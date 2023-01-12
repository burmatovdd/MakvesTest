package methods

import (
	"encoding/csv"
	"os"
)

func (service *Service) Parser(idArray []int) Result {
	result := Result{}

	file, err := os.Open("ueba.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	for {
		record, e := reader.Read()
		if e != nil {
			panic(e)
		}

		answer, id := service.CheckId(record, idArray)
		if answer {
			result.Result = append(result.Result, Item{
				Id:   id,
				Info: record,
			})
		}
	}
	return result
}
