package mappers

import (
	modelCore "gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/model"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/model"
	"reflect"
	"testing"
	"time"
)

func TestCovertNumbersDtoToDataModel(t *testing.T) {
	type args struct {
		numbersDTO *modelCore.NumbersDTO
	}
	tests := []struct {
		name string
		args args
		want *model.Reservation
	}{
		{
			name: "success",
			args: args{
				numbersDTO: &modelCore.NumbersDTO{
					Client: "Client",
					Number: 1,
				},
			},
			want: &model.Reservation{
				Client: "Client",
				Number: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CovertNumbersDtoToDataModel(tt.args.numbersDTO); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CovertNumbersDtoToDataModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCovertNumbersDataModelToDto(t *testing.T) {
	type args struct {
		numbersModel []*model.Reservation
	}
	tests := []struct {
		name string
		args args
		want []*modelCore.NumbersDTO
	}{
		{
			name: "success",
			args: args{
				numbersModel: []*model.Reservation{{
					ID:        1,
					Client:    "Client",
					Number:    1,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}},
			},
			want: []*modelCore.NumbersDTO{{
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
			if got := CovertNumbersDataModelToDto(tt.args.numbersModel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CovertNumbersDataModelToDto() = %v, want %v", got, tt.want)
			}
		})
	}
}
