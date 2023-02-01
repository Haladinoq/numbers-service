package mapper

import (
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/api/model"
	modelCore "gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/model"
	"reflect"
	"testing"
	"time"
)

func TestConvertNumbersRequestToDTO(t *testing.T) {
	type args struct {
		numbersRequest *model.NumbersRequest
	}
	tests := []struct {
		name string
		args args
		want *modelCore.NumbersDTO
	}{
		{
			name: "success",
			args: args{
				numbersRequest: &model.NumbersRequest{
					Client: "Client",
					Number: 1,
				},
			},
			want: &modelCore.NumbersDTO{
				Client: "Client",
				Number: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNumbersRequestToDTO(tt.args.numbersRequest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNumbersRequestToDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertNumbersDTOToResponse(t *testing.T) {
	type args struct {
		numbersDTO []*modelCore.NumbersDTO
	}
	tests := []struct {
		name string
		args args
		want []*model.NumbersResponse
	}{
		{
			name: "success",
			args: args{
				numbersDTO: []*modelCore.NumbersDTO{{
					ID:        1,
					Client:    "Client",
					Number:    1,
					CreatedAt: time.Time{}.UnixMilli(),
					UpdatedAt: time.Time{}.UnixMilli(),
				}},
			},
			want: []*model.NumbersResponse{{
				ID:        1,
				Client:    "Client",
				Number:    1,
				CreatedAt: time.Time{}.UnixMilli(),
				UpdatedAt: time.Time{}.UnixMilli(),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNumbersDTOToResponse(tt.args.numbersDTO); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNumbersDTOToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
