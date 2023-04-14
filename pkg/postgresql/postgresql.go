package postgresql

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/url"
	"time"
)

var (
	ErrorNoDatabaseName       = errors.New("database name is empty")
	ErrorNoUsernameOrPassword = errors.New("no username or password")
	ErrorNoHost               = errors.New("no host is provided")
)

func NewClient(ctx context.Context, config Config) (pool *pgxpool.Pool, err error) {
	connURL, err := makeDatabaseURL(config)
	if err != nil {
		return nil, err
	}

	// This function is used to abstract cancel() call from for loop
	// Calling cancel() in for loop may cause resource leak
	tryToConnect := func(connector func()) {
		connector()
	}

	// Not always app connects to the db on the first try.
	// So, here function tries to connect to the db multiple times.
	for i := uint8(0); i < config.MaxAttempts; i++ {
		tryToConnect(func() {
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			pool, err = pgxpool.Connect(ctx, connURL)
		})
		if err == nil {
			break
		}

		time.Sleep(5 * time.Second)
	}

	return
}

func makeDatabaseURL(config Config) (string, error) {
	if config.Username == "" || config.Password == "" {
		return "", ErrorNoUsernameOrPassword
	}

	if config.Host == "" {
		return "", ErrorNoHost
	}

	if config.DBName == "" {
		return "", ErrorNoDatabaseName
	}

	host := config.Host
	if config.Port != "" {
		host += ":" + config.Port
	}

	u := url.URL{
		Scheme: "postgresql",
		User:   url.UserPassword(config.Username, config.Password),
		Host:   host,
		Path:   config.DBName,
	}

	return u.String(), nil
}
