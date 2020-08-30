package usecase

import (
	"errors"
	"github.com/gosimple/slug"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/messages"
	"qibla-backend/helpers/pusher"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type CrmStoryUseCase struct {
	*UcContract
}

func (uc CrmStoryUseCase) BrowseAll() (res []viewmodel.CrmStoryVm, err error) {
	repository := actions.NewCrmStoryRepository(uc.DB)
	crmStories, err := repository.BrowseAll()
	if err != nil {
		return res, err
	}

	for _, crmStory := range crmStories {
		res = append(res, viewmodel.CrmStoryVm{
			ID:          crmStory.ID,
			Slug:        crmStory.Slug,
			Name:        crmStory.Name,
			ProfitCount: "",
			CreatedAt:   crmStory.CreatedAt,
			UpdatedAt:   crmStory.UpdatedAt,
		})
	}

	return res, err
}

func (uc CrmStoryUseCase) ReadBy(column, value string) (res viewmodel.CrmStoryVm, err error) {
	repository := actions.NewCrmStoryRepository(uc.DB)
	crmStory, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.CrmStoryVm{
		ID:          crmStory.ID,
		Slug:        crmStory.Slug,
		Name:        crmStory.Name,
		ProfitCount: "",
		CreatedAt:   crmStory.CreatedAt,
		UpdatedAt:   crmStory.UpdatedAt,
	}

	return res, err
}

func (uc CrmStoryUseCase) Edit(ID string, input *requests.CrmStoryRequest) (err error) {
	repository := actions.NewCrmStoryRepository(uc.DB)
	pusherUc := PusherUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy(ID, "slug", slug.Make(input.Name))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.CrmStoryVm{
		ID:        ID,
		Slug:      slug.Make(input.Name),
		Name:      input.Name,
		UpdatedAt: now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		return err
	}
	_ = pusherUc.Broadcast(pusher.EventStory, body)

	return nil
}

func (uc CrmStoryUseCase) Add(input *requests.CrmStoryRequest) (err error) {
	repository := actions.NewCrmStoryRepository(uc.DB)
	pusherUc := PusherUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "slug", slug.Make(input.Name))
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.CrmStoryVm{
		Slug:      slug.Make(input.Name),
		Name:      input.Name,
		UpdatedAt: now,
	}
	body.ID, err = repository.Add(body)
	if err != nil {
		return err
	}
	_ = pusherUc.Broadcast(pusher.EventStory, body)

	return nil
}

func (uc CrmStoryUseCase) Delete(ID string) (err error) {
	repository := actions.NewCrmStoryRepository(uc.DB)
	pusherUc := PusherUseCase{UcContract: uc.UcContract}
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
		_ = pusherUc.Broadcast("new-stage", ID)
	}

	return nil
}

func (uc CrmStoryUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewCrmStoryRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}
