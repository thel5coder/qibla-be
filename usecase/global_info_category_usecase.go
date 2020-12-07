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

type GlobalInfoCategoryUseCase struct {
	*UcContract
}

func (uc GlobalInfoCategoryUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.GlobalInfoCategoryVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewGlobalInfoCategoryRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	globalInfoCategories, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, globalInfoCategory := range globalInfoCategories {
		res = append(res, viewmodel.GlobalInfoCategoryVm{
			ID:        globalInfoCategory.ID,
			Name:      globalInfoCategory.Name,
			Slug:      globalInfoCategory.Slug,
			CreatedAt: globalInfoCategory.CreatedAt,
			UpdatedAt: globalInfoCategory.UpdatedAt,
			DeletedAt: globalInfoCategory.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc GlobalInfoCategoryUseCase) ReadBy(column, value string) (res viewmodel.GlobalInfoCategoryVm, err error) {
	repository := actions.NewGlobalInfoCategoryRepository(uc.DB)
	globalInfoCategory, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.GlobalInfoCategoryVm{
		ID:        globalInfoCategory.ID,
		Name:      globalInfoCategory.Name,
		Slug:      globalInfoCategory.Slug,
		CreatedAt: globalInfoCategory.CreatedAt,
		UpdatedAt: globalInfoCategory.UpdatedAt,
		DeletedAt: globalInfoCategory.DeletedAt.String,
	}

	return res, err
}

func (uc GlobalInfoCategoryUseCase) ReadByPk(ID string) (res viewmodel.GlobalInfoCategoryVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		fmt.Println(err)
		return res, errors.New(messages.DataNotFound)
	}

	return res, err
}

func (uc GlobalInfoCategoryUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewGlobalInfoCategoryRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc GlobalInfoCategoryUseCase) Edit(ID string, input *requests.GlobalInfoCategoryRequest) (err error) {
	repository := actions.NewGlobalInfoCategoryRepository(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsNameExist(ID, input.Name)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.GlobalInfoCategoryVm{
		ID:        ID,
		Name:      input.Name,
		Slug:      slug.Make(input.Name),
		UpdatedAt: now.Format(time.RFC3339),
	}
	_, err = repository.Edit(body)

	return err
}

func (uc GlobalInfoCategoryUseCase) Add(input *requests.GlobalInfoCategoryRequest) (error error) {
	repository := actions.NewGlobalInfoCategoryRepository(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsNameExist("", input.Name)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.GlobalInfoCategoryVm{
		Name:      input.Name,
		Slug:      slug.Make(input.Name),
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}
	_, err = repository.Add(body)

	return err
}

func (uc GlobalInfoCategoryUseCase) Delete(ID string) (err error) {
	repository := actions.NewGlobalInfoCategoryRepository(uc.DB)
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

func (uc GlobalInfoCategoryUseCase) CountByPk(ID string) (res int, err error) {
	repository := actions.NewGlobalInfoCategoryRepository(uc.DB)
	res, err = repository.CountBy("", "id", ID)

	return res, err
}

func (uc GlobalInfoCategoryUseCase) IsNameExist(ID, name string) (res bool, err error) {
	count, err := uc.CountBy(ID, "slug", name)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
