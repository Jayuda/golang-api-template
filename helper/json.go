package helper

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadFromRequestBody(c *gin.Context, result interface{}) {
	decoder := json.NewDecoder(c.Request.Body)
	_ = decoder.Decode(result)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
