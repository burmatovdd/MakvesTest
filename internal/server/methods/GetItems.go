package methods

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type Id struct {
	Id []int
}

func (service *Service) GetItems(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	idArray := Id{}
	err = json.Unmarshal(jsonData, &idArray)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(service.JsonParser(service.Parser(idArray.Id))))

}
