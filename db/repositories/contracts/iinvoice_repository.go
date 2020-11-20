package contracts

import "qibla-backend/db/models"

// IInvoiceRepository ...
type IInvoiceRepository interface {
	Browse(order, sort string, limit, offset int) (data []models.Invoice, count int, err error)
}
