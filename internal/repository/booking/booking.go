package booking

import (
	"context"

	bookingModel "github.com/davidsugianto/booking-rpc/internal/model/booking"
)

func (r *Repo) GetBookingByID(ctx context.Context, id string) (data *bookingModel.Booking, err error) {
	err = r.db.WithContext(ctx).
		Table(bookingModel.BookingTable).
		Select("*").
		Where("id = ?", id).
		First(&data).Error
	return data, err
}

func (r *Repo) CreateBooking(ctx context.Context, data *bookingModel.Booking) error {
	err := r.db.WithContext(ctx).Table(bookingModel.BookingTable).Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) SearchBooking(ctx context.Context, name string, limit, offset int) (data *bookingModel.Booking, total int64, err error) {
	baseQuery := r.db.WithContext(ctx).Table(bookingModel.BookingTable).Where("status = 1")

	if name != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+name+"%")
	}

	baseQuery = baseQuery.Unscoped()
	countQuery := baseQuery.Count(&total)

	fetchList := baseQuery.
		Select("*").
		Limit(limit).
		Offset(offset).
		Find(&data)

	if fetchList.Error != nil {
		return data, total, fetchList.Error
	}

	if countQuery.Error != nil {
		return data, total, countQuery.Error
	}

	return data, total, nil
}

func (r *Repo) UpdateBooking(ctx context.Context, data *bookingModel.Booking) error {
	err := r.db.WithContext(ctx).
		Table(bookingModel.BookingTable).
		Where("id = ?", data.ID).
		Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) DeleteBooking(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).
		Table(bookingModel.BookingTable).
		Where("id = ?", id).
		Delete(&bookingModel.Booking{}).Error
	if err != nil {
		return err
	}
	return nil
}
