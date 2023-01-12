package methods

import (
	"strconv"
)

func (service *Service) CheckId(array []string, id []int) (bool, int) {
	arrayItem, err := strconv.Atoi(array[0])
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(id); i++ {
		if arrayItem == id[i] {
			return true, id[i]
		}
		continue
	}
	return false, -1
}
