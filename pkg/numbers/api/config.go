package api

import (
	"github.com/golobby/container/v3"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/api/handlers"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/services"
	"gitlab.palo-it.net/palo/numbers-service/pkg/utils"
)

//RunConfig run config of api layer
func RunConfig(ctn *container.Container) {
	ctn.Singleton(func(numbersService services.INumbersService) handlers.NumbersAPIHandler {
		return handlers.NewNumbersApiHandler(numbersService)
	})

	ctn.Call(func(
		ctx utils.Context,
		numbersHandler handlers.NumbersAPIHandler,
	) {
		handlers.NewNumbersHandlerRoutes(ctx.Router, numbersHandler)
	})

}
