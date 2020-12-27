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

type SatisfactionCategoryUseCase struct {
	*UcContract
}

func (uc SatisfactionCategoryUseCase) BrowseAllBy(filters map[string]interface{},order,sort string) (res []viewmodel.SatisfactionCategoryVm, err error) {
	repository := actions.NewSatisfactionCategoryRepository(uc.DB)
	if order == "" {
		order = "name"
	}
	if sort == ""{
		sort = "asc"
	}
	satisfactionCategories, err := repository.BrowseAllBy(filters,order,sort)
	if err != nil {
		return res, err
	}

	res = uc.buildBody(satisfactionCategories)


	return res, err
}

func (uc SatisfactionCategoryUseCase) ReadBy(column, value string) (res []viewmodel.SatisfactionCategoryVm, err error) {
	repository := actions.NewSatisfactionCategoryRepository(uc.DB)
	satisfactionCategory,err := repository.ReadBy(column,value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-satisfactionCategory-readBy")
		return res,err
	}
	res = uc.buildBody(satisfactionCategory)

	return res, err
}

func (uc SatisfactionCategoryUseCase) Edit(ID, parentID, name, description string, isActive bool) (err error) {
	repository := actions.NewSatisfactionCategoryRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountByParentID(parentID, ID, slug.Make(name))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.SatisfactionCategoryVm{
		ID:          ID,
		ParentID:    parentID,
		Slug:        slug.Make(name),
		Name:        name,
		Description: description,
		IsActive:    isActive,
		UpdatedAt:   now,
	}
	err = repository.Edit(body, uc.TX)

	return err
}

//edit updated at
func (uc SatisfactionCategoryUseCase) EditUpdatedAt(ID string) (err error) {
	repository := actions.NewSatisfactionCategoryRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	model := models.SatisfactionCategory{
		ID:        ID,
		UpdatedAt: now,
	}
	err = repository.EditUpdatedAt(model, uc.TX)
	if err != nil {
		return err
	}

	return nil
}

func (uc SatisfactionCategoryUseCase) Add(parentID, name, description string, isActive bool) (err error) {
	repository := actions.NewSatisfactionCategoryRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountByParentID(parentID, "", slug.Make(name))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.SatisfactionCategoryVm{
		ParentID:    parentID,
		Slug:        slug.Make(name),
		Name:        name,
		Description: description,
		IsActive:    true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	err = repository.Add(body, uc.TX)

	return err
}

func (uc SatisfactionCategoryUseCase) Store(input *requests.SatisfactionCategoryRequest) (err error) {
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	for _, subCategory := range input.SubCategories {
		if subCategory.ID == "" {
			err = uc.Add(input.ParentID, subCategory.Name, subCategory.Description, subCategory.IsActive)
		} else {
			err = uc.Edit(subCategory.ID, input.ParentID, subCategory.Name, subCategory.Description, subCategory.IsActive)
		}

		if err != nil {
			uc.TX.Rollback()

			return err
		}
	}

	err = uc.EditUpdatedAt(input.ParentID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-satisfactionCategory-editUpdatedAtParent")
		uc.TX.Rollback()

		return err
	}
	uc.TX.Commit()

	return nil
}

func (uc SatisfactionCategoryUseCase) DeleteBy(column, value string) (err error) {
	repository := actions.NewSatisfactionCategoryRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	err = repository.DeleteBy(column, value, now, now, uc.TX)

	return err
}

func (uc SatisfactionCategoryUseCase) DeleteByPk(ID string) (err error) {
	count, err := uc.CountBy("", "parent_id", ID)
	if err != nil {
		return err
	}

	uc.TX, err = uc.DB.Begin()
	if err != nil {

		uc.TX.Rollback()

		return err
	}

	if count > 0 {
		err = uc.DeleteBy("parent_id", ID)
		if err != nil {
			uc.TX.Rollback()

			return err
		}
	} else {
		err = uc.DeleteBy("id", ID)
		if err != nil {
			uc.TX.Rollback()

			return err
		}
	}

	uc.TX.Commit()

	return nil
}

func (uc SatisfactionCategoryUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewSatisfactionCategoryRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc SatisfactionCategoryUseCase) CountByParentID(parentID, ID, slug string) (res int, err error) {
	repository := actions.NewSatisfactionCategoryRepository(uc.DB)
	res, err = repository.CountByParentID(parentID, ID, slug)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc SatisfactionCategoryUseCase) buildBody(model []models.SatisfactionCategory) []viewmodel.SatisfactionCategoryVm {
	var parents []viewmodel.SatisfactionCategoryVm

	for _,data := range model {
		if data.ParentID.String == "" {
			parents = append(parents, viewmodel.SatisfactionCategoryVm{
				ID:          data.ID,
				ParentID:    data.ParentID.String,
				Slug:        data.Slug,
				Name:        data.Name,
				Description: data.Description.String,
				IsActive:    data.IsActive,
				CreatedAt:   data.CreatedAt,
				UpdatedAt:   data.UpdatedAt,
				Child:       nil,
			})
		}
	}

	if len(parents) > 0 {
		for i := 0; i < len(parents); i ++ {
			for _,data := range model {
				if parents[i].ID == data.ParentID.String {
					parents[i].Child = append(parents[i].Child, viewmodel.SatisfactionCategoryVm{
						ID:          data.ID,
						ParentID:    data.ParentID.String,
						Slug:        data.Slug,
						Name:        data.Name,
						Description: data.Description.String,
						IsActive:    data.IsActive,
						CreatedAt:   data.CreatedAt,
						UpdatedAt:   data.UpdatedAt,
						Child:       nil,
					})
				}
			}
		}
	}

	return parents
}
