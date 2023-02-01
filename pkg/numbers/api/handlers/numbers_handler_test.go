package handlers

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/mocks"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNumbersAPIHandler_ReserveNumber(t *testing.T) {
	mockNumbers := new(mocks.INumbersService)
	mockNumbers.On("ReserveNumber", mock.Anything).
		Return(nil).Once()

	r := gin.Default()

	handler := NewNumbersApiHandler(mockNumbers)
	NewNumbersHandlerRoutes(r, handler)

	w := httptest.NewRecorder()

	body := `{
		  "client": "patricia",
		  "number": 2
		}`

	req := newRequest(http.MethodPost, "/v1/reservation", body)
	r.ServeHTTP(w, req)

	result := `{"result":"success"}`

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, result, w.Body.String())
}

func TestNumbersAPIHandler_ReserveNumber_ErrorBinding(t *testing.T) {
	mockNumbers := new(mocks.INumbersService)
	mockNumbers.On("ReserveNumber", mock.Anything).
		Return(nil).Once()

	r := gin.Default()

	handler := NewNumbersApiHandler(mockNumbers)
	NewNumbersHandlerRoutes(r, handler)

	w := httptest.NewRecorder()

	body := `{
		  "client": "patricia",
		  "number": "aaaaaa"
		}`

	req := newRequest(http.MethodPost, "/v1/reservation", body)
	r.ServeHTTP(w, req)

	result := `{"type":"validation.payload","reason":"json: cannot unmarshal string into Go struct field NumbersRequest.number of type int64","message":"something happened while reserving the number"}`

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, result, w.Body.String())
}
func TestNumbersAPIHandler_ReserveNumber_ErrorService(t *testing.T) {
	mockNumbers := new(mocks.INumbersService)
	mockNumbers.On("ReserveNumber", mock.Anything).
		Return(errors.New("error")).Once()

	r := gin.Default()

	handler := NewNumbersApiHandler(mockNumbers)
	NewNumbersHandlerRoutes(r, handler)

	w := httptest.NewRecorder()

	body := `{
		  "client": "patricia",
		  "number": 2
		}`

	req := newRequest(http.MethodPost, "/v1/reservation", body)
	r.ServeHTTP(w, req)

	result := `{"type":"numbers-service.application","reason":"error","message":"something happened while reserving the number"}`

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, result, w.Body.String())
}

func newRequest(method, uri, body string) *http.Request {
	req, _ := http.NewRequest(method, uri, bytes.NewBufferString(body))
	ctx := req.Context()
	req = req.WithContext(ctx)
	return req
}

func TestNumbersAPIHandler_GetNumbers(t *testing.T) {
	var (
		reservations = []*model.NumbersDTO{{
			ID:        1,
			Client:    "Client",
			Number:    1,
			CreatedAt: time.Time{}.UnixMilli(),
			UpdatedAt: time.Time{}.UnixMilli(),
		}}
	)

	mockNumbers := new(mocks.INumbersService)
	mockNumbers.On("GetReservationNumbers").
		Return(reservations, nil).Once()

	r := gin.Default()

	handler := NewNumbersApiHandler(mockNumbers)
	NewNumbersHandlerRoutes(r, handler)

	w := httptest.NewRecorder()

	body := ``

	req := newRequest(http.MethodGet, "/v1/reservation", body)
	r.ServeHTTP(w, req)

	result := `[{"id":1,"client":"Client","number":1,"CreatedAt":-62135596800000,"UpdatedAt":-62135596800000}]`

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, result, w.Body.String())
}

func TestNumbersAPIHandler_GetNumbers_ErrorService(t *testing.T) {

	mockNumbers := new(mocks.INumbersService)
	mockNumbers.On("GetReservationNumbers").
		Return(nil, errors.New("error")).Once()

	r := gin.Default()

	handler := NewNumbersApiHandler(mockNumbers)
	NewNumbersHandlerRoutes(r, handler)

	w := httptest.NewRecorder()

	body := ``

	req := newRequest(http.MethodGet, "/v1/reservation", body)
	r.ServeHTTP(w, req)

	result := `[{"id":1,"client":"Client","number":1,"CreatedAt":-62135596800000,"UpdatedAt":-62135596800000}]`

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, result, w.Body.String())
}
