package usecase

import (
	"database/sql"
	"qibla-backend/db/repositories/actions"
)

type PromotionPositionUseCase struct {
	*UcContract
}

func (uc PromotionPositionUseCase) Browse(promotionPlatformID string) (res []string, err error) {
	repository := actions.NewPromotionPositionRepository(uc.DB)
	promotionPositions, err := repository.BrowseByPromotionPlatformID(promotionPlatformID)
	if err != nil {
		return res, err
	}

	for _, promotionPosition := range promotionPositions {
		res = append(res, promotionPosition.Position)
	}

	return res, err
}

func (uc PromotionPositionUseCase) Add(promotionPlatformID, position string, tx *sql.Tx) (err error) {
	repository := actions.NewPromotionPositionRepository(uc.DB)
	err = repository.Add(promotionPlatformID, position, tx)

	return err
}

func (uc PromotionPositionUseCase) Delete(promotionPlatformID string, tx *sql.Tx) (err error) {
	repository := actions.NewPromotionPositionRepository(uc.DB)
	err = repository.Delete(promotionPlatformID, tx)

	return err
}

func (uc PromotionPositionUseCase) Store(promotionPlatformID string, positions []string, tx *sql.Tx) (err error) {
	rows, _ := uc.Browse(promotionPlatformID)

	if len(rows) > 0 {
		err = uc.Delete(promotionPlatformID, tx)
		if err != nil {
			return err
		}
	}

	for _, position := range positions {
		err = uc.Add(promotionPlatformID, position, tx)
		if err != nil {
			return err
		}
	}

	return nil
}
