package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response REST type.
type Response struct {
	Result interface{} `json:"result"`
}

// OK writes a unique result.
func OK(c *gin.Context, result interface{}) {
	Encode(c, &Response{result})
}

// Encode writes a success json response.
func Encode(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, v)
}

// Created convience must write response 201.
func Created(c *gin.Context, result interface{}) {
	OK(c, result)
}

// ResponseList REST type.
type ResponseList struct {
	Result     interface{} `json:"result"`
	Page       int64       `json:"page"`
	Limit      int64       `json:"limit"`
	Pages      int64       `json:"pages"`
	Elements   int64       `json:"total_elements"`
	Next       bool        `json:"next"`
	Catalogues interface{} `json:"catalogues,omitempty"`
	Filters    interface{} `json:"filters,omitempty"`
}

// List writes a unique result.
func List(c *gin.Context, res *Page) {
	ne := res.Page < res.Pages
	rl := &ResponseList{
		Result:     res.List,
		Page:       res.Page,
		Limit:      res.Limit,
		Pages:      res.Pages,
		Elements:   res.SizeOfList,
		Next:       ne,
		Catalogues: res.Catalogues,
		Filters:    res.Filters,
	}
	c.JSON(http.StatusOK, rl)
}
