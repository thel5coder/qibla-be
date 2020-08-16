package usecase

import (
	"errors"
	"github.com/gosimple/slug"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type PromotionPackageUseCase struct {
	*UcContract
}

func (uc PromotionPackageUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.PromotionPackageVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewPromotionPackageRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	promotionPackages, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, promotionPackage := range promotionPackages {
		res = append(res, viewmodel.PromotionPackageVm{
			ID:          promotionPackage.ID,
			Slug:        promotionPackage.Slug,
			PackageName: promotionPackage.PackageName,
			IsActive:    promotionPackage.IsActive,
			CreatedAt:   promotionPackage.CreatedAt,
			UpdatedAt:   promotionPackage.UpdatedAt,
			DeletedAt:   promotionPackage.DeletedAt.String,
		})
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc PromotionPackageUseCase) BrowseAll() (res []viewmodel.PromotionPackageVm,err error){
	repository := actions.NewPromotionPackageRepository(uc.DB)
	promotionPackages,err := repository.BrowseAll()
	if err != nil {
		return res,err
	}

	for _, promotionPackage := range promotionPackages {
		res = append(res, viewmodel.PromotionPackageVm{
			ID:          promotionPackage.ID,
			Slug:        promotionPackage.Slug,
			PackageName: promotionPackage.PackageName,
			IsActive:    promotionPackage.IsActive,
			CreatedAt:   promotionPackage.CreatedAt,
			UpdatedAt:   promotionPackage.UpdatedAt,
			DeletedAt:   promotionPackage.DeletedAt.String,
		})
	}

	return res,err
}

func (uc PromotionPackageUseCase) ReadBy(column, value string) (res viewmodel.PromotionPackageVm, err error) {
	repository := actions.NewPromotionPackageRepository(uc.DB)
	promotionPackage, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.PromotionPackageVm{
		ID:          promotionPackage.ID,
		Slug:        promotionPackage.Slug,
		PackageName: promotionPackage.PackageName,
		IsActive:    promotionPackage.IsActive,
		CreatedAt:   promotionPackage.CreatedAt,
		UpdatedAt:   promotionPackage.UpdatedAt,
		DeletedAt:   promotionPackage.DeletedAt.String,
	}

	return res, err
}

func (uc PromotionPackageUseCase) ReadByPk(ID string) (res viewmodel.PromotionPackageVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc PromotionPackageUseCase) Edit(ID string, input *requests.PromotionPackageRequest) (err error) {
	repository := actions.NewPromotionPackageRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy(ID, "slug", slug.Make(input.PackageName))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.PromotionPackageVm{
		ID:          ID,
		Slug:        slug.Make(input.PackageName),
		PackageName: input.PackageName,
		IsActive:    input.IsActive,
		UpdatedAt:   now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc PromotionPackageUseCase) Add(input *requests.PromotionPackageRequest) (err error) {
	repository := actions.NewPromotionPackageRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "slug", slug.Make(input.PackageName))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.PromotionPackageVm{
		Slug:        slug.Make(input.PackageName),
		PackageName: input.PackageName,
		IsActive:    true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	_, err = repository.Add(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc PromotionPackageUseCase) Delete(ID string) (err error) {
	repository := actions.NewPromotionPackageRepository(uc.DB)
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

func (uc PromotionPackageUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewPromotionPackageRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}
