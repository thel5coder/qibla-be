package usecase

import (
	"fmt"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
)

type YoutubeUseCase struct {
	*UcContract
}

//get video from youtube by channel id
func (uc YoutubeUseCase) GetVideoIDByChannelID(channelIDs []string,videoContentID string) (res []requests.VideoKajianRequest, err error) {
	parts := []string{"snippet"}

	for _, channelID := range channelIDs {
		call := uc.YoutubeService.Search.List(parts).
			Q("").
			ChannelId(channelID).
			MaxResults(defaultMaxResultYoutubeData).
			Type(defaultYoutubeSearchType).
			Order(defaultYoutubeOrder)

		response, err := call.Do()
		if err != nil {
			fmt.Println(err.Error())
			return res, err
		}

		for _, item := range response.Items {
			video,err := uc.GetVideo(item.Id.VideoId)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-youtube-getVideo")
				err = nil
			}

			res = append(res,requests.VideoKajianRequest{
				YoutubeVideoID: item.Id.VideoId,
				VideoContentID: videoContentID,
				Title:          video.Title,
				ChannelName:    video.ChannelName,
				Description:    video.Description,
				EmbeddedPlayer: video.EmbeddedPlayer,
				Thumbnails:     item.Snippet.Thumbnails,
				PublishedAt:    video.PublishedAt,
			})
		}
	}

	return res, err
}

func (uc YoutubeUseCase) GetVideo(videoID string) (res viewmodel.VideoDetailVm, err error) {
	parts := []string{"snippet", "player"}

	call := uc.YoutubeService.Videos.List(parts).
		MaxResults(defaultMaxResultYoutubeData).
		Id(videoID)
	response,err := call.Do()
	if err != nil {
		return res,err
	}
	
	res = viewmodel.VideoDetailVm{
		Title:          response.Items[0].Snippet.Title,
		ChannelName:    response.Items[0].Snippet.ChannelTitle,
		Description:    response.Items[0].Snippet.Description,
		EmbeddedPlayer: response.Items[0].Player.EmbedHtml,
		PublishedAt:    response.Items[0].Snippet.PublishedAt,
	}

	return res,err
}
