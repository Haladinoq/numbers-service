package handlers

import (
	"github.com/gin-gonic/gin"
	logg "github.com/rs/zerolog/log"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/api/mapper"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/api/model"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/services"
	"gitlab.palo-it.net/palo/numbers-service/pkg/rest"
	"net/http"
)

const (
	versionEndpoint = "v1"
	numbersURL      = versionEndpoint + "/reservation"
)

//NumbersAPIHandler handler
type NumbersAPIHandler struct {
	NumbersService services.INumbersService
}

func NewNumbersApiHandler(numbersService services.INumbersService) NumbersAPIHandler {
	return NumbersAPIHandler{NumbersService: numbersService}
}

//NewNumbersHandlerRoutes ..
func NewNumbersHandlerRoutes(router *gin.Engine, numbersHandler NumbersAPIHandler) {
	group := router.Group(numbersURL)
	group.POST("", numbersHandler.ReserveNumber)
	group.GET("", numbersHandler.GetNumbers)
}

//ReserveNumber service for reservation numbers
// @Summary service for reservation numbers
// @Description create numbers
// @Tags Reservation V1
// @Accept  json
// @Produce  json
// @Param numbers body model.NumbersRequest true "The reservation data"
// @Success 200 string rest.MSGSuccess string "success"
// @Success 204 "No Content"
// @Failure 400 "Bad Request"
// @Failure 401 "Unauthorized Request"
// @Failure 404 "Not Found"
// @Router /v1/reservation [post]
func (handlerNumbers NumbersAPIHandler) ReserveNumber(context *gin.Context) {
	var request *model.NumbersRequest
	if err := context.BindJSON(&request); err != nil {
		logg.Error().Caller().Err(err).Msg("")
		rest.Fail(context, rest.Payload, rest.MSGErrorCreate, err)
		return
	}

	numbersDTO := mapper.ConvertNumbersRequestToDTO(request)
	err := handlerNumbers.NumbersService.ReserveNumber(numbersDTO)
	if err != nil {
		logg.Error().Caller().Err(err).Msg("")
		rest.Fail(context, rest.Application, rest.MSGErrorCreate, err)
		return
	}
	rest.Created(context, rest.MSGSuccess)
}

//GetNumbers get reservation numbers
// @Summary get reservation numbers
// @Description get reservation numbers
// @Tags Reservation V1
// @Accept  json
// @Produce  json
// @Success 200 {object} model.NumbersResponse
// @Success 204 "No Content"
// @Failure 400 "Bad Request"
// @Failure 401 "Unauthorized Request"
// @Failure 404 "Not Found"
// @Router /v1/reservation [get]
func (handlerNumbers NumbersAPIHandler) GetNumbers(context *gin.Context) {

	numbers, err := handlerNumbers.NumbersService.GetReservationNumbers()
	if err != nil {
		logg.Error().Caller().Err(err).Msg("")
		rest.Fail(context, rest.Application, rest.MSGErrorOne, err)
		return
	}
	numbersResponse := mapper.ConvertNumbersDTOToResponse(numbers)
	context.JSON(http.StatusOK, numbersResponse)
}
