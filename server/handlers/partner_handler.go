package handlers

import (
	"net/http"
	"qibla-backend/pkg/jwt"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type PartnerHandler struct {
	Handler
}

func (handler PartnerHandler) Browse(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.PartnerUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(search, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler PartnerHandler) BrowseProfilePartner(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.PartnerUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.BrowseProfilePartner(search, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler PartnerHandler) ReadByPk(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.PartnerUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("p.id", ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler PartnerHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(requests.PartnerRegisterRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.PartnerUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(ID, input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler PartnerHandler) EditVerify(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(requests.PartnerVerifyRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.PartnerUseCase{UcContract: handler.UseCaseContract}
	err := uc.EditVerify(ID, input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler PartnerHandler) EditAccountStatus(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(requests.PartnerStatusRequest)
	claims := ctx.Get("user").(*jwt.CustomClaims)
	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.PartnerUseCase{UcContract: handler.UseCaseContract}
	err := uc.EditBoolStatus(ID, "is_active", input.Reason, claims.Id, input.Password, input.IsActive)
	if err != nil {
		if err.Error() == messages.CredentialDoNotMatch {
			return handler.SendResponseUnauthorized(ctx, err)
		}
	}

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler PartnerHandler) EditWebinarStatus(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(requests.PartnerStatusRequest)
	claims := ctx.Get("user").(*jwt.CustomClaims)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.PartnerUseCase{UcContract: handler.UseCaseContract}
	err := uc.EditBoolStatus(ID, "webinar_status", input.Reason, claims.Id, input.Password, input.IsActive)
	if err != nil {
		if err.Error() == messages.CredentialDoNotMatch {
			return handler.SendResponseUnauthorized(ctx, err)
		}
	}

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler PartnerHandler) Add(ctx echo.Context) error {
	input := new(requests.PartnerRegisterRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.PartnerUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Add(input)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler PartnerHandler) DeleteByPk(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.PartnerUseCase{UcContract: handler.UseCaseContract}
	err := uc.DeleteByPk(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
