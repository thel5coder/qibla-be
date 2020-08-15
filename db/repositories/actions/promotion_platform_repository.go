package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type PromotionPlatformRepository struct {
	DB *sql.DB
}

func NewPromotionPlatformRepository(DB *sql.DB) contracts.IPromotionPlatformRepository {
	return &PromotionPlatformRepository{DB: DB}
}

func (repository PromotionPlatformRepository) BrowseByPromotionID(promotionID string) (data []models.PromotionPlatform, err error) {
	statement := `select * from "promotion_platforms" where "promotion_id"=$1`
	rows, err := repository.DB.Query(statement, promotionID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.PromotionPlatform{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.PromotionID,
			&dataTemp.Platform,
		)
		data = append(data, dataTemp)
	}

	return data, err
}

func (PromotionPlatformRepository) Add(promotionID, platform string, tx *sql.Tx) (res string, err error) {
	statement := `insert into "promotion_platforms" ("promotion_id","platform") values($1,$2) returning "id"`
	err = tx.QueryRow(statement, promotionID, platform).Scan(&res)

	return res, err
}

func (PromotionPlatformRepository) Delete(promotionID string, tx *sql.Tx) (err error) {
	statement := `delete from "promotion_platforms" where "promotion_id"=$1`
	_, err = tx.Exec(statement, promotionID)

	return err
}
