package usecase

import (
	"errors"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type VideoContentUseCase struct {
	*UcContract
}

func (uc VideoContentUseCase) Browse(order, sort string, page, limit int) (res []viewmodel.VideoContentVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewVideoContentRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	videoContents, count, err := repository.Browse(order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, videoContent := range videoContents {
		res = append(res, viewmodel.VideoContentVm{
			ID:        videoContent.ID,
			Channel:   videoContent.Channel,
			ChannelID: videoContent.ChannelID,
			Link:      videoContent.Links,
			IsActive:  videoContent.IsActive,
			CreatedAt: videoContent.CreatedAt,
			UpdatedAt: videoContent.UpdatedAt,
			DeletedAt: videoContent.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc VideoContentUseCase) BrowseAll() (res []viewmodel.VideoContentVm, err error) {
	repository := actions.NewVideoContentRepository(uc.DB)

	videoContents, err := repository.BrowseAll()
	if err != nil {
		return res, err
	}

	for _, videoContent := range videoContents {
		res = append(res, viewmodel.VideoContentVm{
			ID:        videoContent.ID,
			Channel:   videoContent.Channel,
			ChannelID: videoContent.ChannelID,
			Link:      videoContent.Links,
			IsActive:  videoContent.IsActive,
			CreatedAt: videoContent.CreatedAt,
			UpdatedAt: videoContent.UpdatedAt,
			DeletedAt: videoContent.DeletedAt.String,
		})
	}

	return res, err
}

func (uc VideoContentUseCase) ReadBy(column, value string) (res viewmodel.VideoContentVm, err error) {
	repository := actions.NewVideoContentRepository(uc.DB)
	videoContent, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.VideoContentVm{
		ID:        videoContent.ID,
		Channel:   videoContent.Channel,
		ChannelID: videoContent.ChannelID,
		Link:      videoContent.Links,
		IsActive:  videoContent.IsActive,
		CreatedAt: videoContent.CreatedAt,
		UpdatedAt: videoContent.UpdatedAt,
		DeletedAt: videoContent.DeletedAt.String,
	}

	return res, err
}

func (uc VideoContentUseCase) Edit(ID string, input *requests.VideoContentRequest) (err error) {
	repository := actions.NewVideoContentRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	link := strings.Split(input.Link, "/")
	channelID := link[len(link)-1]

	count, err := uc.CountBy(ID, "channel_id", channelID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-videoContent-countByChannelID")
		return err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-videoContentDetail-dataAlreadyExist")
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.VideoContentVm{
		ID:        ID,
		Channel:   input.Channel,
		ChannelID: channelID,
		Link:      input.Link,
		IsActive:  true,
		CreatedAt: now,
	}
	_, err = repository.Edit(body)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "query-videoContent-edit")
		return err
	}

	//fetch data from youtube and store to db
	err = uc.YoutubeStore(ID, channelID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-videoContent-youtubeStore")
		err = nil
	}

	return nil
}

func (uc VideoContentUseCase) Add(input *requests.VideoContentRequest) (err error) {
	repository := actions.NewVideoContentRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	link := strings.Split(input.Link, "/")
	channelID := link[len(link)-1]

	count, err := uc.CountBy("", "channel_id", channelID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-videoContent-countByChannelID")
		return err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-videoContentDetail-dataAlreadyExist")
		return errors.New(messages.DataAlreadyExist)
	}

	body := viewmodel.VideoContentVm{
		Channel:   input.Channel,
		ChannelID: channelID,
		Link:      input.Link,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	body.ID, err = repository.Add(body)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "query-videoContent-add")
		return err
	}

	//fetch data from youtube and store to db
	err = uc.YoutubeStore(body.ID, channelID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-videoContent-youtubeStore")
		err = nil
	}

	return nil
}

func (uc VideoContentUseCase) Delete(ID string) (err error) {
	repository := actions.NewVideoContentRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "id", ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-videoContent-countByID")
		return err
	}
	if count > 0 {
		_, err = repository.Delete(ID, now, now)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-videoContent-delete")
			return err
		}

		videoKajianUc := VideoKajianUseCase{UcContract:uc.UcContract}
		err = videoKajianUc.DeleteBy("video_content_id",ID,"=")
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-videoKajian-deleteByVideoContentID")
		}
	}

	return err
}

func (uc VideoContentUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewVideoContentRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}

//youtube usecase, fetch data from youtube and save to postgres
func (uc VideoContentUseCase) YoutubeStore(ID, channelID string) (err error) {
	//fetch video by channel id
	youtubeUc := YoutubeUseCase{UcContract: uc.UcContract}
	videoKajianRequest, err := youtubeUc.GetVideoIDByChannelID([]string{channelID}, ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-youtube-getVideoByChannelID")
	}

	//store video to db
	videoKajianUc := VideoKajianUseCase{UcContract: uc.UcContract}
	err = videoKajianUc.Store(videoKajianRequest)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-videoKajian-store")
	}

	return err
}
