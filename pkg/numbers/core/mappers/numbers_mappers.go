package mappers

import (
	modelCore "gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/model"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/model"
)

// CovertNumbersDtoToDataModel ...
func CovertNumbersDtoToDataModel(numbersDTO *modelCore.NumbersDTO) *model.Reservation {
	return &model.Reservation{
		Client: numbersDTO.Client,
		Number: numbersDTO.Number,
	}

}

// CovertNumbersDataModelToDto ...
func CovertNumbersDataModelToDto(numbersModel []*model.Reservation) []*modelCore.NumbersDTO {
	var numbersDTO []*modelCore.NumbersDTO
	for _, numbers := range numbersModel {
		numberDTO := CovertNumberDataModelToDto(numbers)
		numbersDTO = append(numbersDTO, numberDTO)

	}
	return numbersDTO
}

func CovertNumberDataModelToDto(numberModel *model.Reservation) *modelCore.NumbersDTO {
	return &modelCore.NumbersDTO{
		ID:        numberModel.ID,
		Client:    numberModel.Client,
		Number:    numberModel.Number,
		CreatedAt: numberModel.CreatedAt.UnixMilli(),
		UpdatedAt: numberModel.UpdatedAt.UnixMilli(),
	}
}
