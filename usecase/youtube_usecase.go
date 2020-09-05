package usecase

import (
	"encoding/json"
	"fmt"
	"qibla-backend/usecase/viewmodel"
)

type YoutubeUseCase struct {
	*UcContract
}

func (uc YoutubeUseCase) GetVideoIDByChannelID(channelIDs []string) (res []viewmodel.ChannelVideoVm, err error) {
	parts := []string{"snippet"}

	for _, channelID := range channelIDs {
		call := uc.YoutubeClient.Search.List(parts).
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
			var thumbnails map[string]interface{}
			jsonThumbnail, _ := item.Snippet.Thumbnails.MarshalJSON()
			json.Unmarshal(jsonThumbnail, &thumbnails)

			res = append(res, viewmodel.ChannelVideoVm{
				ID:          item.Id.VideoId,
				Title:       item.Snippet.Title,
				ChannelName: item.Snippet.ChannelTitle,
				Thumbnails:  thumbnails,
				Description: item.Snippet.Description,
				PublishedAt: item.Snippet.PublishedAt,
			})
		}
	}

	return res, err
}

func (uc YoutubeUseCase) GetVideo(videoID string) (res viewmodel.VideoDetailVm, err error) {
	parts := []string{"snippet", "player"}

	call := uc.YoutubeClient.Videos.List(parts).
		MaxResults(defaultMaxResultYoutubeData).
		Id(videoID)
	response,err := call.Do()
	if err != nil {
		return res,err
	}
	
	res = viewmodel.VideoDetailVm{
		Title:          response.Items[1].Snippet.Title,
		ChannelName:    response.Items[1].Snippet.ChannelTitle,
		Description:    response.Items[1].Snippet.Description,
		EmbeddedPlayer: response.Items[1].Player.EmbedHtml,
		PublishedAt:    response.Items[1].Snippet.PublishedAt,
	}

	return res,err
}
