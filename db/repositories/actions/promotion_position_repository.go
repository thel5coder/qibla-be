package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type PromotionPositionRepository struct{
	DB *sql.DB
}

func NewPromotionPositionRepository(DB *sql.DB) contracts.IPromotionPositionRepository{
	return &PromotionPositionRepository{DB: DB}
}

func (repository PromotionPositionRepository) BrowseByPromotionPlatformID(promotionPlatformID string) (data []models.PromotionPosition, err error) {
	statement := `select * from "promotion_positions" where "promotion_platform_id"=$1`
	rows,err := repository.DB.Query(statement,promotionPlatformID)
	if err != nil {
		return data,err
	}

	for rows.Next(){
		dataTemp := models.PromotionPosition{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.PromotionPlatformID,
			&dataTemp.Position,
			)
		if err != nil {
			return data,err
		}

		data = append(data,dataTemp)
	}

	return data,err
}

func (PromotionPositionRepository) Add(promotionPlatformID, position string, tx *sql.Tx) (err error) {
	statement := `insert into "promotion_positions" ("promotion_platform_id","position") values($1,$2)`
	_,err = tx.Exec(statement,promotionPlatformID,position)

	return err
}

func (PromotionPositionRepository) Delete(promotionPlatformID string, tx *sql.Tx) (err error) {
	statement := `delete from "promotion_positions" where "promotion_platform_id"=$1`
	_,err = tx.Exec(statement,promotionPlatformID)

	return err
}

