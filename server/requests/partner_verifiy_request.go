package requests

type PartnerVerifyRequest struct {
	ContractNumber     string `json:"contract_number"`
	DomainSite         string `json:"domain_site"`
	DomainErp          string `json:"domain_erp"`
	Database           string `json:"database"`
	DatabaseUsername   string `json:"database_username"`
	DatabasePassword   string `json:"database_password"`
	Reason             string `json:"reason"`
	DueDateAging       int    `json:"due_date_aging"`
	IsActive           bool   `json:"is_active"`
	InvoicePublishDate string `json:"invoice_publish_date"`
}
