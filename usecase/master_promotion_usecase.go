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

type MasterPromotionUseCase struct {
	*UcContract
}

//browse
func (uc MasterPromotionUseCase) Browse(filters map[string]interface{}, order, sort string, page, limit int) (res []viewmodel.MasterPromotionVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewMasterPromotionRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	masterPromotions, count, err := repository.Browse(filters, order, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterPromotion-browse")
		return res, pagination, err
	}

	for _, masterPromotion := range masterPromotions {
		res = append(res, uc.buildBody(masterPromotion))
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

//browse all
func (uc MasterPromotionUseCase) BrowseAll() (res []viewmodel.MasterPromotionVm,err error){
	repository := actions.NewMasterPromotionRepository(uc.DB)

	masterPromotions,err := repository.BrowseAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterPromotion-browseAll")
		return res,err
	}

	for _, masterPromotion := range masterPromotions {
		res = append(res, uc.buildBody(masterPromotion))
	}

	return res,err
}

//read by
func (uc MasterPromotionUseCase) ReadBy(column, value string) (res viewmodel.MasterPromotionVm, err error) {
	repository := actions.NewMasterPromotionRepository(uc.DB)

	masterPromotion, err := repository.ReadBy(column, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterPromotion-readBy")
		return res, err
	}
	res = uc.buildBody(masterPromotion)

	return res, err
}

//read by pk
func (uc MasterPromotionUseCase) ReadByPk(ID string) (res viewmodel.MasterPromotionVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-masterPromotion-readBy")
		return res, err
	}

	return res, err
}

//edit
func (uc MasterPromotionUseCase) Edit(ID string, input *requests.MasterPromotionRequest) (err error) {
	repository := actions.NewMasterPromotionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy(ID, "slug", slug.Make(input.PackageName))
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-masterPromotion-countBy")
		return err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel,messages.DataAlreadyExist,functioncaller.PrintFuncName(),"uc-masterPromotion-countByDataExist")
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.MasterPromotionVm{
		ID:        ID,
		Slug:      slug.Make(input.PackageName),
		Name:      input.PackageName,
		IsActive:  input.IsActive,
		UpdatedAt: now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterPromotion-edit")
		return err
	}

	return nil
}

//add
func (uc MasterPromotionUseCase) Add(input *requests.MasterPromotionRequest) (err error) {
	repository := actions.NewMasterPromotionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "slug", slug.Make(input.PackageName))
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-masterPromotion-countBy")
		return err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel,messages.DataAlreadyExist,functioncaller.PrintFuncName(),"uc-masterPromotion-countByDataExist")
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.MasterPromotionVm{
		Slug:      slug.Make(input.PackageName),
		Name:      input.PackageName,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err = repository.Add(body)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterPromotion-add")
		return err
	}

	return nil
}

//delete
func (uc MasterPromotionUseCase) Delete(ID string) (err error) {
	repository := actions.NewMasterPromotionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "id", ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-masterPromotion-countById")
		return err
	}
	if count > 0 {
		_, err = repository.Delete(ID, now, now)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterPromotion-delete")
			return err
		}
	}

	return nil
}

//count by
func (uc MasterPromotionUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewMasterPromotionRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-masterPromotion-countBy")
		return res, err
	}

	return res, err
}

//build body
func(uc MasterPromotionUseCase) buildBody(model models.MasterPromotion) viewmodel.MasterPromotionVm{
	return viewmodel.MasterPromotionVm{
		ID:        model.ID,
		Slug:      model.Slug,
		Name:      model.Name,
		IsActive:  model.IsActive,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
