package rest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSuccessREST(t *testing.T) {
	table := []struct {
		Purpose string
		Item    interface{}
		ExpCode int
		Exp     string
	}{
		{"1. OK", map[string]interface{}{"hi": "hello"}, http.StatusOK,
			`{"result":{"hi":"hello"}}`,
		},
	}
	for _, x := range table {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		OK(c, x.Item)

		resp := w.Result()
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(resp.Body)
		assert.Nil(t, err)

		actual := buf.String()
		assert.EqualValues(t, x.Exp, actual, x.Purpose, actual)
		assert.EqualValues(t, x.ExpCode, resp.StatusCode, x.Purpose)
	}
}
