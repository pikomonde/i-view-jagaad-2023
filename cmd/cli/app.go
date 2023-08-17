package main

import (
	"i-view-jagaad-2023/cmd/initialize"
	"i-view-jagaad-2023/repository"
	repoFile "i-view-jagaad-2023/repository/file"
	repoUserProvider "i-view-jagaad-2023/repository/user_provider"
	"i-view-jagaad-2023/service/user"
	"time"
)

func main() {

	// Initialize
	conf := initialize.NewConfig()
	httpCli := initialize.NewHttpCli(10 * time.Second)

	// Repositories
	rUserProviderA := repoUserProvider.NewRepository(httpCli, conf.ProviderAURL)
	rUserProviderB := repoUserProvider.NewRepository(httpCli, conf.ProviderBURL)
	rFile := repoFile.NewRepository("users.csv")

	// Services
	sUser := user.NewService(
		[]repository.UserProviderRepository{rUserProviderA, rUserProviderB},
		rFile,
	)
	sUser.FetchUsersFromProviders()

	// Delivery

	// Start services

}
