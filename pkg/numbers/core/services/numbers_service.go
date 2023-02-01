package services

import (
	model2 "gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/model"
)

//INumbersService interface
type INumbersService interface {
	ReserveNumber(numbers *model2.NumbersDTO) error
	GetReservationNumbers() ([]*model2.NumbersDTO, error)
}
