package usecase

import (
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/datetime"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type CalendarUseCase struct {
	*UcContract
}

func (uc CalendarUseCase) BrowseByYearMonth(yearMonth string) (res []viewmodel.CalendarVm, err error) {
	repository := actions.NewCalendarRepository(uc.DB)
	calendarParticipantUc := CalendarParticipantUseCase{UcContract: uc.UcContract}
	calendars, err := repository.BrowseByYearMonth(yearMonth)
	if err != nil {
		return res, err
	}

	for _, calendar := range calendars {
		var participants []string
		calendarParticipants, _ := calendarParticipantUc.BrowseByCalendarID(calendar.ID)
		for _, calendarParticipant := range calendarParticipants {
			participants = append(participants, calendarParticipant.Email)
		}
		res = append(res, viewmodel.CalendarVm{
			ID:           calendar.ID,
			Title:        calendar.Title,
			Start:        calendar.Start,
			End:          calendar.End,
			Description:  calendar.Description,
			Participants: participants,
			Remember:     calendar.Remember,
			CreatedAt:    calendar.CreatedAt,
			UpdatedAt:    calendar.UpdatedAt,
			DeletedAt:    calendar.DeletedAt.String,
		})
	}

	return res, err
}

func (uc CalendarUseCase) readBy(column, value string) (res viewmodel.CalendarVm, err error) {
	repository := actions.NewCalendarRepository(uc.DB)
	calendarParticipantUc := CalendarParticipantUseCase{UcContract: uc.UcContract}
	var participants []string

	calendar, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	calendarParticipants, _ := calendarParticipantUc.BrowseByCalendarID(calendar.ID)
	for _, calendarParticipant := range calendarParticipants {
		participants = append(participants, calendarParticipant.Email)
	}
	res = viewmodel.CalendarVm{
		ID:           calendar.ID,
		Title:        calendar.Title,
		Start:        calendar.Start,
		End:          calendar.End,
		Description:  calendar.Description,
		Participants: participants,
		Remember:     calendar.Remember,
		CreatedAt:    calendar.CreatedAt,
		UpdatedAt:    calendar.UpdatedAt,
		DeletedAt:    calendar.DeletedAt.String,
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

	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	startDate := datetime.StrParseToTime(input.Start, "2006-01-02 15:04:05")
	endDate := datetime.StrParseToTime(input.End, "2006-01-02 15:04:05")
	loc, _ := time.LoadLocation("UTC")
	body := viewmodel.CalendarVm{
		ID:          ID,
		Title:       input.Title,
		Start:       startDate.In(loc).Format(time.RFC3339),
		End:         endDate.In(loc).Format(time.RFC3339),
		Description: input.Description,
		Remember:    input.Remember,
		UpdatedAt:   now,
	}
	err = repository.Edit(body, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	calendarParticipantUc := CalendarParticipantUseCase{UcContract: uc.UcContract}
	err = calendarParticipantUc.Store(ID, input.Participants)
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	uc.TX.Commit()

	return nil
}

func (uc CalendarUseCase) Add(input *requests.CalendarRequest) (err error) {
	repository := actions.NewCalendarRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	startDate := datetime.StrParseToTime(input.Start, "2006-01-02 15:04:05")
	endDate := datetime.StrParseToTime(input.End, "2006-01-02 15:04:05")
	loc, _ := time.LoadLocation("UTC")
	body := viewmodel.CalendarVm{
		Title:       input.Title,
		Start:       startDate.In(loc).Format(time.RFC3339),
		End:         endDate.In(loc).Format(time.RFC3339),
		Description: input.Description,
		Remember:    input.Remember,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	body.ID, err = repository.Add(body, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	calendarParticipantUc := CalendarParticipantUseCase{UcContract: uc.UcContract}
	err = calendarParticipantUc.Store(body.ID, input.Participants)
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	uc.TX.Commit()

	return nil
}

func (uc CalendarUseCase) Delete(ID string) (err error) {
	repository := actions.NewCalendarRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	count, err := uc.countBy("", "id", ID)
	if err != nil {
		return err
	}

	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	calendarParticipantUc := CalendarParticipantUseCase{UcContract: uc.UcContract}
	if count > 0 {
		err = repository.Delete(ID, now, now, uc.TX)
		if err != nil {
			uc.TX.Rollback()

			return err
		}

		err = calendarParticipantUc.Delete(ID)
		if err != nil {
			uc.TX.Rollback()

			return err
		}
	}
	uc.TX.Commit()

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
