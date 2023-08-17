package cli

import (
	"i-view-jagaad-2023/service"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Cli struct {
	UserService service.UserService
}

func NewDelivery(
	userService service.UserService,
) *Cli {
	return &Cli{
		UserService: userService,
	}
}

func (d *Cli) Start() {
	rootCmd := &cobra.Command{
		Use:   "i-view-jagaad-2023",
		Short: "Cli app to fetch, save, and search user by tag",
	}

	fetchCmd := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch users from providers and save it to csv file",
		Run: func(cmd *cobra.Command, args []string) {
			d.UserService.FetchUsersFromProviders()
		},
	}
	rootCmd.AddCommand(fetchCmd)

	getUsersCmd := &cobra.Command{
		Use:   "get-users",
		Short: "Search users by tags",
		Run: func(cmd *cobra.Command, args []string) {
			tagsStrArr, err := cmd.Flags().GetStringArray("tag")
			if err != nil {
				log.Errorf("Error cannot get tags flag, err: %s", err.Error())
			}

			tags := make([]string, 0)
			for _, tagsStr := range tagsStrArr {
				tags = append(tags, strings.Split(tagsStr, ",")...)
			}

			d.UserService.GetUserByTags(tags)
		},
	}
	rootCmd.AddCommand(getUsersCmd)
	getUsersCmd.Flags().StringArrayP("tag", "t", []string{}, "Search users by tag")

	rootCmd.Execute()
}
