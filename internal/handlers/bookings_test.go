package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/sunil-dev608/space-trouble/config"
	"github.com/sunil-dev608/space-trouble/internal/pkg/logger"
	"github.com/sunil-dev608/space-trouble/internal/pkg/model"
	"github.com/sunil-dev608/space-trouble/internal/service"
	svcmocks "github.com/sunil-dev608/space-trouble/internal/service/mocks"
)

func TestBookingHandler_CreateBooking(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	input := &model.Booking{
		FirstName:     "abc",
		LastName:      "def",
		Gender:        "F",
		Birthday:      "2000-01-01",
		LaunchpadID:   "ksc_lc_39a",
		DestinationID: "Mars",
		LaunchDate:    "2022-01-01",
	}
	inputBytes, _ := json.Marshal(input)
	inputDB, _ := input.ToDB()
	type fields struct {
		bookingService service.BookingService
		cfg            *config.Config
		logger         logger.Logger
	}
	type args struct {
		c   echo.Context
		req *http.Request
		rec *httptest.ResponseRecorder
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
				bookingService: svcmocks.NewMockBookingService(ctrl),
				cfg: &config.Config{
					CompetitorLaunchesAPIURL:   "https://api.spacexdata.com/v5/launches/query",
					CompetitorLaunchpadsAPIURL: "https://api.spacexdata.com/v5/launchpads/query",
					Destinations: map[string]interface{}{
						"Mars": nil,
						"Moon": nil,
					},
					Launchpads:    map[string]string{"ksc_lc_39a": "active"},
					ServerAddress: ":8080",
					DBDsn:         "postgres://space_trouble:space_trouble@localhost:5432/space_trouble",
				},
				logger: nil,
			},
			args: args{
				c:   nil,
				req: httptest.NewRequest(http.MethodPost, "/api/v1/bookings", bytes.NewBufferString(string(inputBytes))),
				rec: httptest.NewRecorder(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &BookingHandler{
				bookingService: tt.fields.bookingService,
				cfg:            tt.fields.cfg,
				logger:         tt.fields.logger,
			}
			tt.args.req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			tt.args.c = echo.New().NewContext(tt.args.req, tt.args.rec)

			h.bookingService.(*svcmocks.MockBookingService).EXPECT().CreateBooking(tt.args.c.Request().Context(), inputDB).Return(int64(1), nil)
			if err := h.CreateBooking(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("BookingHandler.CreateBooking() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.rec.Code != http.StatusCreated {
				t.Errorf("Expected status 201, got %v", tt.args.rec.Code)
			}
		})
	}
}
