package usecase

import (
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
)


// ExcelUseCase ...
type ExcelUseCase struct {
	*UcContract
}

// Disbursement ...
func (uc ExcelUseCase) Disbursement(ids []string) (res string, err error) {
	disbursementUc := DisbursementUseCase{UcContract: uc.UcContract}
	_, err = disbursementUc.BrowseAll("")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "find_donor")
		return res, err
	}
	//
	//data := []map[string]string{
	//	{"A1": "Name", "B1": "Email", "C1": "Address", "D1": "Phone", "E1": "Profession", "F1": "Age", "G1": "Gender", "H1": "Status"},
	//}
	//for index, d := range donor {
	//	i := strconv.Itoa(index + 2)
	//	age, _, _, _, _, _ := timepkg.DiffCustom(d.BirthDate, time.Now())
	//	data = append(data, map[string]string{
	//		"A" + i: d.Name, "B" + i: d.Email, "C" + i: d.Address, "D" + i: d.Phone, "E" + i: d.Profession, "F" + i: strconv.Itoa(age), "G" + i: d.Gender, "H" + i: helper.StatusMapping(d.Status),
	//	})
	//}

	//res = uc.EnvConfig["FILE_STATIC_FILE"] + "/" + xid.New().String() + ".xlsx"

	//err = excelizepkg.Generate(&data, res)
	//if err != nil {
	//	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "generate_excel", uc.ReqID)
	//	return res, err
	//}

	return res, err
}

