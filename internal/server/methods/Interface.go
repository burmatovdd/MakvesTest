package methods

import "github.com/gin-gonic/gin"

type Service struct {
	service *Functions
}

type Functions interface {
	Parser(idArray []int) Result
	CheckId(array []string, id []int) (bool, int)
	GetItems(c *gin.Context)
	JsonParser(result Result) []byte
}
