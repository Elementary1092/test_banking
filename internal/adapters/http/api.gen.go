// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /customer)
	CustomerInfo(w http.ResponseWriter, r *http.Request)

	// (GET /customer/accounts)
	CustomerAccounts(w http.ResponseWriter, r *http.Request)

	// (POST /customer/accounts)
	AccountCreate(w http.ResponseWriter, r *http.Request)

	// (GET /customer/accounts/{account_number})
	AccountGet(w http.ResponseWriter, r *http.Request, accountNumber string)

	// (POST /customer/accounts/{account_number}/replenish)
	AccountReplenish(w http.ResponseWriter, r *http.Request, accountNumber string)

	// (POST /customer/accounts/{account_number}/transfer)
	AccountTransfer(w http.ResponseWriter, r *http.Request, accountNumber string)

	// (POST /customer/accounts/{account_number}/withdraw)
	AccountWithdraw(w http.ResponseWriter, r *http.Request, accountNumber string)

	// (POST /customer/refresh-token)
	RefreshToken(w http.ResponseWriter, r *http.Request)

	// (POST /customer/signin)
	CustomerSignIn(w http.ResponseWriter, r *http.Request)

	// (POST /customer/signup)
	CustomerSignUp(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// CustomerInfo operation middleware
func (siw *ServerInterfaceWrapper) CustomerInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CustomerInfo(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CustomerAccounts operation middleware
func (siw *ServerInterfaceWrapper) CustomerAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CustomerAccounts(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AccountCreate operation middleware
func (siw *ServerInterfaceWrapper) AccountCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AccountCreate(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AccountGet operation middleware
func (siw *ServerInterfaceWrapper) AccountGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "account_number" -------------
	var accountNumber string

	err = runtime.BindStyledParameterWithLocation("simple", false, "account_number", runtime.ParamLocationPath, chi.URLParam(r, "account_number"), &accountNumber)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "account_number", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AccountGet(w, r, accountNumber)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AccountReplenish operation middleware
func (siw *ServerInterfaceWrapper) AccountReplenish(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "account_number" -------------
	var accountNumber string

	err = runtime.BindStyledParameterWithLocation("simple", false, "account_number", runtime.ParamLocationPath, chi.URLParam(r, "account_number"), &accountNumber)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "account_number", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AccountReplenish(w, r, accountNumber)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AccountTransfer operation middleware
func (siw *ServerInterfaceWrapper) AccountTransfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "account_number" -------------
	var accountNumber string

	err = runtime.BindStyledParameterWithLocation("simple", false, "account_number", runtime.ParamLocationPath, chi.URLParam(r, "account_number"), &accountNumber)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "account_number", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AccountTransfer(w, r, accountNumber)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AccountWithdraw operation middleware
func (siw *ServerInterfaceWrapper) AccountWithdraw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "account_number" -------------
	var accountNumber string

	err = runtime.BindStyledParameterWithLocation("simple", false, "account_number", runtime.ParamLocationPath, chi.URLParam(r, "account_number"), &accountNumber)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "account_number", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AccountWithdraw(w, r, accountNumber)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// RefreshToken operation middleware
func (siw *ServerInterfaceWrapper) RefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RefreshToken(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CustomerSignIn operation middleware
func (siw *ServerInterfaceWrapper) CustomerSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CustomerSignIn(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CustomerSignUp operation middleware
func (siw *ServerInterfaceWrapper) CustomerSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CustomerSignUp(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/customer", wrapper.CustomerInfo)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/customer/accounts", wrapper.CustomerAccounts)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/customer/accounts", wrapper.AccountCreate)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/customer/accounts/{account_number}", wrapper.AccountGet)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/customer/accounts/{account_number}/replenish", wrapper.AccountReplenish)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/customer/accounts/{account_number}/transfer", wrapper.AccountTransfer)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/customer/accounts/{account_number}/withdraw", wrapper.AccountWithdraw)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/customer/refresh-token", wrapper.RefreshToken)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/customer/signin", wrapper.CustomerSignIn)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/customer/signup", wrapper.CustomerSignUp)
	})

	return r
}
