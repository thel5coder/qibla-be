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

type MasterZakatUseCase struct {
	*UcContract
}

func (uc MasterZakatUseCase) Browse(search map[string]interface{}, order, sort string, page, limit int) (res []viewmodel.MasterZakatVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	masterZakats, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, masterZakat := range masterZakats {
		res = append(res, viewmodel.MasterZakatVm{
			ID:               masterZakat.ID,
			Slug:             masterZakat.Slug,
			TypeZakat:        masterZakat.TypeZakat,
			Name:             masterZakat.Name,
			Description:      masterZakat.Description,
			Amount:           masterZakat.Amount.Int32,
			CurrentGoldPrice: masterZakat.CurrentGoldPrice.Int32,
			GoldNishab:       masterZakat.GoldNishab.Int32,
			CreatedAt:        masterZakat.CreatedAt,
			UpdatedAt:        masterZakat.UpdatedAt,
			DeletedAt:        masterZakat.DeletedAt.String,
		})
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc MasterZakatUseCase) BrowseAll() (res []viewmodel.MasterZakatVm, err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	masterProducts, err := repository.BrowseAll()
	if err != nil {
		return res, err
	}

	for _, masterZakat := range masterProducts {
		res = append(res, viewmodel.MasterZakatVm{
			ID:               masterZakat.ID,
			Slug:             masterZakat.Slug,
			TypeZakat:        masterZakat.TypeZakat,
			Name:             masterZakat.Name,
			Description:      masterZakat.Description,
			Amount:           masterZakat.Amount.Int32,
			CurrentGoldPrice: masterZakat.CurrentGoldPrice.Int32,
			GoldNishab:       masterZakat.GoldNishab.Int32,
			CreatedAt:        masterZakat.CreatedAt,
			UpdatedAt:        masterZakat.UpdatedAt,
			DeletedAt:        masterZakat.DeletedAt.String,
		})
	}

	return res, err
}

func (uc MasterZakatUseCase) ReadBy(column, value string) (res viewmodel.MasterZakatVm, err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	masterZakat, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.MasterZakatVm{
		ID:               masterZakat.ID,
		Slug:             masterZakat.Slug,
		TypeZakat:        masterZakat.TypeZakat,
		Name:             masterZakat.Name,
		Description:      masterZakat.Description,
		Amount:           masterZakat.Amount.Int32,
		CurrentGoldPrice: masterZakat.CurrentGoldPrice.Int32,
		GoldNishab:       masterZakat.GoldNishab.Int32,
		CreatedAt:        masterZakat.CreatedAt,
		UpdatedAt:        masterZakat.UpdatedAt,
		DeletedAt:        masterZakat.DeletedAt.String,
	}

	return res, err
}

func (uc MasterZakatUseCase) ReadByPk(ID string) (res viewmodel.MasterZakatVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc MasterZakatUseCase) Edit(ID string, input *requests.MasterZakatRequest) (err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy(ID, "slug", slug.Make(input.Name))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.MasterZakatVm{
		ID:               ID,
		Slug:             slug.Make(input.Name),
		TypeZakat:        input.TypeZakat,
		Name:             input.Name,
		Description:      input.Description,
		Amount:           input.Amount,
		CurrentGoldPrice: input.CurrentGoldPrice,
		GoldNishab:       input.GoldNishab,
		UpdatedAt:        now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc MasterZakatUseCase) Add(input *requests.MasterZakatRequest) (err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "slug", slug.Make(input.Name))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.MasterZakatVm{
		Slug:             slug.Make(input.Name),
		TypeZakat:        input.TypeZakat,
		Name:             input.Name,
		Description:      input.Description,
		Amount:           input.Amount,
		CurrentGoldPrice: input.CurrentGoldPrice,
		GoldNishab:       input.GoldNishab,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	_, err = repository.Add(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc MasterZakatUseCase) Delete(ID string) (err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
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

func (uc MasterZakatUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}
