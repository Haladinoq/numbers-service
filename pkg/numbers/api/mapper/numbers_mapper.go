package mapper

import (
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/api/model"
	modelCore "gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/model"
)

// ConvertNumbersRequestToDTO mapper to model core
func ConvertNumbersRequestToDTO(numbersRequest *model.NumbersRequest) *modelCore.NumbersDTO {
	return &modelCore.NumbersDTO{
		Client: numbersRequest.Client,
		Number: numbersRequest.Number,
	}

}

// ConvertNumbersDTOToResponse ...
func ConvertNumbersDTOToResponse(numbersDTO []*modelCore.NumbersDTO) []*model.NumbersResponse {
	var result []*model.NumbersResponse

	for _, number := range numbersDTO {
		numberResponse := ConvertNumberDTOToResponse(number)
		result = append(result, numberResponse)
	}
	return result

}

// ConvertNumberDTOToResponse ...
func ConvertNumberDTOToResponse(numbersDTO *modelCore.NumbersDTO) *model.NumbersResponse {
	return &model.NumbersResponse{
		ID:        numbersDTO.ID,
		Client:    numbersDTO.Client,
		Number:    numbersDTO.Number,
		CreatedAt: numbersDTO.CreatedAt,
		UpdatedAt: numbersDTO.UpdatedAt,
	}

}
