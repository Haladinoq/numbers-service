package sql

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/model"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers/data/persistence/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"testing"
	"time"
)

func TestNumbersSQLRepo_ValidateReservationsNumber(t *testing.T) {

	var (
		reservation = &model.Reservation{
			ID:        1,
			Client:    "Client",
			Number:    1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}

		findSQL = `SELECT * FROM "reservation" WHERE client = $1 or "number" = $2`
	)

	mock, gormDB, err := InitialConection()
	require.NoError(t, err)
	mock.ExpectQuery(regexp.QuoteMeta(findSQL)).WithArgs("client", int64(1)).
		WillReturnRows(sqlmock.NewRows([]string{
			"id",
			"client",
			"number",
			"created_at",
			"update_at",
		}).AddRow(
			reservation.ID,
			reservation.Client,
			reservation.Number,
			reservation.CreatedAt,
			reservation.UpdatedAt,
		))
	mockErr, gormDBErr, err := InitialConection()
	mockErr.ExpectQuery(regexp.QuoteMeta(findSQL)).WithArgs("client", int64(1)).WillReturnError(assert.AnError)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		client string
		number int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Reservation
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db: gormDB},
			args:    args{client: "client", number: 1},
			want:    reservation,
			wantErr: false,
		}, {
			name:    "success",
			fields:  fields{db: gormDBErr},
			args:    args{client: "client", number: 1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbersRepo := NumbersSQLRepo{
				db: tt.fields.db,
			}
			got, err := numbersRepo.ValidateReservationsNumber(tt.args.client, tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateReservationsNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateReservationsNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbersSQLRepo_ReserveNumber(t *testing.T) {

	var (
		insertSQL   = `INSERT INTO "reservation" ("client","number","created_at","updated_at","id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`
		reservation = &model.Reservation{
			ID:        1,
			Client:    "client",
			Number:    1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
	)

	mock, gormDB, errRole := InitialConection()
	require.NoError(t, errRole)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(insertSQL)).
		WithArgs(
			reservation.Client,
			reservation.Number,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			reservation.ID).WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow(int64(1)))
	mock.ExpectCommit()

	mockErr, gormErrDB, errRole := InitialConection()
	require.NoError(t, errRole)
	mockErr.ExpectBegin()
	mockErr.ExpectQuery(regexp.QuoteMeta(insertSQL)).WithArgs(
		reservation.Client,
		reservation.Number,
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		reservation.ID).WillReturnError(assert.AnError)
	mockErr.ExpectRollback()

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		reservation *model.Reservation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success",
			fields:  fields{db: gormDB},
			args:    args{reservation: reservation},
			wantErr: assert.NoError,
		}, {
			name:    "error",
			fields:  fields{db: gormErrDB},
			args:    args{reservation: reservation},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbersRepo := NumbersSQLRepo{
				db: tt.fields.db,
			}
			tt.wantErr(t, numbersRepo.ReserveNumber(tt.args.reservation), fmt.Sprintf("ReserveNumber(%v)", tt.args.reservation))
		})
	}
}

func TestNumbersSQLRepo_GetReservationsNumber(t *testing.T) {
	var (
		reservations = []*model.Reservation{{
			ID:        1,
			Client:    "Client",
			Number:    1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}}

		findSQL = `SELECT * FROM "reservation`
	)

	mock, gormDB, err := InitialConection()
	require.NoError(t, err)
	mock.ExpectQuery(regexp.QuoteMeta(findSQL)).WithArgs().
		WillReturnRows(sqlmock.NewRows([]string{
			"id",
			"client",
			"number",
			"created_at",
			"update_at",
		}).AddRow(
			reservations[0].ID,
			reservations[0].Client,
			reservations[0].Number,
			reservations[0].CreatedAt,
			reservations[0].UpdatedAt,
		))
	mockErr, gormDBErr, err := InitialConection()
	mockErr.ExpectQuery(regexp.QuoteMeta(findSQL)).WithArgs().WillReturnError(assert.AnError)

	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.Reservation
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success",
			fields:  fields{db: gormDB},
			want:    reservations,
			wantErr: assert.NoError,
		}, {
			name:    "success",
			fields:  fields{db: gormDBErr},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbersRepo := NumbersSQLRepo{
				db: tt.fields.db,
			}
			got, err := numbersRepo.GetReservationsNumber()
			if !tt.wantErr(t, err, fmt.Sprintf("GetReservationsNumber()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetReservationsNumber()")
		})
	}
}

func TestNewNumbers(t *testing.T) {
	_, gormDBErr, _ := InitialConection()
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want repo.INumbersRepo
	}{{
		name: "succes",
		args: args{db: gormDBErr},
		want: &NumbersSQLRepo{
			db: gormDBErr,
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewNumbers(tt.args.db), "NewNumbers(%v)", tt.args.db)
		})
	}
}

// InitialConection mock connection
func InitialConection() (sqlmock.Sqlmock, *gorm.DB, error) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return mock, gormDB, nil
}
