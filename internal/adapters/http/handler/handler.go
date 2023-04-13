package handler

import (
	api "github.com/Elementary1092/test_banking/internal/adapters/http"
	"github.com/Elementary1092/test_banking/internal/adapters/http/httperr"
	"github.com/Elementary1092/test_banking/internal/adapters/http/middleware"
	"github.com/Elementary1092/test_banking/internal/app"
	accountCreate "github.com/Elementary1092/test_banking/internal/domain/account/command/create"
	accountReplenish "github.com/Elementary1092/test_banking/internal/domain/account/command/replenish"
	accountTransfer "github.com/Elementary1092/test_banking/internal/domain/account/command/transfer"
	accountWithdraw "github.com/Elementary1092/test_banking/internal/domain/account/command/withdraw"
	accountFind "github.com/Elementary1092/test_banking/internal/domain/account/query/find"
	accountList "github.com/Elementary1092/test_banking/internal/domain/account/query/list"
	customerCreate "github.com/Elementary1092/test_banking/internal/domain/customer/command/create"
	customerAuth "github.com/Elementary1092/test_banking/internal/domain/customer/query/auth"
	customerFind "github.com/Elementary1092/test_banking/internal/domain/customer/query/find"
	"github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type HttpHandler struct {
	app *app.Application
}

func NewHandler(app *app.Application) *HttpHandler {
	if app == nil {
		panic("application is nil in http/handler.NewHandler")
	}

	return &HttpHandler{
		app: app,
	}
}

func (h *HttpHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request api.RefreshTokenRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		httperr.WrapError(w, err)
		return
	}

	idToken, refreshToken, err := h.app.TokenMgr.RefreshToken(ctx, request.RefreshToken)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	Render(w, MapSignInResponse(idToken, refreshToken), http.StatusOK)
}

func (h *HttpHandler) CustomerSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	var request api.SignInRequest
	if err := decoder.Decode(&request); err != nil {
		httperr.BadRequest(w, "request-with-invalid-data")
		return
	}

	query := customerAuth.Query{
		Email:    request.Email,
		Password: request.Password,
	}
	customer, err := h.app.Customer.Queries.Auth.Handle(ctx, query)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	idToken, refreshToken, err := h.app.TokenMgr.GenerateToken(ctx, customer.UUID())
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	Render(w, MapSignInResponse(idToken, refreshToken), http.StatusOK)
}

func (h *HttpHandler) CustomerSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	var request api.SignUpRequest
	if err := decoder.Decode(&request); err != nil {
		httperr.BadRequest(w, "request-with-invalid-data")
		return
	}

	cmd := customerCreate.Customer{
		Email:    request.Email,
		Password: request.Password,
	}
	if err := h.app.Customer.Commands.Create.Handle(ctx, cmd); err != nil {
		httperr.WrapError(w, err)
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) CustomerInfo(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.RetrieveUserID(r.Context())
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	ctx := r.Context()
	query := customerFind.Query{
		UUID: userID,
	}
	customer, err := h.app.Customer.Queries.Find.Handle(ctx, query)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	Render(w, MapCustomerRead(customer), http.StatusOK)
}

func (h *HttpHandler) AccountCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := middleware.RetrieveUserID(ctx)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	var request api.CreateAccountRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		httperr.BadRequest(w, "request-with-invalid-data")
		return
	}

	currency, err := RetrieveCurrency(request.Currency)
	if err != nil {
		httperr.BadRequest(w, "invalid-currency")
		return
	}
	cmd := accountCreate.Command{
		UserID:   userID,
		Currency: currency,
	}
	if err := h.app.Account.Commands.Create.Handle(ctx, cmd); err != nil {
		httperr.WrapError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) AccountGet(
	w http.ResponseWriter,
	r *http.Request,
	accountNumber string,
) {
	ctx := r.Context()
	userID, err := middleware.RetrieveUserID(ctx)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	query := accountFind.Query{
		AccountNumber: userID,
		UserID:        accountNumber,
	}
	account, err := h.app.Account.Queries.Find.Handle(ctx, query)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	Render(w, MapAccountRead(account), http.StatusOK)
}

func (h *HttpHandler) CustomerAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := middleware.RetrieveUserID(ctx)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	// Limit and offset will not be used
	// because assumed that customer will not have large number of accounts
	// but, in the future, they may be added easily
	query := accountList.Query{
		UserID: userID,
	}

	accounts, err := h.app.Account.Queries.List.Handle(ctx, query)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	resp := make([]api.GetAccountResponse, len(accounts))
	for i, account := range accounts {
		resp[i] = MapAccountRead(account)
	}

	Render(w, resp, http.StatusOK)
}

func (h *HttpHandler) AccountReplenish(
	w http.ResponseWriter,
	r *http.Request,
	accountNumber string,
) {
	ctx := r.Context()
	userID, err := middleware.RetrieveUserID(ctx)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	var request api.ReplenishRequest
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&request); err != nil {
		httperr.BadRequest(w, "request-with-invalid-data")
		return
	}

	currency, err := RetrieveCurrency(request.Currency)
	if err != nil {
		httperr.BadRequest(w, "invalid-currency")
		return
	}
	cmd := accountReplenish.Command{
		AccountNumber: accountNumber,
		FromCard:      request.FromCard,
		Amount:        request.Amount,
		Currency:      currency,
		UserID:        userID,
	}
	if err = h.app.Account.Commands.Replenish.Handle(ctx, cmd); err != nil {
		httperr.WrapError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HttpHandler) AccountTransfer(
	w http.ResponseWriter,
	r *http.Request,
	accountNumber string,
) {
	ctx := r.Context()
	userID, err := middleware.RetrieveUserID(ctx)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	var request api.TransferRequest
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&request); err != nil {
		httperr.BadRequest(w, "request-with-invalid-data")
		return
	}

	cmd := accountTransfer.Command{
		From:   accountNumber,
		To:     request.ToAccount,
		Amount: request.Amount,
		UserID: userID,
	}
	if err = h.app.Account.Commands.Transfer.Handle(ctx, cmd); err != nil {
		httperr.WrapError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HttpHandler) AccountWithdraw(
	w http.ResponseWriter,
	r *http.Request,
	accountNumber string,
) {
	ctx := r.Context()
	userID, err := middleware.RetrieveUserID(ctx)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	var request api.WithdrawRequest
	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&request); err != nil {
		httperr.BadRequest(w, "request-with-invalid-data")
		return
	}

	cmd := accountWithdraw.Command{
		From:   accountNumber,
		To:     request.ToCard,
		Amount: request.Amount,
		UserID: userID,
	}
	if err = h.app.Account.Commands.Withdraw.Handle(ctx, cmd); err != nil {
		httperr.WrapError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
