package usecase

import (
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/pusher"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type CrmBoardUseCase struct {
	*UcContract
}

func (uc CrmBoardUseCase) BrowseByCrmStoryID(crmStoryID string) (res []viewmodel.CrmBoardVm, err error) {
	repository := actions.NewCrmBoardRepository(uc.DB)
	crmBoards, err := repository.BrowseByCrmStoryID(crmStoryID)
	if err != nil {
		return res, err
	}

	for _, crmBoard := range crmBoards {
		res = append(res, viewmodel.CrmBoardVm{
			ID:                crmBoard.ID,
			CrmStoryID:        crmBoard.CrmStoryID,
			ContactID:         crmBoard.ContactID,
			Opportunity:       crmBoard.Opportunity,
			ProfitExpectation: crmBoard.ProfitExpectation,
			Star:              crmBoard.Star,
			CreatedAt:         crmBoard.CreatedAt,
			UpdatedAt:         crmBoard.UpdatedAt,
		})
	}

	return res, err
}

func (uc CrmBoardUseCase) ReadBy(column, value string) (res viewmodel.CrmBoardVm, err error) {
	repository := actions.NewCrmBoardRepository(uc.DB)
	crmBoard, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.CrmBoardVm{
		ID:                crmBoard.ID,
		CrmStoryID:        crmBoard.CrmStoryID,
		ContactID:         crmBoard.ContactID,
		Opportunity:       crmBoard.Opportunity,
		ProfitExpectation: crmBoard.ProfitExpectation,
		Star:              crmBoard.Star,
		CreatedAt:         crmBoard.CreatedAt,
		UpdatedAt:         crmBoard.UpdatedAt,
	}

	return res, err
}

func (uc CrmBoardUseCase) Edit(ID string, input *requests.CrmBoardRequest) (err error) {
	repository := actions.NewCrmBoardRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	body := viewmodel.CrmBoardVm{
		ID:                ID,
		CrmStoryID:        input.CrmStoryID,
		ContactID:         input.ContactID,
		Opportunity:       input.Opportunity,
		ProfitExpectation: input.ProfitExpectation,
		Star:              input.Star,
		UpdatedAt:         now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		return err
	}

	return nil
}

func (uc CrmBoardUseCase) EditBoardStory(ID, crmStoryID string) (err error) {
	repository := actions.NewCrmBoardRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	crmBoard,err := uc.ReadBy("id",ID)
	if err != nil {
		return err
	}
	oldStoryID := crmBoard.CrmStoryID
	body := map[string]interface{}{
		"old_story":oldStoryID,
		"new_story":crmStoryID,
	}

	_, err = repository.EditBoardStory(ID, crmStoryID, now)
	if err != nil {
		return err
	}

	pusherUc := PusherUseCase{UcContract:uc.UcContract}
	pusherUc.Broadcast(pusher.EventStory,body)

	return nil
}

func (uc CrmBoardUseCase) Add(input *requests.CrmBoardRequest) (err error) {
	repository := actions.NewCrmBoardRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	body := viewmodel.CrmBoardVm{
		CrmStoryID:        input.CrmStoryID,
		ContactID:         input.ContactID,
		Opportunity:       input.Opportunity,
		ProfitExpectation: input.ProfitExpectation,
		Star:              input.Star,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
	_, err = repository.Add(body)
	if err != nil {
		return err
	}

	pusherBody := map[string]interface{}{
		"old_story":"",
		"new_story":input.CrmStoryID,
	}
	pusherUc := PusherUseCase{UcContract:uc.UcContract}
	pusherUc.Broadcast(pusher.EventStory,pusherBody)


	return nil
}

func (uc CrmBoardUseCase) CountBy(ID, crmStoryID, column, value string) (res int, err error) {
	repository := actions.NewCrmBoardRepository(uc.DB)
	res, err = repository.CountBy(ID, crmStoryID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}
