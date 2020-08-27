package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"qibla-backend/usecase"
	"strings"
)

type FileHandler struct {
	Handler
}

func (handler FileHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.FileUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler FileHandler) Add(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	splits := strings.Split(file.Filename, ".")
	ext := splits[len(splits)-1]
	fmt.Println(ext)

	uc := usecase.FileUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.UploadFile(file)

	return handler.SendResponse(ctx, res, nil, err)
}
