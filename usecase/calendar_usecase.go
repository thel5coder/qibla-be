package usecase

import (
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/datetime"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type CalendarUseCase struct {
	*UcContract
}

func (uc CalendarUseCase) BrowseByYearMonth(yearMonth string) (res []viewmodel.CalendarVm, err error) {
	repository := actions.NewCalendarRepository(uc.DB)
	calendars, err := repository.BrowseByYearMonth(yearMonth)
	if err != nil {
		return res, err
	}

	for _, calendar := range calendars {
		res = append(res, viewmodel.CalendarVm{
			ID:          calendar.ID,
			Title:       calendar.Title,
			Start:       calendar.Start,
			End:         calendar.End,
			Description: calendar.Description,
			CreatedAt:   calendar.CreatedAt,
			UpdatedAt:   calendar.UpdatedAt,
			DeletedAt:   calendar.DeletedAt.String,
		})
	}

	return res, err
}

func (uc CalendarUseCase) readBy(column, value string) (res viewmodel.CalendarVm, err error) {
	repository := actions.NewCalendarRepository(uc.DB)
	calendar, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.CalendarVm{
		ID:          calendar.ID,
		Title:       calendar.Title,
		Start:       calendar.Start,
		End:         calendar.End,
		Description: calendar.Description,
		CreatedAt:   calendar.CreatedAt,
		UpdatedAt:   calendar.UpdatedAt,
		DeletedAt:   calendar.DeletedAt.String,
	}

	return res, err
}

func (uc CalendarUseCase) ReadByPk(ID string) (res viewmodel.CalendarVm, err error) {
	res, err = uc.readBy("id", ID)
	if err != nil {
		return res, err
	}
	return res, err
}

func (uc CalendarUseCase) Edit(ID string, input *requests.CalendarRequest) (err error) {
	repository := actions.NewCalendarRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	body := viewmodel.CalendarVm{
		ID:          ID,
		Title:       input.Title,
		Start:       datetime.StrParseToTime(input.Start,"2006-01-02 15:04:05").Format(time.RFC3339),
		End:         datetime.StrParseToTime(input.End,"2006-01-02 15:04:05").Format(time.RFC3339),
		Description: input.Description,
		UpdatedAt:   now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc CalendarUseCase) Add(input *requests.CalendarRequest) (err error) {
	repository := actions.NewCalendarRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	body := viewmodel.CalendarVm{
		Title:       input.Title,
		Start:       datetime.StrParseToTime(input.Start,"2006-01-02 15:04:05").Format(time.RFC3339),
		End:         datetime.StrParseToTime(input.End,"2006-01-02 15:04:05").Format(time.RFC3339),
		Description: input.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	_, err = repository.Add(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc CalendarUseCase) Delete(ID string) (err error) {
	repository := actions.NewCalendarRepository(uc.DB)
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

func (uc CalendarUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewCalendarRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, nil
}
