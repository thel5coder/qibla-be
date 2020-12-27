package usecase

import (
	"errors"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/hashing"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/messages"
	"qibla-backend/pkg/str"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type PartnerUseCase struct {
	*UcContract
}

func (uc PartnerUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.PartnerVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewParterRepository(uc.DB)

	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	partners, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, partner := range partners {
		res = append(res, uc.buildBody(partner))
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc PartnerUseCase) BrowseProfilePartner(search, order, sort string, page, limit int) (res []viewmodel.PartnerVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewParterRepository(uc.DB)

	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	partners, count, err := repository.BrowseProfilePartner(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, partner := range partners {
		res = append(res, uc.buildBody(partner))
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc PartnerUseCase) ReadBy(column, value string) (res viewmodel.PartnerVm, err error) {
	repository := actions.NewParterRepository(uc.DB)
	partner, err := repository.ReadBy(column, value)
	res = uc.buildBody(partner)

	return res, err
}

func (uc PartnerUseCase) Edit(ID string, input *requests.PartnerRegisterRequest) (err error) {
	repository := actions.NewParterRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	err = uc.ValidateIsPartnerExist(ID, input.ContactID)
	if err != nil {
		return err
	}

	//init transaction
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	partnerExtraProductUc := PartnerExtraProductUseCase{UcContract: uc.UcContract}
	err = partnerExtraProductUc.Store(ID, input.ExtraProducts)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	body := viewmodel.PartnerVm{
		ID:                 ID,
		WebinarStatus:      input.WebinarStatus,
		WebsiteStatus:      input.WebsiteStatus,
		SubscriptionPeriod: input.SubscriptionPeriod,
		UpdatedAt:          now,
		Contact:            viewmodel.ContactVm{ID: input.ContactID},
		Product:            viewmodel.SettingProductVm{ProductID: input.ProductID},
	}

	err = repository.Edit(body, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	uc.TX.Commit()

	return nil
}

func (uc PartnerUseCase) EditVerify(ID string, input *requests.PartnerVerifyRequest) (err error) {
	repository := actions.NewParterRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	body := viewmodel.PartnerVm{
		ID:                 ID,
		DomainSite:         input.DomainSite,
		DomainErp:          input.DomainErp,
		Database:           input.Database,
		DatabaseUsername:   input.DatabaseUsername,
		DatabasePassword:   input.DatabasePassword,
		Reason:             input.Reason,
		DueDateAging:       input.DueDateAging,
		IsActive:           input.IsActive,
		VerifiedAt:         now,
		InvoicePublishDate: input.InvoicePublishDate,
		ContractNumber:     input.ContractNumber,
	}
	_, err = repository.EditVerified(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc PartnerUseCase) EditBoolStatus(ID, column, reason, userID, password string, value bool) (err error) {
	repository := actions.NewParterRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	userUc := UserUseCase{UcContract: uc.UcContract}

	isPasswordValid, err := userUc.IsPasswordValid(userID, password)
	if err != nil {
		return err
	}
	if !isPasswordValid {
		return err
	}
	_, err = repository.EditBoolStatus(ID, column, reason, now, value)
	if err != nil {
		return err
	}

	return nil
}

func (uc PartnerUseCase) EditPaymentStatus(ID string) (err error) {
	repository := actions.NewParterRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	err = repository.EditPaymentStatus(ID, now, now, uc.TX)
	if err != nil {
		return err
	}

	return nil
}

func (uc PartnerUseCase) Add(input *requests.PartnerRegisterRequest) (res viewmodel.TransactionVm, err error) {
	repository := actions.NewParterRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	contactUc := ContactUseCase{UcContract: uc.UcContract}
	contact, err := contactUc.ReadByPk(input.ContactID)
	if err != nil {
		return res, err
	}

	err = uc.ValidateIsPartnerExist("", input.ContactID)
	if err != nil {
		return res, err
	}

	//init transaction
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return res, err
	}

	//add user
	userUc := UserUseCase{UcContract: uc.UcContract}
	password, _ := hashing.HashAndSalt(str.RandomString(6))
	userID, err := userUc.Add(contact.TravelAgentName, input.UserName, contact.Email, contact.PhoneNumber, "ee694491-5166-441b-8262-9745bf866aa9", password, "", true, false)
	if err != nil {
		uc.TX.Rollback()

		return res, err
	}

	//add partners
	body := viewmodel.PartnerVm{
		Contact:                     viewmodel.ContactVm{ID: input.ContactID},
		UserID:                      userID,
		Product:                     viewmodel.SettingProductVm{ProductID: input.ProductID},
		SubscriptionPeriod:          input.SubscriptionPeriod,
		WebsiteStatus:               input.WebsiteStatus,
		WebinarStatus:               input.WebinarStatus,
		IsActive:                    false,
		IsPaid:                      false,
		IsSubscriptionPeriodExpired: false,
		CreatedAt:                   now,
		UpdatedAt:                   now,
	}
	body.ID, err = repository.Add(body, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return res, err
	}

	//add extra product
	if len(input.ExtraProducts) > 0 {
		partnerExtraProductUc := PartnerExtraProductUseCase{UcContract: uc.UcContract}
		err = partnerExtraProductUc.Store(body.ID, input.ExtraProducts)
		if err != nil {
			uc.TX.Rollback()

			return res, err
		}
		transactionUc := TransactionUseCase{UcContract: uc.UcContract}
		res, err = transactionUc.AddTransactionRegisterPartner(userID, input.BankName, input.PaymentMethodCode, 7, input.ExtraProducts, contact)
		if err != nil {
			uc.TX.Rollback()

			return res, err
		}
	}

	uc.TX.Commit()

	return res, nil
}

func (uc PartnerUseCase) DeleteBy(column, value string) (err error) {
	repository := actions.NewParterRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.countBy("", column, value)
	if err != nil {
		return err
	}

	if count > 0 {
		err = repository.DeleteBy(column, value, now, now, uc.TX)
		if err != nil {

			return err
		}
	}

	return nil
}

func (uc PartnerUseCase) DeleteByPk(ID string) (err error) {
	//init transaction
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	err = uc.DeleteBy("id", ID)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	partnerExtraProductUc := PartnerExtraProductUseCase{UcContract: uc.UcContract}
	err = partnerExtraProductUc.DeleteBy("partner_id", ID)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	uc.TX.Commit()

	return nil
}

func (uc PartnerUseCase) ValidateIsPartnerExist(ID, contactID string) (err error) {
	count, err := uc.countBy(ID, "contact_id", contactID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	return nil
}

func (uc PartnerUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewParterRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc PartnerUseCase) buildBody(data models.Partner) (res viewmodel.PartnerVm) {
	fileUc := FileUseCase{UcContract: uc.UcContract}
	settingProductUc := SettingProductUseCase{UcContract: uc.UcContract}
	partnerExtraProductUc := PartnerExtraProductUseCase{UcContract: uc.UcContract}

	file, _ := fileUc.ReadByPk(data.Contact.Logo)
	product, _ := settingProductUc.ReadBy("product_id", data.ProductID)
	fmt.Println(data.ID)
	partnerExtraProducts, err := partnerExtraProductUc.BrowseByPartnerID(data.ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-partnerExtraProduct-browseByPartnerID")
	}
	fmt.Println(partnerExtraProducts)

	return viewmodel.PartnerVm{
		ID:                          data.ID,
		UserID:                      data.UserID,
		UserName:                    data.UserName,
		ContractNumber:              data.ContractNumber.String,
		WebinarStatus:               data.WebinarStatus,
		WebsiteStatus:               data.WebsiteStatus,
		DomainSite:                  data.DomainSite.String,
		DomainErp:                   data.DomainErp.String,
		Database:                    data.Database.String,
		DatabaseUsername:            data.DatabaseUsername.String,
		DatabasePassword:            data.DatabasePassword.String,
		InvoicePublishDate:          data.InvoicePublishDate.String,
		DueDateAging:                int(data.DueDateAging.Int32),
		IsActive:                    data.IsActive,
		IsPaid:                      data.IsPaid,
		Reason:                      data.Reason.String,
		SubscriptionPeriod:          data.SubscriptionPeriod,
		SubscriptionPeriodExpiredAt: data.SubscriptionPeriodExpiredAt.String,
		IsSubscriptionPeriodExpired: data.IsSubscriptionExpired.Bool,
		CreatedAt:                   data.CreatedAt,
		UpdatedAt:                   data.UpdatedAt,
		VerifiedAt:                  data.VerifiedAt.String,
		PaidAt:                      data.PaidAt.String,
		Contact: viewmodel.ContactVm{
			ID:                   data.Contact.ID,
			BranchName:           data.Contact.BranchName.String,
			TravelAgentName:      data.Contact.TravelAgentName.String,
			Address:              data.Contact.Address.String,
			Longitude:            data.Contact.Longitude.String,
			Latitude:             data.Contact.Latitude.String,
			AreaCode:             data.Contact.AreaCode,
			PhoneNumber:          data.Contact.PhoneNumber,
			SKNumber:             data.Contact.SKNumber.String,
			SKDate:               data.Contact.SKDate.String,
			Accreditation:        data.Contact.Accreditation.String,
			AccreditationDate:    data.Contact.AccreditationDate.String,
			DirectorName:         data.Contact.DirectorName.String,
			DirectorContact:      data.Contact.DirectorContact.String,
			PicName:              data.Contact.PicName,
			PicContact:           data.Contact.PicContact,
			FileLogo:             file,
			VirtualAccountNumber: data.Contact.VirtualAccountNumber.String,
			AccountNumber:        data.Contact.AccountNumber,
			AccountName:          data.Contact.AccountName,
			AccountBankName:      data.Contact.AccountBankName,
			AccountBankCode:      data.Contact.AccountBankCode,
			Email:                data.Contact.Email,
			IsZakatPartner:       data.Contact.IsZakatPartner,
			CreatedAt:            data.Contact.CreatedAt,
			UpdatedAt:            data.Contact.UpdatedAt,
			DeletedAt:            data.Contact.DeletedAt.String,
		},
		Product:      product,
		ExtraProduct: partnerExtraProducts,
	}
}
