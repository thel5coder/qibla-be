package usecase

import (
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
)

type CalendarParticipantUseCase struct {
	*UcContract
}

func (uc CalendarParticipantUseCase) BrowseByCalendarID(calendarID string) (res []viewmodel.CalendarParticipantVm, err error) {
	repository := actions.NewCalendarParticipantRepository(uc.DB)
	calendarParticipants, err := repository.BrowseByCalendarID(calendarID)
	if err != nil {
		return res, err
	}

	for _, calendarParticipant := range calendarParticipants {
		res = append(res, viewmodel.CalendarParticipantVm{
			ID:         calendarParticipant.ID,
			CalendarID: calendarParticipant.CalendarID,
			Email:      calendarParticipant.Email,
		})
	}

	return res, err
}

func (uc CalendarParticipantUseCase) Add(calendarID, emailParticipant string) (err error) {
	repository := actions.NewCalendarParticipantRepository(uc.DB)
	err = repository.Add(calendarID, emailParticipant, uc.TX)

	return err
}

func (uc CalendarParticipantUseCase) Delete(calendarID string) (err error) {
	repository := actions.NewCalendarParticipantRepository(uc.DB)
	err = repository.Delete(calendarID, uc.TX)

	return err
}

func (uc CalendarParticipantUseCase) Store(calendarID string, emailParticipants []string) (err error) {
	calendarParticipants, err := uc.BrowseByCalendarID(calendarID)
	if err != nil {
		return err
	}
	if len(calendarParticipants) > 0 {
		err = uc.Delete(calendarID)
		if err != nil {
			return err
		}
	}

	for _, emailParticipant := range emailParticipants {
		err = uc.Add(calendarID, emailParticipant)
		if err != nil {
			return err
		}
	}

	return nil
}
