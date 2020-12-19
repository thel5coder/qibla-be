package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type CrmBoardRepository struct {
	DB *sql.DB
}

func NewCrmBoardRepository(db *sql.DB) contracts.ICrmBoardRepository {
	return &CrmBoardRepository{DB: db}
}

func (repository CrmBoardRepository) BrowseByCrmStoryID(crmStoryID string) (data []models.CrmBoard, err error) {
	statement := `select * from "crm_boards" where "crm_story_id"=$1 and "deleted_at" is null order by "updated_at" asc`
	rows, err := repository.DB.Query(statement, crmStoryID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.CrmBoard{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.CrmStoryID,
			&dataTemp.ContactID,
			&dataTemp.Opportunity,
			&dataTemp.ProfitExpectation,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.Star,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository CrmBoardRepository) ReadBy(column, value string) (data models.CrmBoard, err error) {
	statement := `select * from "crm_boards" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.CrmStoryID,
		&data.ContactID,
		&data.Opportunity,
		&data.ProfitExpectation,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.Star,
	)

	return data, err
}

func (repository CrmBoardRepository) Edit(input viewmodel.CrmBoardVm) (res string, err error) {
	statement := `update "crm_boards" set "crm_story_id"=$1, "contact_id"=$2, "opportunity"=$3, "profit_expectation"=$4,"star"=$5, "updated_at"=$6 where "id"=$7 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.CrmStoryID,
		input.ContactID,
		input.Opportunity,
		input.ProfitExpectation,
		input.Star,
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ID,
	).Scan(&res)

	return res, err
}

func (repository CrmBoardRepository) EditBoardStory(ID, crmStoryID, updatedAt string) (res string, err error) {
	statement := `update "crm_boards" set "crm_story_id"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, crmStoryID, datetime.StrParseToTime(updatedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository CrmBoardRepository) Add(input viewmodel.CrmBoardVm) (res string, err error) {
	statement := `insert into "crm_boards" ("crm_story_id","contact_id","opportunity","profit_expectation","star","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7) returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.CrmStoryID,
		input.ContactID,
		input.Opportunity,
		input.ProfitExpectation,
		input.Star,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (repository CrmBoardRepository) DeleteBy(column, value, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "crm_boards" set "updated_at"=$1, "deleted_at"=$2 where ` + column + `=$3 returning "crm_story_id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), value).Scan(&res)

	return res, err
}

func (repository CrmBoardRepository) CountBy(ID, crmStoryID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") where ` + column + `=$1 where "deleted_at" is null and "crm_story_id"=$2`
		err = repository.DB.QueryRow(statement, value, crmStoryID).Scan(&res)
	} else {
		statement := `select count("id") where (` + column + `=$1 where "deleted_at" is null and "crm_story_id"=$2) and "id"<>$3`
		err = repository.DB.QueryRow(statement, value, crmStoryID, ID).Scan(&res)
	}

	return res, err
}
