package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type PartnerExtraProductRepository struct {
	DB *sql.DB
}

func NewPartnerExtraProductRepository(DB *sql.DB) contracts.IPartnerExtraProductRepository {
	return &PartnerExtraProductRepository{DB: DB}
}

const partnerExtraProductSelectStatement = `select pep."id",pep."partner_id",mp."id",mp."name",mp."subscription_type",sp."price",sp."price_unit",sp."sessions"
                        from "partner_extra_products" pep
                       inner join "master_products" mp on mp."id"=pep."product_id"
                       inner join "setting_products" sp on sp."product_id"=mp."id"`

func (repository PartnerExtraProductRepository) BrowseByPartnerID(partnerID string) (data []models.PartnerExtraProduct, err error) {
	statement := partnerExtraProductSelectStatement + ` where pep."partner_id"=$1`
	rows, err := repository.DB.Query(statement, partnerID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.PartnerExtraProduct{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.PartnerID,
			&dataTemp.Product.ID,
			&dataTemp.Product.Name,
			&dataTemp.Product.SubscriptionType,
			&dataTemp.Product.Price,
			&dataTemp.Product.PriceUnit,
			&dataTemp.Product.Session,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}
	return data, err
}

func (repository PartnerExtraProductRepository) ReadBy(column, value string) (data models.PartnerExtraProduct, err error) {
	statement := partnerExtraProductSelectStatement + ` where ` + column + `=$1 and pep."deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.PartnerID,
		&data.Product.ID,
		&data.Product.Name,
		&data.Product.SubscriptionType,
		&data.Product.Price,
		&data.Product.PriceUnit,
		&data.Product.Session,
	)

	return data, err
}

func (PartnerExtraProductRepository) Add(partnerID, productID string, tx *sql.Tx) (err error) {
	statement := `insert into partner_extra_products ("partner_id","product_id") values($1,$2) returning "id"`
	_, err = tx.Exec(
		statement,
		partnerID,
		productID,
	)

	return err
}

func (PartnerExtraProductRepository) DeleteBy(column, value string, tx *sql.Tx) (err error) {
	statement := `delete from "partner_extra_products" where ` + column + `=$1`
	_, err = tx.Exec(statement, value)

	return err
}

func (repository PartnerExtraProductRepository) CountBy(column, value string) (res int, err error) {
	statement := `select count("id") from "partner_extra_products" where `+column+`=$1`
	err = repository.DB.QueryRow(statement,value).Scan(&res)

	return res,err
}
