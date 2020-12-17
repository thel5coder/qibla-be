package usecase

import (
	"io/ioutil"
	"os"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/number"
	timepkg "qibla-backend/pkg/time"
	"qibla-backend/pkg/wkhtmltopdf"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"

	"github.com/rs/xid"
)

// PdfUseCase ...
type PdfUseCase struct {
	*UcContract
}

// Generate ...
func (uc PdfUseCase) Generate(sourceFile string, replace []viewmodel.PdfReplaceVm) (res string, err error) {
	ctx := "PdfUseCase.Generate"

	htmlFile, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "read_source_html", uc.ReqID)
		return res, err
	}

	// Replace value
	html := string(htmlFile)
	for _, r := range replace {
		html = strings.Replace(html, r.From, r.To, -1)
	}

	htmlNew := "./../html_template/temp/" + xid.New().String() + ".html"
	err = ioutil.WriteFile(htmlNew, []byte(html), 0644)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "write_new_html", uc.ReqID)
		return res, err
	}

	res = "./../html_template/temp/" + xid.New().String() + ".pdf"
	err = wkhtmltopdf.Generate(htmlNew, res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "generate_pdf", uc.ReqID)
		return res, err
	}
	defer os.Remove(htmlNew)

	return res, err
}

// Disbursement ...
func (uc PdfUseCase) Disbursement(id string) (res string, err error) {
	ctx := "PdfUseCase.Disbursement"

	disbursementUc := DisbursementUseCase{UcContract: uc.UcContract}
	disbursement, err := disbursementUc.ReadByPk(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find", uc.ReqID)
		return res, err
	}

	userZakatUc := UserZakatUseCase{UcContract: uc.UcContract}
	userZakat, err := userZakatUc.BrowseAllByDisbursement(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_zakat", uc.ReqID)
		return res, err
	}

	table := ``
	for _, u := range userZakat {
		table += `<tr><td>` + u.TransactionInvoiceNumber + `</td><td>` + number.FormatCurrency(float64(u.Total), "Rp ", ".", ",", 2) + `</td><td>` + u.TypeZakat + `</td><td>` + timepkg.ConvertLocation(u.CreatedAt, time.RFC3339, "02-01-2006 15:04:05", DefaultLocation) + `</td><td>` + timepkg.ConvertLocation(disbursement.StartPeriod, time.RFC3339, "02-01-2006 15:04:05", DefaultLocation) + ` - ` + timepkg.ConvertLocation(disbursement.EndPeriod, time.RFC3339, "02-01-2006 15:04:05", DefaultLocation) + `</td></tr>`
	}

	sourceFile := "./../html_template/invoice/template.html"
	replace := []viewmodel.PdfReplaceVm{
		{
			From: "[base-url]",
			To:   os.Getenv("APP_BASE_URL") + "/html_template/invoice",
		},
		{
			From: "[company-name]",
			To:   disbursement.ContactTravelAgentName,
		},
		{
			From: "[company-address]",
			To:   disbursement.ContactAddress,
		},
		{
			From: "[company-province]",
			To:   "-",
		},
		{
			From: "[company-phone]",
			To:   disbursement.ContactPhoneNumber,
		},
		{
			From: "[total]",
			To:   number.FormatCurrency(disbursement.Total, "Rp ", ".", ",", 2),
		},
		{
			From: "[disbursement-date]",
			To:   timepkg.ConvertLocation(disbursement.DisburseAt, time.RFC3339, "02-01-2006 15:04:05", DefaultLocation),
		},
		{
			From: "[bank-name]",
			To:   disbursement.AccountBankName,
		},
		{
			From: "[table]",
			To:   table,
		},
	}
	res, err = uc.Generate(sourceFile, replace)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "generate_pdf", uc.ReqID)
		return res, err
	}

	return res, err
}
