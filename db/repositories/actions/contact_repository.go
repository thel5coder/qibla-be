package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/helpers/str"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type ContactRepository struct {
	DB *sql.DB
}

func NewContactRepository(DB *sql.DB) contracts.IContactRepository {
	return &ContactRepository{DB: DB}
}

func (repository ContactRepository) Browse(search, order, sort string, limit, offset int) (data []models.Contact, count int, err error) {
	statement := `select * from "contacts"
                  where (lower("branch_name") like $1 or lower("travel_agent_name") like $1 or "phone_number" like $1 or lower("email") like $1) and "deleted_at" is null 
                 order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.Contact{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.BranchName,
			&dataTemp.TravelAgentName,
			&dataTemp.Address,
			&dataTemp.Longitude,
			&dataTemp.Latitude,
			&dataTemp.AreaCode,
			&dataTemp.PhoneNumber,
			&dataTemp.SKNumber,
			&dataTemp.SKDate,
			&dataTemp.Accreditation,
			&dataTemp.AccreditationDate,
			&dataTemp.DirectorName,
			&dataTemp.DirectorContact,
			&dataTemp.PicName,
			&dataTemp.PicContact,
			&dataTemp.Logo,
			&dataTemp.VirtualAccountNumber,
			&dataTemp.AccountName,
			&dataTemp.AccountNumber,
			&dataTemp.AccountBankName,
			&dataTemp.AccountBankCode,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.Email,
			&dataTemp.IsZakatPartner,
		)
		if err != nil {
			return data, count, err
		}
		data = append(data, dataTemp)
	}

	statement = `select count("id") from "contacts"
                  where (lower("branch_name") like $1 or lower("travel_agent_name") like $1 or "phone_number" like $1 or lower("email") like $1) and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository ContactRepository) BrowseAll(search string,isZakatPartner bool) (data []models.Contact, err error) {
	var rows *sql.Rows
	if search == "" {
		statement := `select * from "contacts" where "is_zakat_partner"=$1`
		rows, err = repository.DB.Query(statement,isZakatPartner)
	} else {
		statement := `select * from "contacts" where (lower("travel_agent_name") like $1 or lower("branch_name") like $1) and "is_zakat_partner"=$2`
		rows, err = repository.DB.Query(statement, "%"+strings.ToLower(search)+"%",isZakatPartner)
	}
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.Contact{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.BranchName,
			&dataTemp.TravelAgentName,
			&dataTemp.Address,
			&dataTemp.Longitude,
			&dataTemp.Latitude,
			&dataTemp.AreaCode,
			&dataTemp.PhoneNumber,
			&dataTemp.SKNumber,
			&dataTemp.SKDate,
			&dataTemp.Accreditation,
			&dataTemp.AccreditationDate,
			&dataTemp.DirectorName,
			&dataTemp.DirectorContact,
			&dataTemp.PicName,
			&dataTemp.PicContact,
			&dataTemp.Logo,
			&dataTemp.VirtualAccountNumber,
			&dataTemp.AccountName,
			&dataTemp.AccountNumber,
			&dataTemp.AccountBankName,
			&dataTemp.AccountBankCode,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.Email,
			&dataTemp.IsZakatPartner,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository ContactRepository) ReadBy(column, value string) (data models.Contact, err error) {
	statement := `select * from "contacts"
                  where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.BranchName,
		&data.TravelAgentName,
		&data.Address,
		&data.Longitude,
		&data.Latitude,
		&data.AreaCode,
		&data.PhoneNumber,
		&data.SKNumber,
		&data.SKDate,
		&data.Accreditation,
		&data.AccreditationDate,
		&data.DirectorName,
		&data.DirectorContact,
		&data.PicName,
		&data.PicContact,
		&data.Logo,
		&data.VirtualAccountNumber,
		&data.AccountName,
		&data.AccountNumber,
		&data.AccountBankName,
		&data.AccountBankCode,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.Email,
		&data.IsZakatPartner,
	)

	return data, err
}

func (repository ContactRepository) Edit(input viewmodel.ContactVm) (res string, err error) {
	statement := `update "contacts" set "branch_name"=$1, "travel_agent_name"=$2, "address"=$3, "longitude"=$4, "latitude"=$5, "area_code"=$6, "phone_number"=$7, "sk_number"=$8, "sk_date"=$9,
                 "accreditation"=$10, "accreditation_date"=$11, "director_name"=$12, "director_contact"=$13, "pic_name"=$14, "pic_contact"=$15, "logo"=$16, "virtual_account_number"=$17, "account_number"=$18,
                 "account_name"=$19, "account_bank_name"=$20, "account_bank_code"=$21, "updated_at"=$22, "email"=$23, "is_zakat_partner"=$24 where "id"=$25 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.BranchName,
		input.TravelAgentName,
		str.EmptyString(input.Address),
		str.EmptyString(input.Longitude),
		str.EmptyString(input.Latitude),
		input.AreaCode,
		input.PhoneNumber,
		str.EmptyString(input.SKNumber),
		datetime.EmptyTime(datetime.StrParseToTime(input.SKDate, "2006-01-02")),
		str.EmptyString(input.Accreditation),
		datetime.EmptyTime(datetime.StrParseToTime(input.AccreditationDate, "2006-01-02")),
		str.EmptyString(input.DirectorName),
		str.EmptyString(input.DirectorContact),
		input.PicName,
		input.PicContact,
		input.FileLogo.ID,
		str.EmptyString(input.VirtualAccountNumber),
		input.AccountNumber,
		input.AccountName,
		input.AccountBankName,
		input.AccountBankCode,
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.Email,
		input.IsZakatPartner,
		input.ID,
	).Scan(&res)

	return res, err
}

func (repository ContactRepository) Add(input viewmodel.ContactVm) (res string, err error) {
	statement := `insert into "contacts" ("email","branch_name","travel_agent_name","address","longitude","latitude","area_code","phone_number","sk_number","sk_date","accreditation","accreditation_date","director_name",
                  "director_contact","pic_name","pic_contact","logo","virtual_account_number","account_number","account_name","account_bank_name","account_bank_code","is_zakat_partner","created_at","updated_at") 
                  values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25) returning "id"`

	err = repository.DB.QueryRow(
		statement,
		input.Email,
		input.BranchName,
		input.TravelAgentName,
		str.EmptyString(input.Address),
		str.EmptyString(input.Longitude),
		str.EmptyString(input.Latitude),
		input.AreaCode,
		input.PhoneNumber,
		str.EmptyString(input.SKNumber),
		datetime.EmptyTime(datetime.StrParseToTime(input.SKDate, "2006-01-02")),
		str.EmptyString(input.Accreditation),
		datetime.EmptyTime(datetime.StrParseToTime(input.AccreditationDate, "2006-01-02")),
		str.EmptyString(input.DirectorName),
		str.EmptyString(input.DirectorContact),
		input.PicName,
		input.PicContact,
		input.FileLogo.ID,
		input.VirtualAccountNumber,
		input.AccountNumber,
		input.AccountName,
		input.AccountBankName,
		input.AccountBankCode,
		input.IsZakatPartner,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (repository ContactRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "contacts" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository ContactRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "contacts"
			where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "contacts"
			where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
