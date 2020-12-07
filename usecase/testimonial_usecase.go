package usecase

import (
	"errors"
	"fmt"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type TestimonialUseCase struct {
	*UcContract
}

func (uc TestimonialUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.TestimonialVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewTestimonialRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	testimonials, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, testimonial := range testimonials {
		res = append(res, viewmodel.TestimonialVm{
			ID:                   testimonial.ID,
			WebContentCategoryID: testimonial.WebContentCategoryID,
			FileID:               testimonial.FileID,
			Path:                 testimonial.Path.String,
			CustomerName:         testimonial.CustomerName,
			JobPosition:          testimonial.JobPosition,
			Testimony:            testimonial.Testimony,
			Rating:               testimonial.Rating,
			IsActive:             testimonial.IsActive,
			CreatedAt:            testimonial.CreatedAt,
			UpdatedAt:            testimonial.UpdatedAt,
			DeletedAt:            testimonial.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc TestimonialUseCase) ReadBy(column, value string) (res viewmodel.TestimonialVm, err error) {
	repository := actions.NewTestimonialRepository(uc.DB)
	testimonial, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.TestimonialVm{
		ID:                   testimonial.ID,
		WebContentCategoryID: testimonial.WebContentCategoryID,
		FileID:               testimonial.FileID,
		Path:                 testimonial.Path.String,
		CustomerName:         testimonial.CustomerName,
		JobPosition:          testimonial.JobPosition,
		Testimony:            testimonial.Testimony,
		Rating:               testimonial.Rating,
		IsActive:             testimonial.IsActive,
		CreatedAt:            testimonial.CreatedAt,
		UpdatedAt:            testimonial.UpdatedAt,
		DeletedAt:            testimonial.DeletedAt.String,
	}

	return res, err
}

func (uc TestimonialUseCase) ReadByPk(ID string) (res viewmodel.TestimonialVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		fmt.Println(err)
		return res, errors.New(messages.DataNotFound)
	}

	return res, err
}

func (uc TestimonialUseCase) Edit(ID string, input *requests.TestimonialRequest) (err error){
	repository := actions.NewTestimonialRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	isExist,err := uc.isExist(ID,input.CustomerName)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.TestimonialVm{
		ID:                   ID,
		FileID:               input.FileID,
		CustomerName:         input.CustomerName,
		JobPosition:          input.JobPosition,
		Testimony:            input.Testimony,
		Rating:               input.Rating,
		IsActive:             input.IsActive,
		UpdatedAt:            now,
	}
	_,err = repository.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc TestimonialUseCase) Add(input *requests.TestimonialRequest) (err error){
	repository := actions.NewTestimonialRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	isExist,err := uc.isExist("",input.CustomerName)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.TestimonialVm{
		WebContentCategoryID: input.WebContentCategoryID,
		FileID:               input.FileID,
		CustomerName:         input.CustomerName,
		JobPosition:          input.JobPosition,
		Testimony:            input.Testimony,
		Rating:               input.Rating,
		IsActive:             true,
		CreatedAt:            now,
		UpdatedAt:            now,
	}
	_, err = repository.Add(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc TestimonialUseCase) Delete(ID string) (err error){
	repository := actions.NewTestimonialRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count,err := uc.countBy("","id",ID)
	if err != nil {
		return err
	}

	if count > 0 {
		_,err = repository.Delete(ID,now,now)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc TestimonialUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewTestimonialRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc TestimonialUseCase) isExist(ID, customerName string) (res bool, err error) {
	count, err := uc.countBy(ID, "customer_name", customerName)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
