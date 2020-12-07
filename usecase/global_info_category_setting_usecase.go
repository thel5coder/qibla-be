package usecase

import (
	"errors"
	"fmt"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type GlobalInfoCategorySettingUseCase struct {
	*UcContract
}

func (uc GlobalInfoCategorySettingUseCase) Browse(globalInfoCategorySlug,search, order, sort string, page, limit int) (res []viewmodel.GlobalInfoCategorySettingVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewGlobalInfoCategorySettingRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	globalInfoCategorySettings, count, err := repository.Browse(globalInfoCategorySlug,search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, globalInfoCategory := range globalInfoCategorySettings {
		res = append(res, viewmodel.GlobalInfoCategorySettingVm{
			ID:                     globalInfoCategory.ID,
			GlobalInfoCategoryID:   globalInfoCategory.GlobalInfoCategoryID,
			GlobalInfoCategoryName: globalInfoCategory.GlobalInfoCategoryName,
			Description:            globalInfoCategory.Description,
			IsActive:               globalInfoCategory.IsActive,
			CreatedAt:              globalInfoCategory.CreatedAt,
			UpdatedAt:              globalInfoCategory.UpdatedAt,
			DeletedAt:              globalInfoCategory.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc GlobalInfoCategorySettingUseCase) ReadBy(column, value string) (res viewmodel.GlobalInfoCategorySettingVm, err error) {
	repository := actions.NewGlobalInfoCategorySettingRepository(uc.DB)
	globalInfoCategory, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.GlobalInfoCategorySettingVm{
		ID:                     globalInfoCategory.ID,
		GlobalInfoCategoryID:   globalInfoCategory.GlobalInfoCategoryID,
		GlobalInfoCategoryName: globalInfoCategory.GlobalInfoCategoryName,
		Description:            globalInfoCategory.Description,
		IsActive:               globalInfoCategory.IsActive,
		CreatedAt:              globalInfoCategory.CreatedAt,
		UpdatedAt:              globalInfoCategory.UpdatedAt,
		DeletedAt:              globalInfoCategory.DeletedAt.String,
	}

	return res, err
}

func (uc GlobalInfoCategorySettingUseCase) ReadByPk(ID string) (res viewmodel.GlobalInfoCategorySettingVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		fmt.Println(err)
		return res, errors.New(messages.DataNotFound)
	}

	return res, err
}

func (uc GlobalInfoCategorySettingUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewGlobalInfoCategorySettingRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc GlobalInfoCategorySettingUseCase) Edit(ID string, input *requests.GlobalInfoCategorySettingRequest) (err error) {
	repository := actions.NewGlobalInfoCategorySettingRepository(uc.DB)
	now := time.Now().UTC()

	//isExist, err := uc.IsGlobalInfoCategoryExist(ID, input.GlobalInfoCategoryID)
	//if err != nil {
	//	return err
	//}
	//if isExist {
	//	return errors.New(messages.DataAlreadyExist)
	//}

	body := viewmodel.GlobalInfoCategorySettingVm{
		ID:                   ID,
		GlobalInfoCategoryID: input.GlobalInfoCategoryID,
		Description:          input.Description,
		IsActive:             input.IsActive,
		UpdatedAt:            now.Format(time.RFC3339),
	}
	_, err = repository.Edit(body)

	return err
}

func (uc GlobalInfoCategorySettingUseCase) Add(input *requests.GlobalInfoCategorySettingRequest) (err error) {
	repository := actions.NewGlobalInfoCategorySettingRepository(uc.DB)
	now := time.Now().UTC()

	//isExist, err := uc.IsGlobalInfoCategoryExist("", input.GlobalInfoCategoryID)
	//if err != nil {
	//	return err
	//}
	//if isExist {
	//	return errors.New(messages.DataAlreadyExist)
	//}

	body := viewmodel.GlobalInfoCategorySettingVm{
		GlobalInfoCategoryID: input.GlobalInfoCategoryID,
		Description:          input.Description,
		IsActive:             true,
		CreatedAt:            now.Format(time.RFC3339),
		UpdatedAt:            now.Format(time.RFC3339),
	}
	_, err = repository.Add(body)

	return err
}

func (uc GlobalInfoCategorySettingUseCase) Delete(ID string) (err error) {
	repository := actions.NewGlobalInfoCategorySettingRepository(uc.DB)
	now := time.Now().UTC()

	count, err := uc.CountByPk(ID)
	if err != nil {
		return errors.New(messages.DataNotFound)
	}

	if count > 0 {
		_, err = repository.Delete(ID, now.Format(time.RFC3339), now.Format(time.RFC3339))
	}

	return err
}

func (uc GlobalInfoCategorySettingUseCase) CountByPk(ID string) (res int, err error) {
	repository := actions.NewGlobalInfoCategorySettingRepository(uc.DB)
	res, err = repository.CountBy("", "id", ID)

	return res, err
}

func (uc GlobalInfoCategorySettingUseCase) IsGlobalInfoCategoryExist(ID, globalInfoCategoryID string) (res bool, err error) {
	count, err := uc.CountBy(ID, "global_info_category_id", globalInfoCategoryID)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
