package usecase

import (
	"qibla-backend/db/repositories/actions"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type PrayUseCase struct {
	*UcContract
}

func (uc PrayUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.PrayVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewPrayRepository(uc.DB)
	fileUc := FileUseCase{UcContract: uc.UcContract}
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	prays, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, pray := range prays {
		file, _ := fileUc.ReadByPk(pray.FileID)
		res = append(res, viewmodel.PrayVm{
			ID:        pray.ID,
			Name:      pray.Name,
			FileID:    pray.FileID,
			IsActive:  pray.IsActive,
			File:      file,
			CreatedAt: pray.CreatedAt,
			UpdatedAt: pray.UpdatedAt,
			DeletedAt: pray.DeletedAt.String,
		})
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc PrayUseCase) ReadBy(column, value string) (res viewmodel.PrayVm, err error) {
	repository := actions.NewPrayRepository(uc.DB)
	fileUc := FileUseCase{UcContract: uc.UcContract}
	pray, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	file, _ := fileUc.ReadByPk(pray.FileID)
	res = viewmodel.PrayVm{
		ID:        pray.ID,
		Name:      pray.Name,
		FileID:    pray.FileID,
		IsActive:  pray.IsActive,
		File:      file,
		CreatedAt: pray.CreatedAt,
		UpdatedAt: pray.UpdatedAt,
		DeletedAt: pray.DeletedAt.String,
	}

	return res, err
}

func (uc PrayUseCase) ReadByPk(ID string) (res viewmodel.PrayVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc PrayUseCase) Edit(ID string, input *requests.PrayRequest) (err error) {
	repository := actions.NewPrayRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	body := viewmodel.PrayVm{
		ID:        ID,
		Name:      input.Name,
		FileID:    input.FileID,
		IsActive:  input.IsActive,
		UpdatedAt: now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc PrayUseCase) Add(input *requests.PrayRequest) (err error) {
	repository := actions.NewPrayRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	body := viewmodel.PrayVm{
		Name:      input.Name,
		FileID:    input.FileID,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_, err = repository.Add(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc PrayUseCase) Delete(ID string) (err error) {
	repository := actions.NewPrayRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "id", ID)
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

func (uc PrayUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewPrayRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}
