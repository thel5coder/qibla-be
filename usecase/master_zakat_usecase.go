package usecase

import (
	"errors"
	"github.com/gosimple/slug"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type MasterZakatUseCase struct {
	*UcContract
}

//browse
func (uc MasterZakatUseCase) Browse(search map[string]interface{}, order, sort string, page, limit int) (res []viewmodel.MasterZakatVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	masterZakats, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterZakat-browse")
		return res, pagination, err
	}

	for _, masterZakat := range masterZakats {
		res = append(res, uc.buildBody(masterZakat))
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

//browse all
func (uc MasterZakatUseCase) BrowseAll() (res []viewmodel.MasterZakatVm, err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	masterProducts, err := repository.BrowseAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterZakat-browseAll")
		return res, err
	}

	for _, masterZakat := range masterProducts {
		res = append(res, uc.buildBody(masterZakat))
	}

	return res, err
}

//read by
func (uc MasterZakatUseCase) ReadBy(column, value string) (res viewmodel.MasterZakatVm, err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	masterZakat, err := repository.ReadBy(column, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterZakat-readBy")
		return res, err
	}
	res = uc.buildBody(masterZakat)

	return res, err
}

//read by pk
func (uc MasterZakatUseCase) ReadByPk(ID string) (res viewmodel.MasterZakatVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-masterZakat-readBy")
		return res, err
	}

	return res, err
}

//edit
func (uc MasterZakatUseCase) Edit(ID string, input *requests.MasterZakatRequest) (err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy(ID, "slug", slug.Make(input.Name))
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-masterZakat-countBySlug")
		return err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-masterZakat-countBySlugDataExist")
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.MasterZakatVm{
		ID:               ID,
		Slug:             slug.Make(input.Name),
		TypeZakat:        input.TypeZakat,
		Name:             input.Name,
		Description:      input.Description,
		Amount:           input.Amount,
		CurrentGoldPrice: input.CurrentGoldPrice,
		GoldNishab:       input.GoldNishab,
		UpdatedAt:        now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterZakat-edit")
		return err
	}

	return nil
}

//add
func (uc MasterZakatUseCase) Add(input *requests.MasterZakatRequest) (err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "slug", slug.Make(input.Name))
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-masterZakat-countBySlug")
		return err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-masterZakat-countBySlugDataExist")
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.MasterZakatVm{
		Slug:             slug.Make(input.Name),
		TypeZakat:        input.TypeZakat,
		Name:             input.Name,
		Description:      input.Description,
		Amount:           input.Amount,
		CurrentGoldPrice: input.CurrentGoldPrice,
		GoldNishab:       input.GoldNishab,
		CreatedAt:        now,
		UpdatedAt:        now,
	}
	_, err = repository.Add(body)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterZakat-add")
		return err
	}

	return nil
}

//delete
func (uc MasterZakatUseCase) Delete(ID string) (err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "id", ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-masterZakat-countBy")
		return err
	}
	if count > 0 {
		_, err = repository.Delete(ID, now, now)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterZakat-delete")
			return err
		}
	}

	return nil
}

//count by
func (uc MasterZakatUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewMasterZakatRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterZakat-countBy")
		return res, err
	}

	return res, err
}

//build body
func(uc MasterZakatUseCase) buildBody(model models.MasterZakat) viewmodel.MasterZakatVm{
	return viewmodel.MasterZakatVm{
		ID:               model.ID,
		Slug:             model.Slug,
		TypeZakat:        model.TypeZakat,
		Name:             model.Name,
		Description:      model.Description,
		Amount:           model.Amount.Int32,
		CurrentGoldPrice: model.CurrentGoldPrice.Int32,
		GoldNishab:       model.GoldNishab.Int32,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
	}
}
