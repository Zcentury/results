package results

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type result struct {
	ctx    *gin.Context
	status int
	result map[string]interface{}
}

func Success(c *gin.Context, args ...interface{}) *result {
	r := new(result)
	r.ctx = c
	r.status = http.StatusOK
	r.result = make(map[string]interface{})

	code := 200
	message := "Success"

	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			code = v
		case string:
			message = v
		}
	}
	r.Put("code", code)
	r.Put("message", message)
	return r
}

func Failed(c *gin.Context, args ...interface{}) *result {
	r := new(result)
	r.ctx = c
	r.status = http.StatusOK
	r.result = make(map[string]interface{})

	code := 201
	message := "Failed"
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			code = v
		case string:
			message = v
		}
	}
	r.Put("code", code)
	r.Put("message", message)
	return r
}

func (r *result) Status(status int) *result {
	r.status = status
	return r
}

func (r *result) Put(key string, value interface{}) *result {
	r.result[key] = value
	return r
}

func (r *result) Result() {
	if r.ctx == nil {
		fmt.Println("Gin Context 必须设置")
		return
	}
	r.ctx.JSON(r.status, r.result)
	r.ctx.Abort()
}

func (r *result) Json() (string, error) {
	marshal, err := json.Marshal(r.result)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
