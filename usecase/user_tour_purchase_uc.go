package usecase

import (
	"errors"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

// UserTourPurchaseUseCase ...
type UserTourPurchaseUseCase struct {
	*UcContract
}

// Browse ...
func (uc UserTourPurchaseUseCase) Browse(userID, status, order, sort string, page, limit int) (res []viewmodel.UserTourPurchaseVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewUserTourPurchaseRepository(uc.DB)

	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)
	userTourPurchases, count, err := repository.Browse(userID, status, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, userTourPurchase := range userTourPurchases {
		res = append(res, uc.buildBody(&userTourPurchase))
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

// BrowseAll ...
func (uc UserTourPurchaseUseCase) BrowseAll() (res []viewmodel.UserTourPurchaseVm, err error) {
	repository := actions.NewUserTourPurchaseRepository(uc.DB)

	userTourPurchases, err := repository.BrowseAll()
	if err != nil {
		return res, err
	}

	for _, userTourPurchase := range userTourPurchases {
		res = append(res, uc.buildBody(&userTourPurchase))
	}

	return res, err
}

// BrowseBy ...
func (uc UserTourPurchaseUseCase) BrowseBy(column, value, operator string) (res []viewmodel.UserTourPurchaseVm, err error) {
	repository := actions.NewUserTourPurchaseRepository(uc.DB)
	userTourPurchases, err := repository.BrowseBy(column, value, operator)

	for _, userTourPurchase := range userTourPurchases {
		res = append(res, uc.buildBody(&userTourPurchase))
	}

	return res, err
}

// ReadBy ...
func (uc UserTourPurchaseUseCase) ReadBy(column, value string) (res viewmodel.UserTourPurchaseVm, err error) {
	repository := actions.NewUserTourPurchaseRepository(uc.DB)
	userTourPurchase, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = uc.buildBody(&userTourPurchase)

	return res, err
}

// ReadByPk ...
func (uc UserTourPurchaseUseCase) ReadByPk(ID string) (res viewmodel.UserTourPurchaseVm, err error) {
	res, err = uc.ReadBy("uz.id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc UserTourPurchaseUseCase) checkInput(userID string, input *requests.UserTourPurchaseRequest, oldData *viewmodel.UserTourPurchaseVm) (err error) {

	return err
}

// Add ...
func (uc UserTourPurchaseUseCase) Add(userID string, input *requests.UserTourPurchaseRequest) (res viewmodel.UserTourPurchaseVm, err error) {
	err = uc.checkInput(userID, input, &res)
	if err != nil {
		return res, err
	}

	// transactionUseCase := TransactionUseCase{UcContract: uc.UcContract}
	// transaction, err := transactionUseCase.AddTransactionZakat(input)
	// if err != nil {
	// 	return res, err
	// }

	now := time.Now().UTC()
	res = viewmodel.UserTourPurchaseVm{
		TourPackageID:         input.TourPackageID,
		PaymentType:           input.PaymentType,
		CustomerName:          input.CustomerName,
		CustomerIdentityType:  input.CustomerIdentityType,
		IdentityNumber:        input.IdentityNumber,
		FullName:              input.FullName,
		Sex:                   input.Sex,
		BirthDate:             input.BirthDate,
		BirthPlace:            input.BirthPlace,
		PhoneNumber:           input.PhoneNumber,
		CityID:                input.CityID,
		MaritalStatus:         input.MaritalStatus,
		CustomerAddress:       input.CustomerAddress,
		UserID:                userID,
		ContactID:             input.ContactID,
		OldUserTourPurchaseID: input.OldUserTourPurchaseID,
		CancelationFee:        input.CancelationFee,
		Total:                 input.Total,
		Status:                models.UserTourPurchaseStatusPending,
		CreatedAt:             now.Format(time.RFC3339),
		UpdatedAt:             now.Format(time.RFC3339),
	}
	repository := actions.NewUserTourPurchaseRepository(uc.DB)
	res.ID, err = repository.Add(res, uc.TX)
	if err != nil {
		return res, err
	}

	return res, err
}

// Update ...
func (uc UserTourPurchaseUseCase) Update(userID, id string, input *requests.UserTourPurchaseRequest) (res viewmodel.UserTourPurchaseVm, err error) {
	oldData, err := uc.ReadByPk(id)
	if err != nil {
		return res, err
	}
	if oldData.UserID != userID {
		return res, errors.New(messages.InvalidUser)
	}

	err = uc.checkInput(userID, input, &oldData)
	if err != nil {
		return res, err
	}

	// transactionUseCase := TransactionUseCase{UcContract: uc.UcContract}
	// transaction, err := transactionUseCase.AddTransactionZakat(input)
	// if err != nil {
	// 	return res, err
	// }

	now := time.Now().UTC()
	res = viewmodel.UserTourPurchaseVm{
		ID:                    id,
		TourPackageID:         input.TourPackageID,
		PaymentType:           input.PaymentType,
		CustomerName:          input.CustomerName,
		CustomerIdentityType:  input.CustomerIdentityType,
		IdentityNumber:        input.IdentityNumber,
		FullName:              input.FullName,
		Sex:                   input.Sex,
		BirthDate:             input.BirthDate,
		BirthPlace:            input.BirthPlace,
		PhoneNumber:           input.PhoneNumber,
		CityID:                input.CityID,
		MaritalStatus:         input.MaritalStatus,
		CustomerAddress:       input.CustomerAddress,
		UserID:                userID,
		ContactID:             input.ContactID,
		OldUserTourPurchaseID: input.OldUserTourPurchaseID,
		CancelationFee:        input.CancelationFee,
		Total:                 input.Total,
		Status:                models.UserTourPurchaseStatusPending,
		CreatedAt:             now.Format(time.RFC3339),
		UpdatedAt:             now.Format(time.RFC3339),
	}
	repository := actions.NewUserTourPurchaseRepository(uc.DB)
	err = repository.Edit(res, uc.TX)
	if err != nil {
		return res, err
	}

	return res, err
}

// Delete ...
func (uc UserTourPurchaseUseCase) Delete(ID string) (err error) {
	now := time.Now().UTC().Format(time.RFC3339)
	repository := actions.NewUserTourPurchaseRepository(uc.DB)
	err = repository.Delete(ID, now, now, uc.TX)
	if err != nil {
		return err
	}

	return err
}

func (uc UserTourPurchaseUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewUserTourPurchaseRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc UserTourPurchaseUseCase) buildBody(data *models.UserTourPurchase) (res viewmodel.UserTourPurchaseVm) {
	return viewmodel.UserTourPurchaseVm{
		ID:                     data.ID,
		TourPackageID:          data.TourPackageID.String,
		PaymentType:            data.PaymentType.String,
		CustomerName:           data.CustomerName.String,
		CustomerIdentityType:   data.CustomerIdentityType.String,
		IdentityNumber:         data.IdentityNumber.String,
		FullName:               data.FullName.String,
		Sex:                    data.Sex.String,
		BirthDate:              data.BirthDate.String,
		BirthPlace:             data.BirthPlace.String,
		PhoneNumber:            data.PhoneNumber.String,
		CityID:                 data.CityID.String,
		MaritalStatus:          data.MaritalStatus.String,
		CustomerAddress:        data.CustomerAddress.String,
		UserID:                 data.UserID.String,
		UserEmail:              data.User.Email.String,
		UserName:               data.User.Name.String,
		ContactID:              data.ContactID.String,
		ContactBranchName:      data.Contact.BranchName.String,
		ContactTravelAgentName: data.Contact.TravelAgentName.String,
		OldUserTourPurchaseID:  data.OldUserTourPurchaseID.String,
		CancelationFee:         data.CancelationFee.Float64,
		Total:                  data.Total.Float64,
		CreatedAt:              data.CreatedAt,
		UpdatedAt:              data.UpdatedAt,
		DeletedAt:              data.DeletedAt.String,
	}
}
