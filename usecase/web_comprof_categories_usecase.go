package usecase

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type WebComprofCategoryUseCase struct {
	*UcContract
}

func (uc WebComprofCategoryUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.WebComprofCategoryVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewWebComprofCategoryRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	webComprofCategories, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, webComprofCategory := range webComprofCategories {
		res = append(res, viewmodel.WebComprofCategoryVm{
			ID:           webComprofCategory.ID,
			Slug:         webComprofCategory.Slug,
			Name:         webComprofCategory.Name,
			CategoryType: webComprofCategory.CategoryType,
			CreatedAt:    webComprofCategory.CreatedAt,
			UpdatedAt:    webComprofCategory.UpdatedAt,
			DeletedAt:    webComprofCategory.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc WebComprofCategoryUseCase) ReadBy(column, value string) (res viewmodel.WebComprofCategoryVm, err error) {
	repository := actions.NewWebComprofCategoryRepository(uc.DB)
	webComprofCategory, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.WebComprofCategoryVm{
		ID:           webComprofCategory.ID,
		Slug:         webComprofCategory.Slug,
		Name:         webComprofCategory.Name,
		CategoryType: webComprofCategory.CategoryType,
		CreatedAt:    webComprofCategory.CreatedAt,
		UpdatedAt:    webComprofCategory.UpdatedAt,
		DeletedAt:    webComprofCategory.DeletedAt.String,
	}

	return res, err
}

func (uc WebComprofCategoryUseCase) ReadByPk(ID string) (res viewmodel.WebComprofCategoryVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		fmt.Println(err)
		return res, errors.New(messages.DataNotFound)
	}

	return res, err
}

func (uc WebComprofCategoryUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewWebComprofCategoryRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc WebComprofCategoryUseCase) Edit(ID string, input *requests.WebComprofCategoryRequest) (err error) {
	repository := actions.NewWebComprofCategoryRepository(uc.DB)
	now := time.Now().UTC()
	slugName := slug.Make(input.Name)

	isExist, err := uc.IsNameExist(ID, input.Name)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.WebComprofCategoryVm{
		ID:           ID,
		Slug:         slugName,
		Name:         input.Name,
		CategoryType: input.CategoryType,
		UpdatedAt:    now.Format(time.RFC3339),
	}
	_, err = repository.Edit(body)

	return err
}

func (uc WebComprofCategoryUseCase) Add(input *requests.WebComprofCategoryRequest) (error error) {
	repository := actions.NewWebComprofCategoryRepository(uc.DB)
	now := time.Now().UTC()
	slugName := slug.Make(input.Name)

	isExist, err := uc.IsNameExist("", input.Name)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.WebComprofCategoryVm{
		Slug:         slugName,
		Name:         input.Name,
		CategoryType: input.CategoryType,
		CreatedAt:    now.Format(time.RFC3339),
		UpdatedAt:    now.Format(time.RFC3339),
	}
	_, err = repository.Add(body)

	return err
}

func (uc WebComprofCategoryUseCase) Delete(ID string) (err error) {
	repository := actions.NewWebComprofCategoryRepository(uc.DB)
	now := time.Now().UTC()

	count, err := uc.CountByPk(ID)
	if err != nil {
		return errors.New(messages.DataNotFound)
	}

	if count > 0 {
		_, err = repository.Delete(ID, now.Format(time.RFC3339), now.Format(time.RFC3339))
	}

	return err
}

func (uc WebComprofCategoryUseCase) CountByPk(ID string) (res int, err error) {
	repository := actions.NewWebComprofCategoryRepository(uc.DB)
	res, err = repository.CountBy("", "id", ID)

	return res, err
}

func (uc WebComprofCategoryUseCase) IsNameExist(ID, name string) (res bool, err error) {
	count, err := uc.CountBy(ID, "name", name)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
