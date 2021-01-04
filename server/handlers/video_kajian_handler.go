package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
	"strconv"
)

type VideoKajianHandler struct {
	Handler
}

func (handler VideoKajianHandler) Browse(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	videoContentID := ctx.QueryParam("video_content_id")

	uc := usecase.VideoKajianUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(videoContentID, search, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler VideoKajianHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.VideoKajianUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read(ID)

	return handler.SendResponse(ctx, res, nil, err)
}
