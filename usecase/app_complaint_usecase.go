package usecase

import (
	"qibla-backend/db/repositories/actions"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"strconv"
	"time"
)

type AppComplaintUseCase struct {
	*UcContract
}

func (uc AppComplaintUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.AppComplaintVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewAppComplaintRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	appComplaints, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, appComplaint := range appComplaints {
		res = append(res, viewmodel.AppComplaintVm{
			ID:            appComplaint.ID,
			FullName:      appComplaint.FullName,
			Email:         appComplaint.Email,
			TicketNumber:  appComplaint.TicketNumber,
			ComplaintType: appComplaint.ComplaintType,
			Complaint:     appComplaint.Complaint,
			Solution:      appComplaint.Solution.String,
			Status:        appComplaint.Status,
			CreatedAt:     appComplaint.CreatedAt,
			UpdatedAt:     appComplaint.UpdatedAt,
			DeletedAt:     appComplaint.DeletedAt.String,
		})
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc AppComplaintUseCase) readBy(column, value string) (res viewmodel.AppComplaintVm, err error) {
	repository := actions.NewAppComplaintRepository(uc.DB)
	appComplaint, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.AppComplaintVm{
		ID:            appComplaint.ID,
		FullName:      appComplaint.FullName,
		Email:         appComplaint.Email,
		TicketNumber:  appComplaint.TicketNumber,
		ComplaintType: appComplaint.ComplaintType,
		Complaint:     appComplaint.Complaint,
		Solution:      appComplaint.Solution.String,
		Status:        appComplaint.Status,
		CreatedAt:     appComplaint.CreatedAt,
		UpdatedAt:     appComplaint.UpdatedAt,
		DeletedAt:     appComplaint.DeletedAt.String,
	}

	return res, err
}

func (uc AppComplaintUseCase) Edit(ID string, input *requests.AppComplaintRequest) (err error) {
	repository := actions.NewAppComplaintRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	body := viewmodel.AppComplaintVm{
		ID:        ID,
		Complaint: input.Complaint,
		Solution:  input.Solution,
		Status:    input.Status,
		UpdatedAt: now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc AppComplaintUseCase) Add(input *requests.AppComplaintRequest) (err error) {
	repository := actions.NewAppComplaintRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	body := viewmodel.AppComplaintVm{
		FullName:      input.FullName,
		Email:         input.Email,
		TicketNumber:  input.TicketNumber,
		ComplaintType: input.ComplaintType,
		Complaint:     input.Complaint,
		Solution:      input.Solution,
		Status:        "open",
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	_, err = repository.Add(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc AppComplaintUseCase) Delete(ID string) (err error) {
	repository := actions.NewAppComplaintRepository(uc.DB)
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

func (uc AppComplaintUseCase) ReadByPk(ID string) (res viewmodel.AppComplaintVm, err error) {
	res, err = uc.readBy("id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc AppComplaintUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewAppComplaintRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc AppComplaintUseCase) GetTicketNumber() (res string, err error) {
	repository := actions.NewAppComplaintRepository(uc.DB)
	count, err := repository.CountAll()
	if err != nil {
		return res, err
	}
	count = count + 1

	if count > 99999 {
		res = strconv.Itoa(count)
	} else if count > 9999 {
		res = `0` + strconv.Itoa(count)
	} else if count > 999 {
		res = `00` + strconv.Itoa(count)
	} else if count > 99 {
		res = `000` + strconv.Itoa(count)
	} else if count > 9 {
		res = `0000` + strconv.Itoa(count)
	} else {
		res = `00000` + strconv.Itoa(count)
	}

	return res, err
}
