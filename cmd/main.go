package main

import (
	"context"

	lg "github.com/veerLakshya/my-go-all/logger"
	"github.com/veerlakshya/user-service/config"
	"github.com/veerlakshya/user-service/internal/infrastructure/repository"
)

func main() {
	logger := lg.GetConsoleLogger()
	appConfig := config.GetAppConfig()

	ctx := context.Background()

	db := config.GetDBConn(ctx, logger, appConfig.DB)

	userDetailsRepo := repository.NewUserDetailsRepository(db)
	_ = userDetailsRepo
}
