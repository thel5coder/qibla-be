package usecase

import (
	"qibla-backend/db/repositories/actions"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type VideoContentUseCase struct {
	*UcContract
}

func (uc VideoContentUseCase) Browse(order,sort string,page,limit int) (res []viewmodel.VideoContentVm,pagination viewmodel.PaginationVm,err error){
	repository := actions.NewVideoContentRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	videoContents,count,err := repository.Browse(order,sort,limit,offset)
	if err != nil {
		return res,pagination,err
	}

	for _,videoContent := range videoContents{
		res = append(res,viewmodel.VideoContentVm{
			ID:        videoContent.ID,
			Channel:   videoContent.Channel,
			Link:      videoContent.Links,
			CreatedAt: videoContent.CreatedAt,
			UpdatedAt: videoContent.UpdatedAt,
			DeletedAt: videoContent.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc VideoContentUseCase) Add(input *requests.VideoContentRequest) (err error){
	repository := actions.NewVideoContentRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	body := viewmodel.VideoContentVm{
		Channel:   input.Channel,
		Link:      input.Link,
		CreatedAt: now,
		UpdatedAt: now,
	}
	_,err = repository.Add(body)
	if err != nil {
		return err
	}

	return nil
}
