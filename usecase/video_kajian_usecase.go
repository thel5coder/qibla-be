package usecase

import "qibla-backend/usecase/viewmodel"

type VideoKajianUseCase struct {
	*UcContract
}

func (uc VideoKajianUseCase) Browse() (res []viewmodel.ChannelVideoVm,err error){
	videoContentUc := VideoContentUseCase{UcContract:uc.UcContract}
	youtubeUc := YoutubeUseCase{UcContract:uc.UcContract}
	var channelIDs []string

	videoContents,err := videoContentUc.BrowseAll()
	if err != nil {
		return res,err
	}

	for _,videoContent := range videoContents{
		channelIDs = append(channelIDs,videoContent.ChannelID)
	}

	res,err = youtubeUc.GetVideoIDByChannelID(channelIDs)
	if err != nil {
		return res,err
	}

	return res,err
}

func (uc VideoKajianUseCase) Read(videoID string) (res viewmodel.VideoDetailVm,err error){
	youtubeUc := YoutubeUseCase{UcContract:uc.UcContract}

	res,err = youtubeUc.GetVideo(videoID)
	if err != nil {
		return res,err
	}

	return res,err
}
