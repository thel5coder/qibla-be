package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type ITestimonialRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.Testimonial, count int, err error)

	ReadBy(column, value string) (data models.Testimonial, err error)

	Edit(input viewmodel.TestimonialVm) (res string, err error)

	Add(input viewmodel.TestimonialVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
