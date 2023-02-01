package business

import (
	"errors"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/model"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/core/services"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/mocks"
	modelData "gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/model"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/repo"
	"reflect"
	"testing"
	"time"
)

func TestNumbersBusinessLogic_ReserveNumber(t *testing.T) {

	var (
		numbersDTO = &model.NumbersDTO{
			Client: "Client",
			Number: 1,
		}

		reservation = &modelData.Reservation{
			Client: "Client",
			Number: 1,
		}

		reservationSameClient = &modelData.Reservation{
			Client: "Client",
			Number: 2,
		}

		reservationSameNumber = &modelData.Reservation{
			Client: "Client2",
			Number: 1,
		}
	)

	mockNumbers := new(mocks.INumbersRepo)
	mockNumbers.On("ValidateReservationsNumber", numbersDTO.Client, numbersDTO.Number).Return(nil, nil).Once()
	mockNumbers.On("ReserveNumber", reservation).Return(nil).Once()

	mockNumbersErr := new(mocks.INumbersRepo)
	mockNumbersErr.On("ValidateReservationsNumber", numbersDTO.Client, numbersDTO.Number).Return(reservationSameClient, nil).Once()

	mockNumbersErr2 := new(mocks.INumbersRepo)
	mockNumbersErr2.On("ValidateReservationsNumber", numbersDTO.Client, numbersDTO.Number).Return(reservationSameNumber, nil).Once()

	mockNumbersErr3 := new(mocks.INumbersRepo)
	mockNumbersErr3.On("ValidateReservationsNumber", numbersDTO.Client, numbersDTO.Number).Return(nil, errors.New("error")).Once()

	mockNumbersErr4 := new(mocks.INumbersRepo)
	mockNumbersErr4.On("ValidateReservationsNumber", numbersDTO.Client, numbersDTO.Number).Return(nil, nil).Once()
	mockNumbersErr4.On("ReserveNumber", reservation).Return(errors.New("error")).Once()

	type fields struct {
		numbersRepo repo.INumbersRepo
	}
	type args struct {
		numbersDTO *model.NumbersDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				numbersRepo: mockNumbers,
			},
			args: args{
				numbersDTO: numbersDTO,
			},
			wantErr: false,
		}, {
			name: "error client exist",
			fields: fields{
				numbersRepo: mockNumbersErr,
			},
			args: args{
				numbersDTO: numbersDTO,
			},
			wantErr: true,
		}, {
			name: "error number exist",
			fields: fields{
				numbersRepo: mockNumbersErr2,
			},
			args: args{
				numbersDTO: numbersDTO,
			},
			wantErr: true,
		}, {
			name: "error validate ",
			fields: fields{
				numbersRepo: mockNumbersErr3,
			},
			args: args{
				numbersDTO: numbersDTO,
			},
			wantErr: true,
		}, {
			name: "error reservation number ",
			fields: fields{
				numbersRepo: mockNumbersErr4,
			},
			args: args{
				numbersDTO: numbersDTO,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbersBusiness := NumbersBusinessLogic{
				numbersRepo: tt.fields.numbersRepo,
			}
			if err := numbersBusiness.ReserveNumber(tt.args.numbersDTO); (err != nil) != tt.wantErr {
				t.Errorf("ReserveNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNumbersBusinessLogic_GetReservationNumbers(t *testing.T) {

	var (
		reservations = []*modelData.Reservation{{
			ID:        1,
			Client:    "Client",
			Number:    1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}}

		want = []*model.NumbersDTO{{
			ID:        1,
			Client:    "Client",
			Number:    1,
			CreatedAt: time.Time{}.UnixMilli(),
			UpdatedAt: time.Time{}.UnixMilli(),
		}}
	)

	mockNumbers := new(mocks.INumbersRepo)
	mockNumbers.On("GetReservationsNumber").Return(reservations, nil).Once()

	mockNumbersErr := new(mocks.INumbersRepo)
	mockNumbersErr.On("GetReservationsNumber").Return(nil, errors.New("error")).Once()

	type fields struct {
		numbersRepo repo.INumbersRepo
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.NumbersDTO
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				numbersRepo: mockNumbers,
			},
			want:    want,
			wantErr: false,
		}, {
			name: "error get reservations",
			fields: fields{
				numbersRepo: mockNumbersErr,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbersBusiness := NumbersBusinessLogic{
				numbersRepo: tt.fields.numbersRepo,
			}
			got, err := numbersBusiness.GetReservationNumbers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReservationNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReservationNumbers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNumbersBusinessLogic(t *testing.T) {
	mockNumbers := new(mocks.INumbersRepo)

	type args struct {
		numbersRepo repo.INumbersRepo
	}
	tests := []struct {
		name string
		args args
		want services.INumbersService
	}{
		{
			name: "success",
			args: args{
				numbersRepo: mockNumbers,
			},
			want: &NumbersBusinessLogic{numbersRepo: mockNumbers},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNumbersBusinessLogic(tt.args.numbersRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNumbersBusinessLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}
