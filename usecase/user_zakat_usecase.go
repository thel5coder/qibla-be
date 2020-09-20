package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/enums"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

// UserZakatUseCase ...
type UserZakatUseCase struct {
	*UcContract
}

// Browse ...
func (uc UserZakatUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.UserZakatVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewUserZakatRepository(uc.DB)

	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)
	userZakats, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, userZakat := range userZakats {
		res = append(res, uc.buildBody(&userZakat))
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

// BrowseAll ...
func (uc UserZakatUseCase) BrowseAll() (res []viewmodel.UserZakatVm, err error) {
	repository := actions.NewUserZakatRepository(uc.DB)

	userZakats, err := repository.BrowseAll()
	if err != nil {
		return res, err
	}

	for _, userZakat := range userZakats {
		res = append(res, uc.buildBody(&userZakat))
	}

	return res, err
}

// BrowseBy ...
func (uc UserZakatUseCase) BrowseBy(column, value, operator string) (res []viewmodel.UserZakatVm, err error) {
	repository := actions.NewUserZakatRepository(uc.DB)
	userZakats, err := repository.BrowseBy(column, value, operator)

	for _, userZakat := range userZakats {
		res = append(res, uc.buildBody(&userZakat))
	}

	return res, err
}

// ReadBy ...
func (uc UserZakatUseCase) ReadBy(column, value string) (res viewmodel.UserZakatVm, err error) {
	repository := actions.NewUserZakatRepository(uc.DB)
	userZakat, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = uc.buildBody(&userZakat)

	return res, err
}

// ReadByPk ...
func (uc UserZakatUseCase) ReadByPk(ID string) (res viewmodel.UserZakatVm, err error) {
	res, err = uc.ReadBy("sp.id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc UserZakatUseCase) checkInput(input *requests.UserZakatRequest) (res int, err error) {
	masterZakatUc := MasterZakatUseCase{UcContract: uc.UcContract}
	masterZakat, err := masterZakatUc.ReadBy("type_zakat", input.TypeZakat)
	if err != nil {
		return res, err
	}
	input.MasterZakatID = masterZakat.ID
	input.CurrentGoldPrice = masterZakat.CurrentGoldPrice
	input.GoldNishab = masterZakat.GoldNishab

	if input.TypeZakat == enums.KeyTypeZakatEnum1 {
		input.Total = int32(float64(input.Wealth-input.CurrentGoldPrice*input.GoldNishab) * 2.5 / 100)
	} else {
		input.Total = int32(float64(input.Wealth) * 2.5 / 100)
	}

	return res, err
}

// Edit ...
func (uc UserZakatUseCase) Edit(ID string, input *requests.UserZakatRequest) (err error) {
	now := time.Now().UTC().Format(time.RFC3339)
	body := viewmodel.UserZakatVm{
		ID:               ID,
		UserID:           uc.UserID,
		TransactionID:    input.TransactionID,
		ContactID:        input.ContactID,
		MasterZakatID:    input.MasterZakatID,
		TypeZakat:        input.TypeZakat,
		CurrentGoldPrice: input.CurrentGoldPrice,
		GoldNishab:       input.GoldNishab,
		Wealth:           input.Wealth,
		Total:            input.Total,
		UpdatedAt:        now,
	}
	repository := actions.NewUserZakatRepository(uc.DB)
	err = repository.Edit(body, uc.TX)
	if err != nil {
		return err
	}

	return nil
}

// Add ...
func (uc UserZakatUseCase) Add(input *requests.UserZakatRequest) (err error) {
	transactionUseCase := TransactionUseCase{UcContract: uc.UcContract}
	transaction, err := transactionUseCase.AddTransactionZakat(input)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	body := viewmodel.UserZakatVm{
		UserID:           uc.UserID,
		TransactionID:    transaction.ID,
		ContactID:        input.ContactID,
		MasterZakatID:    input.MasterZakatID,
		TypeZakat:        input.TypeZakat,
		CurrentGoldPrice: input.CurrentGoldPrice,
		GoldNishab:       input.GoldNishab,
		Wealth:           input.Wealth,
		Total:            input.Total,
		CreatedAt:        now.Format(time.RFC3339),
		UpdatedAt:        now.Format(time.RFC3339),
	}
	repository := actions.NewUserZakatRepository(uc.DB)
	body.ID, err = repository.Add(body, uc.TX)
	if err != nil {
		return err
	}

	return err
}

// Delete ...
func (uc UserZakatUseCase) Delete(ID string) (err error) {
	now := time.Now().UTC().Format(time.RFC3339)
	repository := actions.NewUserZakatRepository(uc.DB)
	err = repository.Delete(ID, now, now, uc.TX)
	if err != nil {
		return err
	}

	return err
}

func (uc UserZakatUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewUserZakatRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc UserZakatUseCase) buildBody(data *models.UserZakat) (res viewmodel.UserZakatVm) {
	return viewmodel.UserZakatVm{
		UserID:                       data.UserID.String,
		UserEmail:                    data.User.Email.String,
		UserName:                     data.User.Name.String,
		TransactionID:                data.TransactionID.String,
		TransactionInvoiceNumber:     data.Transaction.InvoiceNumber.String,
		TransactionPaymentMethodCode: data.Transaction.PaymentMethodCode.Int32,
		TransactionPaymentStatus:     data.Transaction.PaymentStatus.String,
		ContactID:                    data.ContactID.String,
		ContactBranchName:            data.Contact.BranchName.String,
		ContactTravelAgentName:       data.Contact.TravelAgentName.String,
		MasterZakatID:                data.MasterZakatID.String,
		TypeZakat:                    data.TypeZakat.String,
		CurrentGoldPrice:             data.CurrentGoldPrice.Int32,
		GoldNishab:                   data.GoldNishab.Int32,
		Wealth:                       data.Wealth.Int32,
		Total:                        data.Total.Int32,
		CreatedAt:                    data.CreatedAt,
		UpdatedAt:                    data.UpdatedAt,
		DeletedAt:                    data.DeletedAt.String,
	}
}
