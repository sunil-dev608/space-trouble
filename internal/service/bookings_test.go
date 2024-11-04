package service

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sunil-dev608/space-trouble/internal/competitors"
	"github.com/sunil-dev608/space-trouble/internal/pkg/model"
	"github.com/sunil-dev608/space-trouble/internal/repository"
	repomocks "github.com/sunil-dev608/space-trouble/internal/repository/mocks"
)

var (
	defaultBirthday   = "2000-01-01"
	defaultLaunchDate = "2022-01-01"
)

func Test_bookingService_CreateBooking(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	defaultBirthdayDate, _ := time.Parse(model.DateLayout, defaultBirthday)
	defaultLaunchDate, _ := time.Parse(model.DateLayout, defaultLaunchDate)
	type fields struct {
		competitorLaunchesProvier competitors.CompetitorLaunchesProvier
		bookingRepo               repository.BookingRepository
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
	}{
		{
			name: "success",
			fields: fields{
				competitorLaunchesProvier: competitors.NewCompetitorLaunchesProvier("https://api.spacexdata.com/v5/launches/query"),
				bookingRepo:               repomocks.NewMockBookingRepository(ctrl),
			},
			args: args{
				ctx: context.Background(),
				booking: &model.BookingDB{
					FirstName:     "John",
					LastName:      "Doe",
					Gender:        "Male",
					Birthday:      defaultBirthdayDate,
					LaunchpadID:   "5eb87d5e4585a20024765ebc",
					DestinationID: "Moon",
					LaunchDate:    defaultLaunchDate,
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookingService{
				competitorLaunchesProvier: tt.fields.competitorLaunchesProvier,
				bookingRepo:               tt.fields.bookingRepo,
			}

			s.bookingRepo.(*repomocks.MockBookingRepository).EXPECT().HasConflictingFlight(tt.args.ctx, tt.args.booking).Return(false, nil)
			s.bookingRepo.(*repomocks.MockBookingRepository).EXPECT().CreateBooking(tt.args.ctx, tt.args.booking).Return(int64(1), nil)
			got, err := s.CreateBooking(tt.args.ctx, tt.args.booking)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookingService.CreateBooking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("bookingService.CreateBooking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookingService_DeleteBooking(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		competitorLaunchesProvier competitors.CompetitorLaunchesProvier
		bookingRepo               repository.BookingRepository
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
	}{
		{
			name: "success",
			fields: fields{
				competitorLaunchesProvier: competitors.NewCompetitorLaunchesProvier("https://api.spacexdata.com/v5/launches/query"),
				bookingRepo:               repomocks.NewMockBookingRepository(ctrl),
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &bookingService{
				competitorLaunchesProvier: tt.fields.competitorLaunchesProvier,
				bookingRepo:               tt.fields.bookingRepo,
			}
			s.bookingRepo.(*repomocks.MockBookingRepository).EXPECT().DeleteBooking(tt.args.ctx, tt.args.id).Return(nil)
			if err := s.DeleteBooking(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("bookingService.DeleteBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
