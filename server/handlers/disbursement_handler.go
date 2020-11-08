package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
	"strconv"
)

// DisbursementHandler ...
type DisbursementHandler struct {
	Handler
}

// Browse ..
func (handler DisbursementHandler) Browse(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	contactTravelAgentName := ctx.QueryParam("contact_travel_agent_name")
	contactBranchName := ctx.QueryParam("contact_branch_name")
	total := ctx.QueryParam("total")
	startPeriod := ctx.QueryParam("start_period")
	endPeriod := ctx.QueryParam("end_period")
	contactAccountBankName := ctx.QueryParam("contact_account_bank_name")
	status := ctx.QueryParam("status")
	disburseAt := ctx.QueryParam("disburse_at")
	originAccountBankName := ctx.QueryParam("origin_account_bank_name")
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.DisbursementUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(search, contactTravelAgentName, contactBranchName, total, startPeriod, endPeriod, contactAccountBankName, status, disburseAt, originAccountBankName, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

// BrowseAll ...
func (handler DisbursementHandler) BrowseAll(ctx echo.Context) error {
	uc := usecase.DisbursementUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll()

	return handler.SendResponse(ctx, res, nil, err)
}

// Read ...
func (handler DisbursementHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.DisbursementUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}
