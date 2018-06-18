package cmd

import (
	"fmt"

	apple "github.com/nareshganesan/reviews/appstore"
	"github.com/spf13/cobra"
)

var searchApps bool
var appName string
var appDetails bool
var reviewDetails bool
var appID int
var countryID int
var pageNo int

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Reviews server exposes API to access reviews",
	Long: `Reviews server exposes API to access reviews.

Some of the Features are listed below.

- Search apps based on app name
- Get reviews for apps based on app id and country id
- Get app details based on app id and country id

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")

		if reviewDetails {
			app := apple.GetReviews(appID, countryID, pageNo)
			// app := apple.GetAllReviews(appID, countryID)
			fmt.Println(app)
		} else if searchApps {
			fmt.Println(appName)
			apps := apple.SearchApps(appName)
			for _, app := range apps {
				fmt.Println(app.TrackID, app.TrackName)
			}
		} else {
			app := apple.GetAppDetails(appID)
			fmt.Println(app)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")
	serverCmd.PersistentFlags().BoolVarP(&searchApps, "search", "s", false, "Search apps (Default: false)")
	serverCmd.PersistentFlags().BoolVarP(&appDetails, "app", "a", true, "Get app details (Default: true)")
	serverCmd.PersistentFlags().BoolVarP(&reviewDetails, "review", "r", false, "Get app review details (Default: false)")
	serverCmd.PersistentFlags().StringVarP(&appName, "appname", "n", "whatsapp", "App name to search (Default: Whatsapp)")
	serverCmd.PersistentFlags().IntVarP(&appID, "appid", "i", 368677368, "App id to get reviews (Default: UBER)")
	serverCmd.PersistentFlags().IntVarP(&countryID, "countryid", "c", 143441, "Country id of app (Default: US)")
	serverCmd.PersistentFlags().IntVarP(&pageNo, "page", "p", 0, "Page no of the reviews list")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
