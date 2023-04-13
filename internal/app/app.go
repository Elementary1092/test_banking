package app

import (
	"context"
	"fmt"
	"github.com/Elementary1092/test_banking/internal"
	accountDao "github.com/Elementary1092/test_banking/internal/adapters/dao/account"
	customerDao "github.com/Elementary1092/test_banking/internal/adapters/dao/customer"
	tokenDao "github.com/Elementary1092/test_banking/internal/adapters/dao/token"
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
	"github.com/Elementary1092/test_banking/pkg/postgresql"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Application struct {
	Account  AccountActions
	Customer CustomerActions
	TokenMgr *token.Handler // tokens manager

	// database client
	client *pgxpool.Pool
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

func NewApplication(config internal.Config) *Application {
	dbConfig := postgresql.Config{
		Username:    config.DB.Username,
		Password:    config.DB.Password,
		Host:        config.DB.Host,
		Port:        config.DB.Port,
		DBName:      config.DB.DBName,
		MaxAttempts: config.DB.MaxAttempts,
	}

	client, err := postgresql.NewClient(context.Background(), dbConfig)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize app: %v", err))
	}
	accountCommandDAO := accountDao.NewCommandDAO(client)
	accountQueryDAO := accountDao.NewQueryDAO(client)
	customerCommandDAO := customerDao.NewCommandDAO(client)
	customerQueryDAO := customerDao.NewQueryDAO(client)
	tokenMgrDAO := tokenDao.NewDAO(client)

	return &Application{
		Account: AccountActions{
			Commands: AccountCommands{
				Create:    accountCreate.NewCreateHandler(accountCommandDAO),
				Replenish: accountReplenish.NewHandler(accountCommandDAO),
				Transfer:  accountTransfer.NewHandler(accountCommandDAO),
				Withdraw:  accountWithdraw.NewHandler(accountCommandDAO),
			},
			Queries: AccountQueries{
				Find:             accountFind.NewHandler(accountQueryDAO),
				FindTransactions: accountTransactions.NewHandler(accountQueryDAO),
				List:             accountList.NewHandler(accountQueryDAO),
			},
		},
		Customer: CustomerActions{
			Commands: CustomerCommands{
				Create: customerCreate.NewHandler(customerCommandDAO),
			},
			Queries: CustomerQueries{
				Auth: customerAuth.NewHandler(customerQueryDAO),
				Find: customerFind.NewHandler(customerQueryDAO),
			},
		},
		TokenMgr: token.NewHandler(
			tokenMgrDAO,
			config.TokenGen.RefreshExpr,
			config.TokenGen.IdExpr,
			config.TokenGen.Issuer,
			config.TokenGen.Secret,
		),

		client: client,
	}
}

func (a *Application) Close() {
	a.client.Close()
}
