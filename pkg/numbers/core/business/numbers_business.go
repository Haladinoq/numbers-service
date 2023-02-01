package business

import (
	"errors"
	logg "github.com/rs/zerolog/log"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/mappers"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/model"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/services"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/repo"
)

const (
	errorNumber = `the number has already been reserved`
	errorClient = `the client already has a reservation`
)

//NumbersBusinessLogic struct
type NumbersBusinessLogic struct {
	numbersRepo repo.INumbersRepo
}

//NewNumbersBusinessLogic ..
func NewNumbersBusinessLogic(numbersRepo repo.INumbersRepo) services.INumbersService {
	return &NumbersBusinessLogic{numbersRepo}
}

func (numbersBusiness NumbersBusinessLogic) ReserveNumber(numbersDTO *model.NumbersDTO) error {

	reservationValidate, e := numbersBusiness.numbersRepo.ValidateReservationsNumber(numbersDTO.Client, numbersDTO.Number)
	if e != nil {
		logg.Error().Caller().Err(e).Msg("")
		return e
	}

	if reservationValidate != nil {
		if numbersDTO.Number == reservationValidate.Number {
			return errors.New(errorNumber)
		}
		if numbersDTO.Client == reservationValidate.Client {
			return errors.New(errorClient)
		}

	}

	numbersModel := mappers.CovertNumbersDtoToDataModel(numbersDTO)
	err := numbersBusiness.numbersRepo.ReserveNumber(numbersModel)
	if err != nil {
		logg.Error().Caller().Err(err).Msg("")
		return err
	}
	return nil

}

func (numbersBusiness NumbersBusinessLogic) GetReservationNumbers() ([]*model.NumbersDTO, error) {
	numbers, e := numbersBusiness.numbersRepo.GetReservationsNumber()
	if e != nil {
		logg.Error().Caller().Err(e).Msg("")
		return nil, e
	}

	numbersDTO := mappers.CovertNumbersDataModelToDto(numbers)
	return numbersDTO, nil
}
