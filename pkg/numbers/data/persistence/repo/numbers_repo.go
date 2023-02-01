package repo

import "gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/model"

//INumbersRepo interface
type INumbersRepo interface {
	ValidateReservationsNumber(client string, number int64) (*model.Reservation, error)
	ReserveNumber(reservation *model.Reservation) error
	GetReservationsNumber() ([]*model.Reservation, error)
}
