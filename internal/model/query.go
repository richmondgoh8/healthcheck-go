package model

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swagger:parameters HealthRequest
type InputQuery struct {
	//in: query
	//
	//required: true
	IsListening int `json:"isListening"`
}

func (inputQuery *InputQuery) ParseValidate(c *gin.Context) error {
	val, exist := c.GetQuery("isListening")
	if !exist {
		return errors.New("missing key: 'isListening'")
	}
	trueVal, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	inputQuery.IsListening = trueVal

	return nil
}
