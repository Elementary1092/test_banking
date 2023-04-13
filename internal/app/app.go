package app

import (
	accountCreate "github.com/Elementary1092/test_banking/internal/domain/account/command/create"
	accountReplenish "github.com/Elementary1092/test_banking/internal/domain/account/command/replenish"
	accountTransfer "github.com/Elementary1092/test_banking/internal/domain/account/command/transfer"
	accountWithdraw "github.com/Elementary1092/test_banking/internal/domain/account/command/withdraw"
	accountFind "github.com/Elementary1092/test_banking/internal/domain/account/query/find"
	accountList "github.com/Elementary1092/test_banking/internal/domain/account/query/list"
	accountTransactions "github.com/Elementary1092/test_banking/internal/domain/account/query/transactions"
	customerCreate "github.com/Elementary1092/test_banking/internal/domain/customer/command/create"
	customerAuth "github.com/Elementary1092/test_banking/internal/domain/customer/query/auth"
	customerFind "github.com/Elementary1092/test_banking/internal/domain/customer/query/find"
	"github.com/Elementary1092/test_banking/internal/domain/token"
)

type Application struct {
	Account  AccountActions
	Customer CustomerActions
	TokenMgr *token.Handler // tokens manager
}

type AccountActions struct {
	Commands AccountCommands
	Queries  AccountQueries
}

type CustomerActions struct {
	Commands CustomerCommands
	Queries  CustomerQueries
}

type AccountCommands struct {
	Create    *accountCreate.Handler
	Replenish *accountReplenish.Handler
	Transfer  *accountTransfer.Handler
	Withdraw  *accountWithdraw.Handler
}

type AccountQueries struct {
	Find             *accountFind.Handler
	FindTransactions *accountTransactions.Handler
	List             *accountList.Handler
}

type CustomerCommands struct {
	Create *customerCreate.Handler
}

type CustomerQueries struct {
	Auth *customerAuth.Handler
	Find *customerFind.Handler
}
