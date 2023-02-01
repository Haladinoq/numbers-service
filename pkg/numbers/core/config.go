package core

import (
	"github.com/golobby/container/v3"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/business"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/services"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/repo"
)

//RunConfig ..
func RunConfig(container *container.Container) {
	container.Singleton(func(numbersRepo repo.INumbersRepo) services.INumbersService {
		return business.NewNumbersBusinessLogic(numbersRepo)
	})
}
