package main

import (
	"context"
	"fmt"
	"os"
	"tp-project/internal/delivery"
	"tp-project/pkg/db/postgres"
	"tp-project/pkg/logger"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New("Sila")
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)

	postgresClient := postgres.New(ctx, os.Getenv("POSTGRES"))
	mainLogger.Info(ctx, fmt.Sprintf("Client DB: %v", postgresClient))
	delivery.Start(postgresClient.DB, &mainLogger)
}
