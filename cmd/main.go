package main

import (
	"github.com/Elementary1092/test_banking/internal"
	"github.com/Elementary1092/test_banking/internal/adapters/http/handler"
	httpServer "github.com/Elementary1092/test_banking/internal/adapters/http/server"
	"github.com/Elementary1092/test_banking/internal/app"
	"github.com/joho/godotenv"
	"log"
)

// entry point
func main() {
	godotenv.Load(".env")

	config := internal.Parse()

	application := app.NewApplication(config)
	defer application.Close()

	httpHandler := handler.NewHandler(application, config)

	server := httpServer.New(httpServer.Config{
		Address:      config.HTTPServer.Address,
		Router:       httpHandler.Router,
		WriteTimeout: config.HTTPServer.WriteTimeout,
		ReadTimeout:  config.HTTPServer.ReadTimeout,
	})

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
