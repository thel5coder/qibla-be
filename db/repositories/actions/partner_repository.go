package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type PartnerRepository struct {
	DB *sql.DB
}

func NewParterRepository(DB *sql.DB) contracts.IPartnerRepository {
	return &PartnerRepository{DB: DB}
}

const (
	partnerSelectStatement = `select p.*,u."username",c.* from "partners" p 
                 inner join "users" u on u."id"=p."user_id"
                 inner join "setting_products" sp on sp."product_id"=p."product_id"
                 inner join "contacts" c on c."id"=p."contact_id"`
)

func (repository PartnerRepository) Browse(search, order, sort string, limit, offset int) (data []models.Partner, count int, err error) {
	statement := partnerSelectStatement + ` where (lower(c."travel_agent_name") like $1 or lower(c."email") like $1) and p."deleted_at" is null
                 order by p.` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.Partner{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.UserID,
			&dataTemp.ContactID,
			&dataTemp.ContractNumber,
			&dataTemp.WebinarStatus,
			&dataTemp.WebsiteStatus,
			&dataTemp.DomainSite,
			&dataTemp.DomainErp,
			&dataTemp.Database,
			&dataTemp.DueDateAging,
			&dataTemp.IsActive,
			&dataTemp.Reason,
			&dataTemp.ProductID,
			&dataTemp.SubscriptionPeriod,
			&dataTemp.SubscriptionPeriodExpiredAt,
			&dataTemp.IsSubscriptionExpired,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.VerifiedAt,
			&dataTemp.IsPaid,
			&dataTemp.InvoicePublishDate,
			&dataTemp.PaidAt,
			&dataTemp.DatabaseUsername,
			&dataTemp.DatabasePassword,

			&dataTemp.UserName,

			&dataTemp.Contact.ID,
			&dataTemp.Contact.BranchName,
			&dataTemp.Contact.TravelAgentName,
			&dataTemp.Contact.Address,
			&dataTemp.Contact.Longitude,
			&dataTemp.Contact.Latitude,
			&dataTemp.Contact.AreaCode,
			&dataTemp.Contact.PhoneNumber,
			&dataTemp.Contact.SKNumber,
			&dataTemp.Contact.SKDate,
			&dataTemp.Contact.Accreditation,
			&dataTemp.Contact.AccreditationDate,
			&dataTemp.Contact.DirectorName,
			&dataTemp.Contact.DirectorContact,
			&dataTemp.Contact.PicName,
			&dataTemp.Contact.PicContact,
			&dataTemp.Contact.Logo,
			&dataTemp.Contact.VirtualAccountNumber,
			&dataTemp.Contact.AccountNumber,
			&dataTemp.Contact.AccountName,
			&dataTemp.Contact.AccountBankName,
			&dataTemp.Contact.AccountBankCode,
			&dataTemp.Contact.CreatedAt,
			&dataTemp.Contact.UpdatedAt,
			&dataTemp.Contact.DeletedAt,
			&dataTemp.Contact.Email,
			&dataTemp.Contact.IsZakatPartner,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count(p."id") from "partners" p 
                 inner join "users" u on u."id"=p."user_id"
                 inner join "contacts" c on c."id"=p."contact_id"
                 where (lower(c."travel_agent_name") like $1 or lower(c."email") like $1) and p."deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository PartnerRepository) BrowseProfilePartner(search, order, sort string, limit, offset int) (data []models.Partner, count int, err error) {
	statement := partnerSelectStatement + ` where (lower(c."travel_agent_name") like $1 or lower(c."email") like $1) and p."deleted_at" is null and p."verified_at" is not null
                 order by p.` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.Partner{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.UserID,
			&dataTemp.ContactID,
			&dataTemp.ContractNumber,
			&dataTemp.WebinarStatus,
			&dataTemp.WebsiteStatus,
			&dataTemp.DomainSite,
			&dataTemp.DomainErp,
			&dataTemp.Database,
			&dataTemp.DueDateAging,
			&dataTemp.IsActive,
			&dataTemp.Reason,
			&dataTemp.ProductID,
			&dataTemp.SubscriptionPeriod,
			&dataTemp.SubscriptionPeriodExpiredAt,
			&dataTemp.IsSubscriptionExpired,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.VerifiedAt,
			&dataTemp.IsPaid,
			&dataTemp.InvoicePublishDate,
			&dataTemp.PaidAt,
			&dataTemp.DatabaseUsername,
			&dataTemp.DatabasePassword,

			&dataTemp.UserName,

			&dataTemp.Contact.ID,
			&dataTemp.Contact.BranchName,
			&dataTemp.Contact.TravelAgentName,
			&dataTemp.Contact.Address,
			&dataTemp.Contact.Longitude,
			&dataTemp.Contact.Latitude,
			&dataTemp.Contact.AreaCode,
			&dataTemp.Contact.PhoneNumber,
			&dataTemp.Contact.SKNumber,
			&dataTemp.Contact.SKDate,
			&dataTemp.Contact.Accreditation,
			&dataTemp.Contact.AccreditationDate,
			&dataTemp.Contact.DirectorName,
			&dataTemp.Contact.DirectorContact,
			&dataTemp.Contact.PicName,
			&dataTemp.Contact.PicContact,
			&dataTemp.Contact.Logo,
			&dataTemp.Contact.VirtualAccountNumber,
			&dataTemp.Contact.AccountNumber,
			&dataTemp.Contact.AccountName,
			&dataTemp.Contact.AccountBankName,
			&dataTemp.Contact.AccountBankCode,
			&dataTemp.Contact.CreatedAt,
			&dataTemp.Contact.UpdatedAt,
			&dataTemp.Contact.DeletedAt,
			&dataTemp.Contact.Email,
			&dataTemp.Contact.IsZakatPartner,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count(p."id") from "partners" p 
                 inner join "users" u on u."id"=p."user_id"
                 inner join "contacts" c on c."id"=p."contact_id"
                 where (lower(c."travel_agent_name") like $1 or lower(c."email") like $1) and p."deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository PartnerRepository) ReadBy(column, value string) (data models.Partner, err error) {
	statement := partnerSelectStatement + ` where ` + column + `=$1 and p."deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.UserID,
		&data.ContactID,
		&data.ContractNumber,
		&data.WebinarStatus,
		&data.WebsiteStatus,
		&data.DomainSite,
		&data.DomainErp,
		&data.Database,
		&data.DueDateAging,
		&data.IsActive,
		&data.Reason,
		&data.ProductID,
		&data.SubscriptionPeriod,
		&data.SubscriptionPeriodExpiredAt,
		&data.IsSubscriptionExpired,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.VerifiedAt,
		&data.IsPaid,
		&data.InvoicePublishDate,
		&data.PaidAt,
		&data.DatabaseUsername,
		&data.DatabasePassword,

		&data.UserName,

		&data.Contact.ID,
		&data.Contact.BranchName,
		&data.Contact.TravelAgentName,
		&data.Contact.Address,
		&data.Contact.Longitude,
		&data.Contact.Latitude,
		&data.Contact.AreaCode,
		&data.Contact.PhoneNumber,
		&data.Contact.SKNumber,
		&data.Contact.SKDate,
		&data.Contact.Accreditation,
		&data.Contact.AccreditationDate,
		&data.Contact.DirectorName,
		&data.Contact.DirectorContact,
		&data.Contact.PicName,
		&data.Contact.PicContact,
		&data.Contact.Logo,
		&data.Contact.VirtualAccountNumber,
		&data.Contact.AccountNumber,
		&data.Contact.AccountName,
		&data.Contact.AccountBankName,
		&data.Contact.AccountBankCode,
		&data.Contact.CreatedAt,
		&data.Contact.UpdatedAt,
		&data.Contact.DeletedAt,
		&data.Contact.Email,
		&data.Contact.IsZakatPartner,
	)

	return data, err
}

func (PartnerRepository) Edit(input viewmodel.PartnerVm, tx *sql.Tx) (err error) {
	statement := `update "partners" set "contact_id"=$1,"product_id"=$2,"webinar_status"=$3,"website_status"=$4,"subscription_period"=$5,"updated_at"=$6 where "id"=$7`
	_, err = tx.Exec(
		statement,
		input.Contact.ID,
		input.Product.ProductID,
		input.WebinarStatus,
		input.WebsiteStatus,
		input.SubscriptionPeriod,
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ID,
	)

	return err
}

func (repository PartnerRepository) EditVerified(input viewmodel.PartnerVm) (res string, err error) {
	statement := `update "partners" set 
                 "domain_site"=$1, "domain_erp"=$2, "database"=$3, "reason"=$4, "due_date_aging"=$5, "is_active"=$6, "verified_at"=$7, "invoice_publish_date"=$8, "contract_number"=$9, 
                 "updated_at"=$10 where "id"=$11 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.DomainSite,
		input.DomainErp,
		input.Database,
		input.Reason,
		input.DueDateAging,
		input.IsActive,
		datetime.StrParseToTime(input.VerifiedAt, time.RFC3339),
		datetime.StrParseToTime(input.InvoicePublishDate, "2006-01-02"),
		input.ContractNumber,
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ID,
	).Scan(&res)

	return res, err
}

func (repository PartnerRepository) EditBoolStatus(ID, column, reason, updatedAt string, value bool) (res string, err error) {
	statement := `update "partners" set ` + column + `=$1, "reason"=$2, "updated_at"=$3 where "id"=$4 returning "id"`
	err = repository.DB.QueryRow(statement, value, reason, datetime.StrParseToTime(updatedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (PartnerRepository) EditPaymentStatus(ID, paidAt, updatedAt string, tx *sql.Tx) (err error) {
	statement := `update "partners" set  "paid_at"=$1, "is_paid"=$2, "updated_at"=$3 where "id"=$4 returning "id"`
	_,err = tx.Exec(statement, datetime.StrParseToTime(paidAt,time.RFC3339), datetime.StrParseToTime(updatedAt,time.RFC3339), ID)

	return err
}


func (PartnerRepository) Add(input viewmodel.PartnerVm, tx *sql.Tx) (res string, err error) {
	statement := `insert into "partners"
                 ("contact_id","user_id","product_id","subscription_period","webinar_status","website_status","is_active","is_paid","is_subscription_expired",
                 "created_at","updated_at")
                 values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning "id"`
	err = tx.QueryRow(
		statement,
		input.Contact.ID,
		input.UserID,
		input.Product.ProductID,
		input.SubscriptionPeriod,
		input.WebinarStatus,
		input.WebsiteStatus,
		input.IsActive,
		input.IsPaid,
		input.IsSubscriptionPeriodExpired,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (PartnerRepository) DeleteBy(column, value, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "partners" set "partners" set "updated_at"=$1, "deleted_at"=$2 where ` + column + `=$3`
	_, err = tx.Exec(
		statement,
		datetime.StrParseToTime(updatedAt, time.RFC3339),
		datetime.StrParseToTime(deletedAt, time.RFC3339),
		value,
	)

	return err
}

func (repository PartnerRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "partners" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "partners" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}


func scanRows(rows *sql.Rows) (data models.Partner,err error){

	return data,err
}