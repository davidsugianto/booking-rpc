package booking

import (
	"context"

	bookingModel "github.com/davidsugianto/booking-rpc/internal/model/booking"
)

type bookingRepo interface {
	GetBookingByID(ctx context.Context, id string) (data *bookingModel.Booking, err error)
	CreateBooking(ctx context.Context, data *bookingModel.Booking) error
	UpdateBooking(ctx context.Context, data *bookingModel.Booking) error
	DeleteBooking(ctx context.Context, id string) error
	SearchBooking(ctx context.Context, name string, limit, offset int) (data *bookingModel.Booking, total int64, err error)
}

type UseCase struct {
	bookingRepo bookingRepo
}

func New(bookingRepo bookingRepo) *UseCase {
	return &UseCase{
		bookingRepo: bookingRepo,
	}
}
