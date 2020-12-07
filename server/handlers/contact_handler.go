package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/pkg/str"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"
)

type ContactHandler struct {
	Handler
}

func (handler ContactHandler) Browse(ctx echo.Context) error {
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	filters := make(map[string]interface{})
	if ctx.QueryParam("branch_name") != "" {
		filters["branch_name"] = ctx.QueryParam("branch_name")
	}

	if ctx.QueryParam("travel_agent_name") != "" {
		filters["travel_agent_name"] = ctx.QueryParam("travel_agent_name")
	}

	if ctx.QueryParam("address") != "" {
		filters["address"] = ctx.QueryParam("address")
	}

	if ctx.QueryParam("phone_number") != "" {
		filters["phone_number"] = ctx.QueryParam("phone_number")
	}

	if ctx.QueryParam("email") != "" {
		filters["email"] = ctx.QueryParam("email")
	}

	uc := usecase.ContactUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(filters, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler ContactHandler) BrowseAll(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	isZakatPartner := str.StringToBool(ctx.QueryParam("isZakatPartner"))

	uc := usecase.ContactUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll(search,isZakatPartner)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler ContactHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.ContactUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler ContactHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(requests.ContactRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.ContactUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(ID, input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler ContactHandler) Add(ctx echo.Context) error {
	input := new(requests.ContactRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.ContactUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler ContactHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.ContactUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
