package usecase

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type SatisfactionCategoryUseCase struct {
	*UcContract
}

func (uc SatisfactionCategoryUseCase) BrowseAllBy(column, value string) (res []viewmodel.SatisfactionCategoryVm, err error) {
	repository := actions.NewSatisfactionCategoryRepository(uc.DB)
	satisfactionCategories, err := repository.BrowseAllBy(column, value)
	if err != nil {
		return res, err
	}

	for _, satisfactionCategory := range satisfactionCategories {
		res = append(res, viewmodel.SatisfactionCategoryVm{
			ID:          satisfactionCategory.ID,
			ParentID:    satisfactionCategory.ParentID.String,
			Slug:        satisfactionCategory.Slug,
			Name:        satisfactionCategory.Name,
			Description: satisfactionCategory.Description.String,
			IsActive:    satisfactionCategory.IsActive,
			CreatedAt:   satisfactionCategory.CreatedAt,
			UpdatedAt:   satisfactionCategory.UpdatedAt,
			DeletedAt:   satisfactionCategory.DeletedAt.String,
			Child:       nil,
		})
	}

	return res, err
}

func (uc SatisfactionCategoryUseCase) BrowseParent() (res []viewmodel.SatisfactionCategoryVm, err error) {
	res, err = uc.BrowseAllBy("parent_id", "")
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc SatisfactionCategoryUseCase) BrowseChild(parentID string) (res []viewmodel.SatisfactionCategoryVm) {
	satisfactionCategories, err := uc.BrowseAllBy("parent_id", parentID)
	if err != nil {
		return res
	}

	for _, satisfactionCategory := range satisfactionCategories {
		res = append(res, viewmodel.SatisfactionCategoryVm{
			ID:          satisfactionCategory.ID,
			ParentID:    satisfactionCategory.ParentID,
			Slug:        satisfactionCategory.Slug,
			Name:        satisfactionCategory.Name,
			Description: satisfactionCategory.Description,
			IsActive:    satisfactionCategory.IsActive,
			CreatedAt:   satisfactionCategory.CreatedAt,
			UpdatedAt:   satisfactionCategory.UpdatedAt,
			DeletedAt:   satisfactionCategory.DeletedAt,
			Child:       uc.BrowseChild(satisfactionCategory.ID),
		})
	}

	return res
}

func (uc SatisfactionCategoryUseCase) BrowseAllTree() (res []viewmodel.SatisfactionCategoryVm, err error) {
	parents, err := uc.BrowseParent()
	if err != nil {
		return res, err
	}

	for _, parent := range parents {
		res = append(res, viewmodel.SatisfactionCategoryVm{
			ID:          parent.ID,
			ParentID:    parent.ParentID,
			Slug:        parent.Slug,
			Name:        parent.Name,
			Description: parent.Description,
			IsActive:    parent.IsActive,
			CreatedAt:   parent.CreatedAt,
			UpdatedAt:   parent.UpdatedAt,
			DeletedAt:   parent.DeletedAt,
			Child:       uc.BrowseChild(parent.ID),
		})
	}

	return res, err
}

func (uc SatisfactionCategoryUseCase) ReadBy(column, value string) (res viewmodel.SatisfactionCategoryVm, err error) {
	repository := actions.NewSatisfactionCategoryRepository(uc.DB)
	satisfactionCategory, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.SatisfactionCategoryVm{
		ID:          satisfactionCategory.ID,
		ParentID:    satisfactionCategory.ParentID.String,
		Slug:        satisfactionCategory.Slug,
		Name:        satisfactionCategory.Name,
		Description: satisfactionCategory.Description.String,
		IsActive:    satisfactionCategory.IsActive,
		CreatedAt:   satisfactionCategory.CreatedAt,
		UpdatedAt:   satisfactionCategory.UpdatedAt,
		DeletedAt:   satisfactionCategory.DeletedAt.String,
		Child:       uc.BrowseChild(satisfactionCategory.ID),
	}

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
		fmt.Println(1)
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
	count, err := uc.CountBy("", "id", ID)
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
