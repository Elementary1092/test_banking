package handler

import (
	"github.com/Elementary1092/test_banking/internal"
	api "github.com/Elementary1092/test_banking/internal/adapters/http"
	"github.com/Elementary1092/test_banking/internal/adapters/http/httperr"
	httpMiddleware "github.com/Elementary1092/test_banking/internal/adapters/http/middleware"
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
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/json-iterator/go"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type HttpHandler struct {
	Router http.Handler
	app    *app.Application
}

func NewHandler(application *app.Application, config internal.Config) *HttpHandler {
	if application == nil {
		panic("application is nil in http/handler.NewHandler")
	}

	httpHandler := &HttpHandler{
		app: application,
	}

	r := chi.NewRouter()

	errHandlerFunc := func(w http.ResponseWriter, r *http.Request, err error) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	wrapper := api.ServerInterfaceWrapper{
		Handler:          httpHandler,
		ErrorHandlerFunc: errHandlerFunc,
	}

	r.Use(middleware.NoCache)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	addCorsMiddleware(r)

	jwtMiddleware := httpMiddleware.JWTMiddleware{
		Secret: config.TokenGen.Secret,
	}

	r.Group(func(r chi.Router) {
		r.Use(jwtMiddleware.Middleware)

		r.Get("/customer", wrapper.CustomerInfo)
		r.Get("/customer/accounts", wrapper.CustomerAccounts)
		r.Post("/customer/accounts", wrapper.AccountCreate)
		r.Get("/customer/accounts/{account_number}", wrapper.AccountGet)
		r.Post("/customer/accounts/{account_number}/replenish", wrapper.AccountReplenish)
		r.Post("/customer/accounts/{account_number}/transfer", wrapper.AccountTransfer)
		r.Post("/customer/accounts/{account_number}/withdraw", wrapper.AccountWithdraw)
	})
	r.Group(func(r chi.Router) {
		r.Post("/customer/refresh-token", wrapper.RefreshToken)
		r.Post("/customer/signin", wrapper.CustomerSignIn)
		r.Post("/customer/signup", wrapper.CustomerSignUp)
	})

	r.Mount("/api", r)

	httpHandler.Router = r

	return httpHandler
}

func addCorsMiddleware(router *chi.Mux) {
	corsMiddleware := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)
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
	userID, err := httpMiddleware.RetrieveUserID(r.Context())
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
	userID, err := httpMiddleware.RetrieveUserID(ctx)
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

	cmd := accountCreate.Command{
		UserID:   userID,
		Currency: request.Currency,
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
	userID, err := httpMiddleware.RetrieveUserID(ctx)
	if err != nil {
		httperr.WrapError(w, err)
		return
	}

	query := accountFind.Query{
		AccountNumber: accountNumber,
		UserID:        userID,
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
	userID, err := httpMiddleware.RetrieveUserID(ctx)
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
	userID, err := httpMiddleware.RetrieveUserID(ctx)
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

	cmd := accountReplenish.Command{
		AccountNumber: accountNumber,
		FromCard:      request.FromCard,
		Amount:        request.Amount,
		Currency:      request.Currency,
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
	userID, err := httpMiddleware.RetrieveUserID(ctx)
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
	userID, err := httpMiddleware.RetrieveUserID(ctx)
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
