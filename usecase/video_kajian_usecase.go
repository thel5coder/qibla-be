package usecase

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/interfacepkg"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type VideoKajianUseCase struct {
	*UcContract
}

//browse
func (uc VideoKajianUseCase) Browse(videoContentID, search, order, sort string, page, limit int) (res []viewmodel.VideoDetailVm, pagination viewmodel.PaginationVm, err error) {
	//videoContentUc := VideoContentUseCase{UcContract: uc.UcContract}
	//youtubeUc := YoutubeUseCase{UcContract: uc.UcContract}
	//var channelIDs []string
	//
	//videoContents, err := videoContentUc.BrowseAll()
	//if err != nil {
	//	return res, err
	//}
	//
	//for _, videoContent := range videoContents {
	//	channelIDs = append(channelIDs, videoContent.ChannelID)
	//}
	//
	//res, err = youtubeUc.GetVideoIDByChannelID(channelIDs)
	//if err != nil {
	//	return res, err
	//}
	repository := actions.NewVideContentDetailRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	videoKajians, count, err := repository.BrowseByVideoContentID(videoContentID, search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, videoKajian := range videoKajians {
		res = append(res, uc.buildBody(videoKajian))
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

//read
func (uc VideoKajianUseCase) Read(ID string) (res viewmodel.VideoDetailVm, err error) {
	//youtubeUc := YoutubeUseCase{UcContract:uc.UcContract}
	//
	//res,err = youtubeUc.GetVideo(videoID)
	//if err != nil {
	//	return res,err
	//}
	//
	//return res,err

	repository := actions.NewVideContentDetailRepository(uc.DB)
	videoKajian, err := repository.Read(ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-videoContentDetail-read")
		return res, err
	}
	res = uc.buildBody(videoKajian)

	return res, nil
}

//add
func (uc VideoKajianUseCase) Add(input *requests.VideoKajianRequest) (err error) {
	repository := actions.NewVideContentDetailRepository(uc.DB)
	now := time.Now().UTC()



	model := models.VideoContentDetails{
		ID:             "",
		VideoContentID: input.VideoContentID,
		Title:          input.Title,
		ChannelTitle:   input.ChannelName,
		EmbeddedUrl:    input.EmbeddedPlayer,
		Thumbnails:     interfacepkg.Marshall(input.Thumbnails),
		Description:    input.Description,
		PublishedAt:    input.PublishedAt,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	_, err = repository.Add(model)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-videoContentDetail-add")
		return err
	}

	return nil
}

//delete by
func (uc VideoKajianUseCase) DeleteBy(column, value, operator string) (err error) {
	repository := actions.NewVideContentDetailRepository(uc.DB)
	now := time.Now().UTC()

	model := models.VideoContentDetails{
		UpdatedAt: now,
		DeletedAt: sql.NullTime{Time: now},
	}
	_, err = repository.DeleteBy(column, value, operator, model)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-videoContentDetail-deleteBy")
		return err
	}

	return nil
}

//store
func (uc VideoKajianUseCase) Store(inputs []requests.VideoKajianRequest) (err error) {
	for _, input := range inputs {
		count, err := uc.countBy("youtube_video_id", input.YoutubeVideoID, "=")
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-videoContentDetail-countByYoutubeVideoID")
			return err
		}

		if count == 0 {
			err = uc.Add(&input)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-videoContentDetail-add")
				return err
			}
		}
	}

	return nil
}

//count by
func (uc VideoKajianUseCase) countBy(column, value, operator string) (res int, err error) {
	repository := actions.NewVideContentDetailRepository(uc.DB)
	res, err = repository.CountBy(column, value, operator)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-videoContentDetail-countBy")
		return res, err
	}

	return res, nil
}

//build body
func (uc VideoKajianUseCase) buildBody(model models.VideoContentDetails) viewmodel.VideoDetailVm {
	var thumbnails map[string]interface{}
	interfacepkg.UnmarshallCb(model.Thumbnails,&thumbnails)

	return viewmodel.VideoDetailVm{
		ID:             model.ID,
		Title:          model.Title,
		ChannelName:    model.ChannelTitle,
		Description:    model.Description,
		EmbeddedPlayer: model.EmbeddedUrl,
		Thumbnails:     thumbnails,
		PublishedAt:    model.PublishedAt,
		CreatedAt:      model.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      model.UpdatedAt.Format(time.RFC3339),
	}
}
