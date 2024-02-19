package cmd

import (
	"fmt"
	"netflix/config"
	"netflix/models"
	"netflix/services"

	"github.com/spf13/cobra"
)

var TitlesWithCreditsCmd = &cobra.Command{
	Use:   "title-with-credits",
	Short: "Retrieve detailed information about titles along with their associated credits.",
	Long: `The getTitlesWithCredits command allows you to fetch comprehensive details about titles, including their descriptions, release information, and genres, along with credits such as actors and their roles.
	This command provides a comprehensive overview of each title's cast and crew, facilitating a deeper understanding of the content.`,
	Run: func(cmd *cobra.Command, args []string) {

		titles, err := models.ReadTitles(config.TitleFilePath)
		if err != nil {
			fmt.Println("Error reading titles:", err)
			return
		}

		credits, err := models.ReadCredits(config.CreditFilePath)
		if err != nil {
			fmt.Println("Error reading credits:", err)
			return
		}

		titleWithCredits := services.ListTitlesWithCredits(titles, credits)

		for _, titleWithCredit := range titleWithCredits {
			fmt.Printf("Title: %s\n", titleWithCredit.Title)
			fmt.Printf("Type: %s\n", titleWithCredit.Type)
			fmt.Printf("Description: %s\n", titleWithCredit.Description)
			// Add more fields as needed
			fmt.Println("Credits:")
			for _, credit := range titleWithCredit.Credits {
				fmt.Printf("- %s as %s\n", credit.Name, credit.Character)
			}
			fmt.Println("------------------------------------")
		}

	},
}

func init() {
	titleCmd.AddCommand(TitlesWithCreditsCmd)
}
