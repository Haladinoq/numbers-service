package rest

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestErrorREST(t *testing.T) {
	table := []struct {
		Purpose, Code, Msg string
		Err                error
		ExpCode            int
		Exp                string
	}{
		{"1. OK", Payload, "invalid payload", errors.New("invalid payload at field: email"), http.StatusBadRequest,
			`{"type":"validation.payload","reason":"invalid payload at field: email","message":"invalid payload"}`,
		},
	}
	for _, x := range table {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		
		Fail(c, x.Code, x.Msg, x.Err)

		resp := w.Result()

		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(resp.Body)
		assert.Nil(t, err)

		actual := buf.String()
		assert.EqualValues(t, x.Exp, actual, x.Purpose, actual)
		assert.EqualValues(t, x.ExpCode, resp.StatusCode, x.Purpose)
	}
}
