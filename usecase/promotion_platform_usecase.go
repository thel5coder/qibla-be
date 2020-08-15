package usecase

import (
	"database/sql"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
)

type PromotionPlatformUseCase struct {
	*UcContract
}

func (uc PromotionPlatformUseCase) Browse(promotionID string) (res []viewmodel.PromotionPlatformPositionVm, err error) {
	repository := actions.NewPromotionPlatformRepository(uc.DB)
	promotionPositionUc := PromotionPositionUseCase{UcContract: uc.UcContract}

	promotionPlatforms, err := repository.BrowseByPromotionID(promotionID)
	if err != nil {
		return res, err
	}

	for _, promotionPlatform := range promotionPlatforms {
		promotionPositions, _ := promotionPositionUc.Browse(promotionPlatform.ID)
		res = append(res, viewmodel.PromotionPlatformPositionVm{
			ID:       promotionPlatform.ID,
			Platform: promotionPlatform.Platform,
			Position: promotionPositions,
		})
	}

	return res, err
}

func (uc PromotionPlatformUseCase) Add(promotionID, platform string, tx *sql.Tx) (res string, err error) {
	repository := actions.NewPromotionPlatformRepository(uc.DB)
	res, err = repository.Add(promotionID, platform, tx)

	return res, err
}

func (uc PromotionPlatformUseCase) Delete(promotionID string, tx *sql.Tx) (err error) {
	repository := actions.NewPromotionPlatformRepository(uc.DB)
	err = repository.Delete(promotionID, tx)

	return err
}

func (uc PromotionPlatformUseCase) Store(promotionID string, promotionPositionPlatforms []viewmodel.PromotionPlatformPositionVm, tx *sql.Tx) (err error) {
	promotionPositionUc := PromotionPositionUseCase{UcContract: uc.UcContract}
	rows, _ := uc.Browse(promotionID)

	if len(rows) > 0 {
		err = uc.Delete(promotionID, tx)
		if err != nil {
			return err
		}
	}

	for _, promotionPositionPlatform := range promotionPositionPlatforms {
		promotionPlatformID, err := uc.Add(promotionID, promotionPositionPlatform.Platform, tx)
		if err != nil {
			return err
		}

		err = promotionPositionUc.Store(promotionPlatformID, promotionPositionPlatform.Position, tx)
		if err != nil {
			return err
		}
	}

	return nil
}
