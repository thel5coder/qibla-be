package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"
)

// DisbursementHandler ...
type DisbursementHandler struct {
	Handler
}

// Browse ..
func (handler DisbursementHandler) Browse(ctx echo.Context) error {
	filters := make(map[string]interface{})

	if ctx.QueryParam("contact_travel_agent_name") != "" {
		filters["contact_travel_agent_name"] = ctx.QueryParam("contact_travel_agent_name")
	}
	if ctx.QueryParam("contact_branch_name") != "" {
		filters["contact_branch_name"] = ctx.QueryParam("contact_branch_name")
	}
	if ctx.QueryParam("total") != "" {
		filters["total"] = ctx.QueryParam("total")
	}
	if ctx.QueryParam("start_period") != "" {
		filters["start_period"] = ctx.QueryParam("start_period")
	}
	if ctx.QueryParam("end_period") != "" {
		filters["end_period"] = ctx.QueryParam("end_period")
	}
	if ctx.QueryParam("contact_account_bank_name") != "" {
		filters["contact_account_bank_name"] = ctx.QueryParam("contact_account_bank_name")
	}
	if ctx.QueryParam("status") != "" {
		filters["status"] = ctx.QueryParam("status")
	}
	if ctx.QueryParam("disburse_at") != "" {
		filters["disburse_at"] = ctx.QueryParam("disburse_at")
	}
	if ctx.QueryParam("origin_account_bank_name") != "" {
		filters["origin_account_bank_name"] = ctx.QueryParam("origin_account_bank_name")
	}

	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.DisbursementUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(filters, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

// BrowseAll ...
func (handler DisbursementHandler) BrowseAll(ctx echo.Context) error {
	status := ctx.QueryParam("status")

	uc := usecase.DisbursementUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll(status)

	return handler.SendResponse(ctx, res, nil, err)
}

// Read ...
func (handler DisbursementHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.DisbursementUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

// PdfExport ...
func (handler DisbursementHandler) PdfExport(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.PdfUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Disbursement(ID)

	return handler.SendResponseFile(ctx, res, "application/pdf", err)
}

// Request ...
func (handler DisbursementHandler) Request(ctx echo.Context) error {
	input := new(requests.DisbursementReqRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	var err error
	handler.UseCaseContract.TX, err = handler.UseCaseContract.DB.Begin()
	if err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	uc := usecase.DisbursementUseCase{UcContract: handler.UseCaseContract}
	err = uc.DisbursementReq(input)
	if err != nil {
		handler.UseCaseContract.TX.Rollback()
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	handler.UseCaseContract.TX.Commit()

	return handler.SendResponse(ctx, nil, nil, err)
}
