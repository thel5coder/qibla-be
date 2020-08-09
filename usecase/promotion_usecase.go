package usecase

import (
	"errors"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type PromotionUseCase struct {
	*UcContract
}

func (uc PromotionUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.PromotionVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	promotions, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, promotion := range promotions {
		res = append(res, viewmodel.PromotionVm{
			ID:                 promotion.ID,
			PromotionPackageID: promotion.PromotionPackageID,
			PackagePromotion:   promotion.PackagePromotion,
			StartDate:          promotion.StartDate,
			EndDate:            promotion.EndDate,
			Platform:           promotion.Platform,
			Position:           promotion.Position,
			Price:              promotion.Price,
			Description:        promotion.Description,
			IsActive:           promotion.IsActive,
			CreatedAt:          promotion.CreatedAt,
			UpdatedAt:          promotion.UpdatedAt,
			DeletedAt:          promotion.DeletedAt.String,
		})
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc PromotionUseCase) ReadBy(column, value string) (res viewmodel.PromotionVm, err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	promotion, err := repository.ReadBy(column, value)

	res = viewmodel.PromotionVm{
		ID:                 promotion.ID,
		PromotionPackageID: promotion.PromotionPackageID,
		PackagePromotion:   promotion.PackagePromotion,
		StartDate:          promotion.StartDate,
		EndDate:            promotion.EndDate,
		Platform:           promotion.Platform,
		Position:           promotion.Position,
		Price:              promotion.Price,
		Description:        promotion.Description,
		IsActive:           promotion.IsActive,
		CreatedAt:          promotion.CreatedAt,
		UpdatedAt:          promotion.UpdatedAt,
		DeletedAt:          promotion.DeletedAt.String,
	}

	return res, err
}

func (uc PromotionUseCase) ReadByPk(ID string) (res viewmodel.PromotionVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc PromotionUseCase) Edit(ID string, input *requests.PromotionRequest) (err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy(ID, "promotion_package_id", input.PromotionPackageID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.PromotionVm{
		ID:                 ID,
		PromotionPackageID: input.PromotionPackageID,
		PackagePromotion:   input.PackagePromotion,
		StartDate:          input.StartDate,
		EndDate:            input.EndDate,
		Platform:           input.Platform,
		Position:           input.Position,
		Price:              input.Price,
		Description:        input.Description,
		IsActive:           input.IsActive,
		UpdatedAt:          now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc PromotionUseCase) Add(input *requests.PromotionRequest) (err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "promotion_package_id", input.PromotionPackageID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.PromotionVm{
		PromotionPackageID: input.PromotionPackageID,
		PackagePromotion:   input.PackagePromotion,
		StartDate:          input.StartDate,
		EndDate:            input.EndDate,
		Platform:           input.Platform,
		Position:           input.Position,
		Price:              input.Price,
		Description:        input.Description,
		IsActive:           input.IsActive,
		CreatedAt:          now,
		UpdatedAt:          now,
	}
	_, err = repository.Add(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc PromotionUseCase) Delete(ID string) (err error){
	repository := actions.NewPromotionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "id", ID)
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

func (uc PromotionUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}
