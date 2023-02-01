package sql

import (
	"errors"
	logg "github.com/rs/zerolog/log"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/model"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/repo"
	"gorm.io/gorm"
)

const (
	whereReservation = `client = ? or "number" = ?`
)

//NumbersSQLRepo struct
type NumbersSQLRepo struct {
	db *gorm.DB
}

//NewNumbers constructor
func NewNumbers(db *gorm.DB) repo.INumbersRepo {
	return &NumbersSQLRepo{
		db: db,
	}
}

func (numbersRepo NumbersSQLRepo) ValidateReservationsNumber(client string, number int64) (*model.Reservation, error) {
	var reservations model.Reservation
	result := numbersRepo.db.Model(&model.Reservation{}).
		Where(whereReservation, client, number).
		Find(&reservations)

	if result.Error != nil {
		logg.Error().Caller().Err(result.Error).Msg("")
		return nil, errors.New("error loading reservation numbers" + result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &reservations, nil
}

func (numbersRepo NumbersSQLRepo) ReserveNumber(reservation *model.Reservation) error {
	return numbersRepo.db.Transaction(func(tx *gorm.DB) error {
		tx.Statement.Omits = nil

		result := tx.Create(&reservation)
		if result.Error != nil {
			logg.Error().Caller().Err(result.Error).Msg("")
			return errors.New("error reservation numbers: " + result.Error.Error())
		}
		return nil
	})
}

func (numbersRepo NumbersSQLRepo) GetReservationsNumber() ([]*model.Reservation, error) {
	var reservations []*model.Reservation
	result := numbersRepo.db.Model(&model.Reservation{}).
		Find(&reservations)

	if result.Error != nil {
		logg.Error().Caller().Err(result.Error).Msg("")
		return nil, errors.New("error loading reservation numbers" + result.Error.Error())
	}
	return reservations, nil
}
