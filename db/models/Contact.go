package models

import "database/sql"

type Contact struct {
	ID                   string         `db:"id"`
	BranchName           sql.NullString `db:"branch_name"`
	TravelAgentName      sql.NullString `db:"travel_agent_name"`
	Address              sql.NullString `db:"address"`
	Longitude            sql.NullString `db:"longitude"`
	Latitude             sql.NullString `db:"latitude"`
	AreaCode             string         `db:"area_code"`
	PhoneNumber          string         `db:"phone_number"`
	SKNumber             sql.NullString `db:"sk_number"`
	SKDate               sql.NullString `db:"sk_date"`
	Accreditation        sql.NullString `db:"accreditation"`
	AccreditationDate    sql.NullString `db:"accreditation_date"`
	DirectorName         sql.NullString `db:"director_name"`
	DirectorContact      sql.NullString `db:"director_contact"`
	PicName              string         `db:"pic_name"`
	PicContact           string         `db:"pic_contact"`
	Logo                 string         `db:"logo"`
	VirtualAccountNumber sql.NullString `db:"virtual_account_number"`
	AccountNumber        string         `db:"account_number"`
	AccountName          string         `db:"account_name"`
	AccountBankName      string         `db:"account_bank_name"`
	AccountBankCode      string         `db:"account_bank_code"`
	Email                string         `db:"email"`
	IsZakatPartner       bool           `db:"is_zakat_partner"`
	CreatedAt            string         `db:"created_at"`
	UpdatedAt            string         `db:"updated_at"`
	DeletedAt            sql.NullString `db:"deleted_at"`
	LogoPath             sql.NullString `db:"path"`
	LogoName             sql.NullString `db:"logo_name"`
}
