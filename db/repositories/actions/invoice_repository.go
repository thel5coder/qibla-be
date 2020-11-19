package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

// InvoiceRepository ...
type InvoiceRepository struct {
	DB *sql.DB
}

// NewInvoiceRepository ...
func NewInvoiceRepository(DB *sql.DB) contracts.IInvoiceRepository {
	return &InvoiceRepository{DB: DB}
}

const (
	// selectStatement ...
	invoiceSelectStatement = `SELECT inv."id", inv."transaction_type", inv."invoice_number", inv."fee_qibla",
	inv."total", inv."due_date", inv."due_date_period", inv."payment_status", inv."paid_date", inv."direction", inv."transaction_date", inv."updated_at"`
)

var (
	// DefaultDirection ...
	DefaultDirection = "out"

	// InvoiceFieldDate ...
	InvoiceFieldDate = "transaction_date"

	// DefaultSortAsc ...
	DefaultSortAsc = "asc"
	// DefaultSortDesc ...
	DefaultSortDesc = "desc"
)

// scanRow ...
func (repository InvoiceRepository) scanRow(row *sql.Row) (res models.Invoice, err error) {
	err = row.Scan(&res.ID, &res.Name, &res.TransactionType, &res.InvoiceNumber,
		&res.FeeQibla, &res.Total, &res.DueDate, &res.BillingStatus, &res.DueDatePeriod,
		&res.PaymentStatus, &res.PaidDate, &res.InvoiceStatus, &res.Direction, &res.TransactionDate, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// scanRows ...
func (repository InvoiceRepository) scanRows(rows *sql.Rows) (res models.Invoice, err error) {
	err = rows.Scan(&res.ID, &res.Name, &res.TransactionType, &res.InvoiceNumber,
		&res.FeeQibla, &res.Total, &res.DueDate, &res.BillingStatus, &res.DueDatePeriod,
		&res.PaymentStatus, &res.PaidDate, &res.InvoiceStatus, &res.Direction, &res.TransactionDate, &res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Browse ...
func (repository InvoiceRepository) Browse(order, sort string, limit, offset int) (data []models.Invoice, count int, err error) {
	statement := invoiceSelectStatement + ` FROM transactions inv
	WHERE inv."direction" = $1 AND inv."deleted_at" IS NULL
	ORDER BY inv.` + order + ` ` + sort + ` limit $2 offset $3`

	rows, err := repository.DB.Query(statement, DefaultDirection, limit, offset)
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

	statement = `SELECT COUNT(1) FROM transactions inv
	WHERE inv."direction" = $1 AND inv."deleted_at" IS NULL`
	err = repository.DB.QueryRow(statement, DefaultDirection).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}
