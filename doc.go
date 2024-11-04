package doc

//go:generate mockgen -source=./internal/repository/bookings.go -destination=./internal/repository/mocks/mock_bookings.go -package=mocks
//go:generate mockgen -source=./internal/service/bookings.go -destination=./internal/service/mocks/mock_bookings.go -package=mocks
