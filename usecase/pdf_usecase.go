package usecase

import (
	"github.com/rs/xid"
	"io/ioutil"
	"os"
	"qibla-backend/helpers/logruslogger"
	"qibla-backend/helpers/wkhtmltopdf"
	"qibla-backend/usecase/viewmodel"
	"strings"
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

	htmlNew := "../html_template/temp/" + xid.New().String() + ".html"
	err = ioutil.WriteFile(htmlNew, []byte(html), 0644)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "write_new_html", uc.ReqID)
		return res, err
	}

	res = "../html_template/temp/" + xid.New().String() + ".pdf"
	err = wkhtmltopdf.Generate(htmlNew, res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "generate_pdf", uc.ReqID)
		return res, err
	}

	os.Remove(htmlNew)

	return res, err
}

// Disbursement ...
func (uc PdfUseCase) Disbursement(id string) (res string, err error) {
	ctx := "PdfUseCase.Disbursement"

	disbursementUc := DisbursementUseCase{UcContract: uc.UcContract}
	_, err = disbursementUc.ReadByPk(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find", uc.ReqID)
		return res, err
	}

	// table := `<tr class="table-header"><td>Name</td><td>Email</td><td>Address</td><td>Phone</td><td>Profession</td><td>Age</td><td>Gender</td><td>Status</td></tr>`
	// for _, d := range donor {
	// 	age, _, _, _, _, _ := timepkg.DiffCustom(d.BirthDate, time.Now())
	// 	table += `<tr><td>` + d.Name + `</td><td>` + d.Email + `</td><td>` + d.Address + `</td><td>` + d.Phone + `</td><td>` + d.Profession + `</td><td>` + strconv.Itoa(age) + `</td><td>` + d.Gender + `</td><td>` + helper.StatusMapping(d.Status) + `</td></tr>`
	// }

	sourceFile := "../html_template/invoice/template.html"
	replace := []viewmodel.PdfReplaceVm{
		{
			From: "[title]",
			To:   "Donor",
		},
		// {
		// 	From: "[date]",
		// 	To:   timepkg.InFormatNoErr(time.Now(), DefaultLocation, "02-01-2006 15:04:05"),
		// },
		// {
		// 	From: "[table]",
		// 	To:   table,
		// },
	}
	res, err = uc.Generate(sourceFile, replace)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "generate_pdf", uc.ReqID)
		return res, err
	}

	return res, err
}
