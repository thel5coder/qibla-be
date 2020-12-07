package usecase

import (
	"errors"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type TermConditionUseCase struct {
	*UcContract
}

func (uc TermConditionUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.TermConditionVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewTermConditionRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	termConditions, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, termCondition := range termConditions {
		res = append(res, viewmodel.TermConditionVm{
			ID:          termCondition.ID,
			TermName:    termCondition.TermName,
			TermType:    termCondition.TermType,
			Description: termCondition.Description,
			CreatedAt:   termCondition.CreatedAt,
			UpdatedAt:   termCondition.UpdatedAt,
			DeletedAt:   termCondition.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc TermConditionUseCase) ReadBy(column, value string) (res viewmodel.TermConditionVm, err error) {
	repository := actions.NewTermConditionRepository(uc.DB)
	termCondition, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.TermConditionVm{
		ID:          termCondition.ID,
		TermName:    termCondition.TermName,
		TermType:    termCondition.TermType,
		Description: termCondition.Description,
		CreatedAt:   termCondition.CreatedAt,
		UpdatedAt:   termCondition.UpdatedAt,
		DeletedAt:   termCondition.DeletedAt.String,
	}

	return res, err
}

func (uc TermConditionUseCase) ReadByPk(ID string) (res viewmodel.TermConditionVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		return res, errors.New(messages.DataNotFound)
	}

	return res, err
}

func (uc TermConditionUseCase) Edit(ID string, input *requests.TermConditionRequest) (err error) {
	repository := actions.NewTermConditionRepository(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsNameExist(ID, input.TermName)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.TermConditionVm{
		ID:          ID,
		TermName:    input.TermName,
		TermType:    input.TermType,
		Description: input.Description,
		UpdatedAt:   now.Format(time.RFC3339),
	}
	_, err = repository.Edit(body)

	return err
}

func (uc TermConditionUseCase) Add(input *requests.TermConditionRequest) (err error) {
	repository := actions.NewTermConditionRepository(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsNameExist("", input.TermName)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.TermConditionVm{
		TermName:    input.TermName,
		TermType:    input.TermType,
		Description: input.Description,
		CreatedAt:   now.Format(time.RFC3339),
		UpdatedAt:   now.Format(time.RFC3339),
	}
	_, err = repository.Add(body)

	return err
}

func (uc TermConditionUseCase) Delete(ID string) (err error) {
	repository := actions.NewTermConditionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountByPk(ID)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New(messages.DataNotFound)
	}

	_, err = repository.Delete(ID, now, now)

	return err
}

func (uc TermConditionUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewTermConditionRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc TermConditionUseCase) CountByPk(ID string) (res int, err error) {
	repository := actions.NewTermConditionRepository(uc.DB)
	res, err = repository.CountByPk(ID)

	return res, err
}

func (uc TermConditionUseCase) IsNameExist(ID, name string) (res bool, err error) {
	count, err := uc.CountBy(ID, "term_name", name)
	if err != nil {
		return res, err
	}

	return count > 0, err
}
