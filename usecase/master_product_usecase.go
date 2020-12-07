package usecase

import (
	"errors"
	"github.com/gosimple/slug"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type MasterProductUseCase struct {
	*UcContract
}

func (uc MasterProductUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.MasterProductVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewMasterProductRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	masterProducts, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, masterProduct := range masterProducts {
		res = append(res, viewmodel.MasterProductVm{
			ID:        masterProduct.ID,
			Slug:      masterProduct.Slug,
			Name:      masterProduct.Name,
			SubscriptionType: masterProduct.SubscriptionType,
			CreatedAt: masterProduct.CreatedAt,
			UpdatedAt: masterProduct.UpdatedAt,
			DeletedAt: masterProduct.DeletedAt.String,
		})
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc MasterProductUseCase) BrowseAll() (res []viewmodel.MasterProductVm,err error){
	repository := actions.NewMasterProductRepository(uc.DB)
	masterProducts, err := repository.BrowseAll()
	if err != nil {
		return res, err
	}

	for _, masterProduct := range masterProducts {
		res = append(res, viewmodel.MasterProductVm{
			ID:        masterProduct.ID,
			Slug:      masterProduct.Slug,
			Name:      masterProduct.Name,
			SubscriptionType: masterProduct.SubscriptionType,
			CreatedAt: masterProduct.CreatedAt,
			UpdatedAt: masterProduct.UpdatedAt,
			DeletedAt: masterProduct.DeletedAt.String,
		})
	}

	return res,err
}

func (uc MasterProductUseCase) BrowseExtraProducts() (res []viewmodel.MasterProductVm,err error){
	repository := actions.NewMasterProductRepository(uc.DB)
	masterProducts, err := repository.BrowseExtraProducts()
	if err != nil {
		return res, err
	}

	for _, masterProduct := range masterProducts {
		res = append(res, viewmodel.MasterProductVm{
			ID:        masterProduct.ID,
			Slug:      masterProduct.Slug,
			Name:      masterProduct.Name,
			SubscriptionType: masterProduct.SubscriptionType,
			CreatedAt: masterProduct.CreatedAt,
			UpdatedAt: masterProduct.UpdatedAt,
			DeletedAt: masterProduct.DeletedAt.String,
		})
	}

	return res,err
}

func (uc MasterProductUseCase) ReadBy(column, value string) (res viewmodel.MasterProductVm, err error) {
	repository := actions.NewMasterProductRepository(uc.DB)
	masterProduct, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.MasterProductVm{
		ID:               masterProduct.ID,
		Slug:             masterProduct.Slug,
		Name:             masterProduct.Name,
		SubscriptionType: masterProduct.SubscriptionType,
		CreatedAt:        masterProduct.CreatedAt,
		UpdatedAt:        masterProduct.UpdatedAt,
		DeletedAt:        masterProduct.DeletedAt.String,
	}

	return res, err
}

func (uc MasterProductUseCase) ReadByPk(ID string) (res viewmodel.MasterProductVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc MasterProductUseCase) Edit(ID string, input *requests.MasterProductRequest) (err error) {
	repository := actions.NewMasterProductRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy(ID, "slug", slug.Make(input.Name))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.MasterProductVm{
		ID:               ID,
		Slug:             slug.Make(input.Name),
		Name:             input.Name,
		SubscriptionType: input.SubscriptionType,
		UpdatedAt:        now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc MasterProductUseCase) Add(input *requests.MasterProductRequest) (err error) {
	repository := actions.NewMasterProductRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "slug", slug.Make(input.Name))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.MasterProductVm{
		Slug:             slug.Make(input.Name),
		Name:             input.Name,
		SubscriptionType: input.SubscriptionType,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	_, err = repository.Add(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc MasterProductUseCase) Delete(ID string) (err error) {
	repository := actions.NewMasterProductRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "id", ID)
	if err != nil {
		return err
	}
	if count > 0 {
		_, err = repository.Delete(ID, now, now)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc MasterProductUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewMasterProductRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}
