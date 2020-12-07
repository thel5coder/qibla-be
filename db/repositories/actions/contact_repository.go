package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/enums"
	"qibla-backend/pkg/str"
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

const (
	contactSelectStatement = `select c.id,c.branch_name,c.travel_agent_name,c.address,c.longitude,c.latitude,c.area_code,c.phone_number,c.sk_number,c.sk_date,accreditation,
                              c.accreditation_date,c.director_name,c.director_contact,c.pic_name,c.pic_contact,c.logo,f.name,f.path,c.virtual_account_number,c.account_number,c.account_name,
                              c.account_bank_name,c.account_bank_code,c.email,c.is_zakat_partner,c.created_at,c.updated_at`
	contactJoinStatement = `left join "files" f on f.id = c.logo`
)

func (repository ContactRepository) scanRows(rows *sql.Rows) (res models.Contact, err error) {
	err = rows.Scan(&res.ID, &res.BranchName, &res.TravelAgentName, &res.Address, &res.Longitude, &res.Latitude, &res.AreaCode, &res.PhoneNumber, &res.SKNumber, &res.SKDate, &res.Accreditation,
		&res.AccreditationDate, &res.DirectorName, &res.DirectorContact, &res.PicName, &res.PicContact, &res.Logo, &res.LogoName, &res.LogoPath, &res.VirtualAccountNumber, &res.AccountNumber, &res.AccountName, &res.AccountBankName,
		&res.AccountBankCode, &res.Email, &res.IsZakatPartner, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ContactRepository) scanRow(row *sql.Row) (res models.Contact, err error) {
	err = row.Scan(&res.ID, &res.BranchName, &res.TravelAgentName, &res.Address, &res.Longitude, &res.Latitude, &res.AreaCode, &res.PhoneNumber, &res.SKNumber, &res.SKDate, &res.Accreditation,
		&res.AccreditationDate, &res.DirectorName, &res.DirectorContact, &res.PicName, &res.PicContact, &res.Logo, &res.LogoName, &res.LogoPath, &res.VirtualAccountNumber, &res.AccountNumber, &res.AccountName, &res.AccountBankName,
		&res.AccountBankCode, &res.Email, &res.IsZakatPartner, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ContactRepository) Browse(filters map[string]interface{}, order, sort string, limit, offset int) (data []models.Contact, count int, err error) {
	filterStatement := ``

	if val, ok := filters["branch_name"]; ok {
		filterStatement += ` and lower(c.branch_name) like '%` + val.(string) + `%'`
	}

	if val, ok := filters["travel_agent_name"]; ok {
		filterStatement += ` and lower(c.travel_agent_name) like '%` + val.(string) + `%'`
	}

	if val, ok := filters["address"]; ok {
		filterStatement += ` and lower(c.address) like '%` + val.(string) + `%'`
	}

	if val, ok := filters["phone_number"]; ok {
		filterStatement += ` and c.phone_number like '%` + val.(string) + `%'`
	}

	if val, ok := filters["email"]; ok {
		filterStatement += ` and (lower(c.email) like '%` + val.(string) + `%'`
	}

	statement := contactSelectStatement + ` from "contacts" c ` + contactJoinStatement + ` where c.deleted_at is null` + filterStatement + ` order by c.` + order + ` ` + sort + ` limit $1 offset $2`
	rows, err := repository.DB.Query(statement, limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, count, err
		}
		data = append(data, temp)
	}

	statement = `select count(c."id") from "contacts" c ` + contactJoinStatement + ` where c."deleted_at" is null` + filterStatement
	err = repository.DB.QueryRow(statement).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository ContactRepository) BrowseAll(search string, isZakatPartner bool) (data []models.Contact, err error) {
	var rows *sql.Rows
	whereStatement := `where c."is_zakat_partner"=$1 and c."deleted_at" is null`
	whereParams := []interface{}{isZakatPartner}
	if search != "" {
		whereStatement += `and (lower(c."travel_agent_name") like $2 or lower(c."branch_name") like $2)`
		whereParams = append(whereParams, "%"+strings.ToLower(search)+"%")
	}
	statement := contactSelectStatement + ` from "contacts" c ` + contactJoinStatement + ` ` + whereStatement
	rows, err = repository.DB.Query(statement, whereParams...)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}

func (repository ContactRepository) BrowseAllZakatDisbursement() (data []models.Contact, err error) {
	statement := `SELECT def.* FROM "contacts" def
	JOIN "user_zakats" uz ON uz."contact_id" = def."id"
	JOIN "transactions" t ON t."id" = uz."transaction_id"
	WHERE def."deleted_at" IS NULL AND t."transaction_type" = $1 AND t."payment_status" = $2
	AND t."is_disburse_allowed" = $3 AND t."is_disburse" = $4
	GROUP BY def."id"`
	rows, err := repository.DB.Query(statement,
		enums.KeyTransactionType1, enums.KeyPaymentStatus3, true, false,
	)
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
	statement := contactSelectStatement + ` from "contacts" c ` + contactJoinStatement + ` where c.` + column + `=$1 and c."deleted_at" is null`
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
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
