package main

import (
	"i-view-jagaad-2023/cmd/initialize"
	deliveryCli "i-view-jagaad-2023/delivery/cli"
	"i-view-jagaad-2023/repository"
	repoFile "i-view-jagaad-2023/repository/file"
	repoUserProvider "i-view-jagaad-2023/repository/user_provider"
	serviceUser "i-view-jagaad-2023/service/user"
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
	sUser := serviceUser.NewService(
		[]repository.UserProviderRepository{rUserProviderA, rUserProviderB},
		rFile,
	)

	// Delivery
	dCli := deliveryCli.NewDelivery(sUser)

	// Start services
	dCli.Start()

}
