package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/functioncaller"
	"qibla-backend/helpers/logruslogger"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type PromotionUseCase struct {
	*UcContract
}

func (uc PromotionUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.PromotionVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	promotionPlatformPositionUc := PromotionPlatformUseCase{UcContract: uc.UcContract}
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	promotions, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, promotion := range promotions {
		promotionPlatforms, _ := promotionPlatformPositionUc.Browse(promotion.ID)
		res = append(res, viewmodel.PromotionVm{
			ID:                   promotion.ID,
			PromotionPackageID:   promotion.PromotionPackageID,
			PromotionPackageName: promotion.PackageName,
			PackagePromotion:     promotion.PackagePromotion,
			Positions:            promotionPlatforms,
			StartDate:            promotion.StartDate,
			EndDate:              promotion.EndDate,
			Price:                promotion.Price,
			Description:          promotion.Description,
			IsActive:             promotion.IsActive,
			CreatedAt:            promotion.CreatedAt,
			UpdatedAt:            promotion.UpdatedAt,
			DeletedAt:            promotion.DeletedAt.String,
		})
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc PromotionUseCase) BrowseAll(filters map[string]interface{}) (res []viewmodel.PromotionVm,err error){
	repository := actions.NewPromotionRepository(uc.DB)

	promotions,err := repository.BrowseAll(filters)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-browseAll-promotion")
		return res,err
	}

	for _, promotion := range promotions{
		res = append(res,uc.buildBody(promotion))
	}

	return res,err
}

func (uc PromotionUseCase) ReadBy(column, value string) (res viewmodel.PromotionVm, err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	promotionPlatformPositionUc := PromotionPlatformUseCase{UcContract: uc.UcContract}

	promotion, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}
	promotionPlatforms, _ := promotionPlatformPositionUc.Browse(promotion.ID)
	res = viewmodel.PromotionVm{
		ID:                   promotion.ID,
		PromotionPackageID:   promotion.PromotionPackageID,
		PromotionPackageName: promotion.PackageName,
		PackagePromotion:     promotion.PackagePromotion,
		Positions:            promotionPlatforms,
		StartDate:            promotion.StartDate,
		EndDate:              promotion.EndDate,
		Price:                promotion.Price,
		Description:          promotion.Description,
		IsActive:             promotion.IsActive,
		CreatedAt:            promotion.CreatedAt,
		UpdatedAt:            promotion.UpdatedAt,
		DeletedAt:            promotion.DeletedAt.String,
	}

	return res, err
}

func (uc PromotionUseCase) ReadByPk(ID string) (res viewmodel.PromotionVm, err error) {
	res, err = uc.ReadBy("p.id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc PromotionUseCase) Edit(ID string, input *requests.PromotionRequest) (err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	promotionPlatformUc := PromotionPlatformUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)
	var promotionPlatformBody []viewmodel.PromotionPlatformPositionVm

	//count, err := uc.CountBy("", input.PromotionPackageID, "p.package_promotion", input.PackagePromotion)
	//if err != nil {
	//	return err
	//}
	//
	//if count > 0 {
	//	return errors.New(messages.DataAlreadyExist)
	//}

	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	body := viewmodel.PromotionVm{
		ID:                 ID,
		PromotionPackageID: input.PromotionPackageID,
		PackagePromotion:   input.PackagePromotion,
		StartDate:          input.StartDate,
		EndDate:            input.EndDate,
		Price:              input.Price,
		Description:        input.Description,
		IsActive:           input.IsActive,
		UpdatedAt:          now,
	}
	_, err = repository.Edit(body, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	for _, promotionPlatformPosition := range input.Position {
		promotionPlatformBody = append(promotionPlatformBody, viewmodel.PromotionPlatformPositionVm{
			ID:       ID,
			Platform: promotionPlatformPosition.Platform,
			Position: promotionPlatformPosition.Position,
		})
		err = promotionPlatformUc.Store(ID, promotionPlatformBody, uc.TX)
		if err != nil {
			uc.TX.Rollback()

			return err
		}
	}
	uc.TX.Commit()

	return nil
}

func (uc PromotionUseCase) Add(input *requests.PromotionRequest) (err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	promotionPlatformUc := PromotionPlatformUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)
	var promotionPlatformBody []viewmodel.PromotionPlatformPositionVm

	//count, err := uc.CountBy("", input.PromotionPackageID, "p.package_promotion", input.PackagePromotion)
	//if err != nil {
	//	fmt.Println(1)
	//	return err
	//}
	//
	//if count > 0 {
	//	return errors.New(messages.DataAlreadyExist)
	//}

	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	body := viewmodel.PromotionVm{
		PromotionPackageID: input.PromotionPackageID,
		PackagePromotion:   input.PackagePromotion,
		StartDate:          input.StartDate,
		EndDate:            input.EndDate,
		Price:              input.Price,
		Description:        input.Description,
		IsActive:           true,
		CreatedAt:          now,
		UpdatedAt:          now,
	}
	body.ID, err = repository.Add(body, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	for _, promotionPlatformPosition := range input.Position {
		promotionPlatformBody = append(promotionPlatformBody, viewmodel.PromotionPlatformPositionVm{
			ID:       body.ID,
			Platform: promotionPlatformPosition.Platform,
			Position: promotionPlatformPosition.Position,
		})
	}
	err = promotionPlatformUc.Store(body.ID, promotionPlatformBody, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	uc.TX.Commit()

	return nil
}

func (uc PromotionUseCase) Delete(ID string) (err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	promotionPlatformUc := PromotionPlatformUseCase{UcContract: uc.UcContract}
	promotionPositionUc := PromotionPositionUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "", "p.id", ID)
	if err != nil {
		return err
	}

	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	if count > 0 {
		_, err = repository.Delete(ID, now, now, uc.TX)
		if err != nil {
			uc.TX.Rollback()

			return err
		}

		promotionPlatforms, err := promotionPlatformUc.Browse(ID)
		if err != nil {
			uc.TX.Rollback()

			return err
		}

		for _, promotionPlatform := range promotionPlatforms {
			err = promotionPositionUc.Delete(promotionPlatform.ID, uc.TX)
			if err != nil {
				return err
			}
		}

		err = promotionPlatformUc.Delete(ID, uc.TX)
		if err != nil {
			uc.TX.Rollback()

			return err
		}
	}
	uc.TX.Commit()

	return nil
}

func (uc PromotionUseCase) CountBy(ID, promotionPackageID, column, value string) (res int, err error) {
	repository := actions.NewPromotionRepository(uc.DB)
	res, err = repository.CountBy(ID, promotionPackageID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc PromotionUseCase) buildBody(model models.Promotion) viewmodel.PromotionVm {
	var platformVm []viewmodel.PromotionPlatformVm

	platforms := strings.Split(model.Platform, ",")
	for _, platform := range platforms {
		platformArr := strings.Split(platform, ":")
		platformVm = append(platformVm, viewmodel.PromotionPlatformVm{
			ID:       platformArr[0],
			Platform: platformArr[1],
			Position: nil,
		})
	}

	positions := strings.Split(model.Position, ",")
	for _, position := range positions {
		positionArr := strings.Split(position, ":")
		for i := 0; i < len(platformVm); i++ {
			if positionArr[1] == platformVm[i].ID {
				platformVm[i].Position = append(platformVm[i].Position, viewmodel.PlatformPositionVm{
					ID:       positionArr[0],
					Position: positionArr[2],
				})
			}
		}
	}

	return viewmodel.PromotionVm{
		ID:                   model.ID,
		PromotionPackageID:   model.PromotionPackageID,
		PromotionPackageName: model.PackageName,
		PackagePromotion:     model.PackagePromotion,
		StartDate:            model.StartDate,
		EndDate:              model.EndDate,
		Platform:             platformVm,
		Price:                model.Price,
		Description:          model.Description,
		IsActive:             model.IsActive,
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
		DeletedAt:            model.DeletedAt.String,
	}
}
