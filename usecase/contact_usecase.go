package usecase

import (
	"errors"
	"fmt"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type ContactUseCase struct {
	*UcContract
}

func (uc ContactUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.ContactVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewContactRepository(uc.DB)
	fileUc := FileUseCase{}
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	contacts, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, contact := range contacts {
		file, _ := fileUc.ReadByPk(contact.Logo)
		res = append(res, viewmodel.ContactVm{
			ID:                   contact.ID,
			BranchName:           contact.BranchName,
			TravelAgentName:      contact.TravelAgentName,
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
			CreatedAt:            contact.CreatedAt,
			UpdatedAt:            contact.UpdatedAt,
			DeletedAt:            contact.DeletedAt.String,
		})
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc ContactUseCase) BrowseAll(search string) (res []viewmodel.ContactVm, err error) {
	repository := actions.NewContactRepository(uc.DB)
	fileUc := FileUseCase{}
	contacts, err := repository.BrowseAll(search)
	if err != nil {
		return res, err
	}

	for _, contact := range contacts {
		file, _ := fileUc.ReadByPk(contact.Logo)
		res = append(res, viewmodel.ContactVm{
			ID:                   contact.ID,
			BranchName:           contact.BranchName,
			TravelAgentName:      contact.TravelAgentName,
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
			CreatedAt:            contact.CreatedAt,
			UpdatedAt:            contact.UpdatedAt,
			DeletedAt:            contact.DeletedAt.String,
		})
	}

	return res, err
}

func (uc ContactUseCase) ReadBy(column, value string) (res viewmodel.ContactVm, err error) {
	repository := actions.NewContactRepository(uc.DB)
	fileUc := FileUseCase{UcContract: uc.UcContract}
	contact, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	file, _ := fileUc.ReadByPk(contact.Logo)

	res = viewmodel.ContactVm{
		ID:                   contact.ID,
		BranchName:           contact.BranchName,
		TravelAgentName:      contact.TravelAgentName,
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
		CreatedAt:            contact.CreatedAt,
		UpdatedAt:            contact.UpdatedAt,
		DeletedAt:            contact.DeletedAt.String,
	}

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
		ID:                   ID,
		BranchName:           input.BranchName,
		TravelAgentName:      input.TravelAgentName,
		Address:              input.Address,
		Longitude:            input.Longitude,
		Latitude:             input.Latitude,
		AreaCode:             input.AreaCode,
		PhoneNumber:          input.PhoneNumber,
		SKNumber:             input.SKNumber,
		SKDate:               input.SKDate,
		Accreditation:        input.Accreditation,
		AccreditationDate:    input.AccreditationDate,
		DirectorName:         input.DirectorName,
		DirectorContact:      input.DirectorContact,
		PicName:              input.PicName,
		PicContact:           input.PicContact,
		Logo:                 input.Logo,
		VirtualAccountNumber: input.VirtualAccountNumber,
		AccountNumber:        input.AccountNumber,
		AccountName:          input.AccountName,
		AccountBankName:      input.AccountBankName,
		AccountBankCode:      input.AccountBankCode,
		Email:                input.Email,
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
		BranchName:           input.BranchName,
		TravelAgentName:      input.TravelAgentName,
		Address:              input.Address,
		Longitude:            input.Longitude,
		Latitude:             input.Latitude,
		AreaCode:             input.AreaCode,
		PhoneNumber:          input.PhoneNumber,
		SKNumber:             input.SKNumber,
		SKDate:               input.SKDate,
		Accreditation:        input.Accreditation,
		AccreditationDate:    input.AccreditationDate,
		DirectorName:         input.DirectorName,
		DirectorContact:      input.DirectorContact,
		PicName:              input.PicName,
		PicContact:           input.PicContact,
		Logo:                 input.Logo,
		VirtualAccountNumber: input.VirtualAccountNumber,
		AccountNumber:        input.AccountNumber,
		AccountName:          input.AccountName,
		AccountBankName:      input.AccountBankName,
		AccountBankCode:      input.AccountBankCode,
		Email:                input.Email,
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
		fmt.Println(1)
		return err
	}
	if skCount > 0 {
		return errors.New("nomor sk sudah ada")
	}

	//cek akreditasi
	accreditationCount, err := uc.countBy(ID, "accreditation", input.Accreditation)
	if err != nil {
		fmt.Println(1)
		return err
	}
	if accreditationCount > 0 {
		return errors.New("akreditasi sudah ada")
	}

	return nil
}

func (uc ContactUseCase) ReadByPk(ID string) (res viewmodel.ContactVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
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
