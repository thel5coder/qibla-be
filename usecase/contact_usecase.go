package usecase

import (
	"errors"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type ContactUseCase struct {
	*UcContract
}

func (uc ContactUseCase) Browse(filters map[string]interface{}, order, sort string, page, limit int) (res []viewmodel.ContactVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewContactRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	contacts, count, err := repository.Browse(filters, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, contact := range contacts {
		res = append(res, uc.buildBody(contact))
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc ContactUseCase) BrowseAll(search string, isZakatPartner bool) (res []viewmodel.ContactVm, err error) {
	repository := actions.NewContactRepository(uc.DB)
	contacts, err := repository.BrowseAll(search, isZakatPartner)
	if err != nil {
		return res, err
	}

	for _, contact := range contacts {
		res = append(res, uc.buildBody(contact))
	}

	return res, err
}

func (uc ContactUseCase) BrowseAllZakatDisbursement() (res []viewmodel.ContactVm, err error) {
	repository := actions.NewContactRepository(uc.DB)
	fileUc := FileUseCase{UcContract: uc.UcContract}
	contacts, err := repository.BrowseAllZakatDisbursement()
	if err != nil {
		return res, err
	}

	for _, contact := range contacts {
		file, _ := fileUc.ReadByPk(contact.Logo)
		res = append(res, viewmodel.ContactVm{
			ID:                   contact.ID,
			BranchName:           contact.BranchName.String,
			TravelAgentName:      contact.TravelAgentName.String,
			Address:              contact.Address.String,
			Longitude:            contact.Longitude.String,
			Latitude:             contact.Latitude.String,
			AreaCode:             contact.AreaCode,
			PhoneNumber:          contact.PhoneNumber,
			SKNumber:             contact.SKNumber.String,
			SKDate:               contact.SKDate.String,
			Accreditation:        contact.Accreditation.String,
			AccreditationDate:    contact.AccreditationDate.String,
			DirectorName:         contact.DirectorName.String,
			DirectorContact:      contact.DirectorContact.String,
			PicName:              contact.PicName,
			PicContact:           contact.PicContact,
			FileLogo:             file,
			VirtualAccountNumber: contact.VirtualAccountNumber.String,
			AccountNumber:        contact.AccountNumber,
			AccountName:          contact.AccountName,
			AccountBankName:      contact.AccountBankName,
			AccountBankCode:      contact.AccountBankCode,
			Email:                contact.Email,
			IsZakatPartner:       contact.IsZakatPartner,
			CreatedAt:            contact.CreatedAt,
			UpdatedAt:            contact.UpdatedAt,
			DeletedAt:            contact.DeletedAt.String,
		})
	}

	return res, err
}

func (uc ContactUseCase) BrowseAllZakatPlace(search string) (res []viewmodel.ZakatPlaceVm, err error) {
	contactZakats, err := uc.BrowseAll(search, true)
	if err != nil {
		return res, err
	}

	for _, contactZakat := range contactZakats {
		res = append(res, viewmodel.ZakatPlaceVm{
			ID:   contactZakat.ID,
			Name: contactZakat.TravelAgentName,
		})
	}

	return res, err
}

func (uc ContactUseCase) ReadBy(column, value string) (res viewmodel.ContactVm, err error) {
	repository := actions.NewContactRepository(uc.DB)
	contact, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}
	res = uc.buildBody(contact)

	return res, err
}

func (uc ContactUseCase) Edit(ID string, input *requests.ContactRequest) (err error) {
	repository := actions.NewContactRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	err = uc.duplicateCheck(ID, input)
	if err != nil {
		return err
	}

	body := viewmodel.ContactVm{
		ID:                ID,
		BranchName:        input.BranchName,
		TravelAgentName:   input.TravelAgentName,
		Address:           input.Address,
		Longitude:         input.Longitude,
		Latitude:          input.Latitude,
		AreaCode:          input.AreaCode,
		PhoneNumber:       input.PhoneNumber,
		SKNumber:          input.SKNumber,
		SKDate:            input.SKDate,
		Accreditation:     input.Accreditation,
		AccreditationDate: input.AccreditationDate,
		DirectorName:      input.DirectorName,
		DirectorContact:   input.DirectorContact,
		PicName:           input.PicName,
		PicContact:        input.PicContact,
		FileLogo: viewmodel.FileVm{
			ID: input.Logo,
		},
		VirtualAccountNumber: input.VirtualAccountNumber,
		AccountNumber:        input.AccountNumber,
		AccountName:          input.AccountName,
		AccountBankName:      input.AccountBankName,
		AccountBankCode:      input.AccountBankCode,
		Email:                input.Email,
		IsZakatPartner:       input.IsZakatPartner,
		UpdatedAt:            now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc ContactUseCase) Add(input *requests.ContactRequest) (err error) {
	repository := actions.NewContactRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	err = uc.duplicateCheck("", input)
	if err != nil {
		return err
	}

	body := viewmodel.ContactVm{
		BranchName:        input.BranchName,
		TravelAgentName:   input.TravelAgentName,
		Address:           input.Address,
		Longitude:         input.Longitude,
		Latitude:          input.Latitude,
		AreaCode:          input.AreaCode,
		PhoneNumber:       input.PhoneNumber,
		SKNumber:          input.SKNumber,
		SKDate:            input.SKDate,
		Accreditation:     input.Accreditation,
		AccreditationDate: input.AccreditationDate,
		DirectorName:      input.DirectorName,
		DirectorContact:   input.DirectorContact,
		PicName:           input.PicName,
		PicContact:        input.PicContact,
		FileLogo: viewmodel.FileVm{
			ID: input.Logo,
		},
		VirtualAccountNumber: input.VirtualAccountNumber,
		AccountNumber:        input.AccountNumber,
		AccountName:          input.AccountName,
		AccountBankName:      input.AccountBankName,
		AccountBankCode:      input.AccountBankCode,
		Email:                input.Email,
		IsZakatPartner:       input.IsZakatPartner,
		CreatedAt:            now,
		UpdatedAt:            now,
	}
	_, err = repository.Add(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc ContactUseCase) Delete(ID string) (err error) {
	repository := actions.NewContactRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.countBy("", "id", ID)
	if err != nil {
		return err
	}

	if count > 0 {
		_, err = repository.Delete(ID, now, now)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc ContactUseCase) duplicateCheck(ID string, input *requests.ContactRequest) (err error) {
	//cek sk
	skCount, err := uc.countBy(ID, "sk_number", input.SKNumber)
	if err != nil {
		return err
	}
	if skCount > 0 {
		return errors.New("nomor sk sudah ada")
	}

	//cek akreditasi
	accreditationCount, err := uc.countBy(ID, "accreditation", input.Accreditation)
	if err != nil {
		return err
	}
	if accreditationCount > 0 {
		return errors.New("akreditasi sudah ada")
	}

	//check email
	emailCount, err := uc.countBy(ID, "email", input.Email)
	if err != nil {
		return err
	}
	if emailCount > 0 {
		return errors.New(messages.EmailAlreadyExist)
	}

	return nil
}

func (uc ContactUseCase) ReadByPk(ID string) (res viewmodel.ContactVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}

	return res, err
}

func (uc ContactUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewContactRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}

//buld body
func (uc ContactUseCase) buildBody(model models.Contact) viewmodel.ContactVm {
	path, err := uc.AWSS3.GetURL(model.LogoPath.String)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-aws-s3-getUrl")
	}

	fileLogo := viewmodel.FileVm{
		ID:   model.Logo,
		Name: model.LogoName.String,
		Path: path,
	}

	return viewmodel.ContactVm{
		ID:                   model.ID,
		BranchName:           model.BranchName.String,
		TravelAgentName:      model.TravelAgentName.String,
		Address:              model.Address.String,
		Longitude:            model.Longitude.String,
		Latitude:             model.Latitude.String,
		AreaCode:             model.AreaCode,
		PhoneNumber:          model.PhoneNumber,
		SKNumber:             model.SKNumber.String,
		SKDate:               model.SKDate.String,
		Accreditation:        model.Accreditation.String,
		AccreditationDate:    model.AccreditationDate.String,
		DirectorName:         model.DirectorName.String,
		DirectorContact:      model.DirectorContact.String,
		PicName:              model.PicName,
		PicContact:           model.PicContact,
		FileLogo:             fileLogo,
		VirtualAccountNumber: model.VirtualAccountNumber.String,
		AccountNumber:        model.AccountNumber,
		AccountName:          model.AccountName,
		AccountBankName:      model.AccountBankName,
		AccountBankCode:      model.AccountBankCode,
		Email:                model.Email,
		IsZakatPartner:       model.IsZakatPartner,
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
		DeletedAt:            model.DeletedAt.String,
	}
}
