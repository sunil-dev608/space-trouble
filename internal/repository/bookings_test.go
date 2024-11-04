package repository

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sunil-dev608/space-trouble/internal/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	mockDB *sql.DB
	mock   sqlmock.Sqlmock
)

func setupTestDB() *gorm.DB {
	mockDB, mock, _ = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}
	return db
}

var (
	defaultBirthday   = "2000-01-01"
	defaultLaunchDate = "2022-01-01"
)

func Test_bookingRepository_CreateBooking(t *testing.T) {

	defaultBirthdayDate, _ := time.Parse(model.DateLayout, defaultBirthday)
	defaultLaunchDate, _ := time.Parse(model.DateLayout, defaultLaunchDate)
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx     context.Context
		booking *model.BookingDB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
		prep    func()
	}{
		{
			name: "success",
			fields: fields{
				db: setupTestDB(),
			},
			args: args{
				ctx: context.Background(),
				booking: &model.BookingDB{
					FirstName:     "John",
					LastName:      "Doe",
					Gender:        "Male",
					Birthday:      defaultBirthdayDate,
					LaunchpadID:   "1",
					DestinationID: "1",
					LaunchDate:    defaultLaunchDate,
				},
			},
			want:    1,
			wantErr: false,
			prep: func() {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(`INSERT INTO "bookings" (.+)`).WillReturnRows(rows)
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bookingRepository{
				db: tt.fields.db,
			}

			if tt.prep != nil {
				tt.prep()
			}
			got, err := r.CreateBooking(tt.args.ctx, tt.args.booking)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookingRepository.CreateBooking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("bookingRepository.CreateBooking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookingRepository_DeleteBooking(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		prep    func()
	}{
		{
			name: "error",
			fields: fields{
				db: setupTestDB(),
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: true,
			prep: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`DELETE FROM "bookings" WHERE "bookings"."id" = ?`).WillReturnError(gorm.ErrRecordNotFound)
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &bookingRepository{
				db: tt.fields.db,
			}
			if err := r.DeleteBooking(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("bookingRepository.DeleteBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
