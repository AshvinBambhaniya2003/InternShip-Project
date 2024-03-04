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

		titleWithCredits, err = services.Paginate(titleWithCredits, skip, limit, orderBy, order)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if selects == "" {
			for _, titleWithCredit := range titleWithCredits {
				fmt.Printf("Title: %s\n", titleWithCredit.Title.Title)
				fmt.Printf("Type: %s\n", titleWithCredit.Type)
				fmt.Printf("Description: %s\n", titleWithCredit.Description)
				// Add more fields as needed
				fmt.Println("Credits:")
				for _, credit := range titleWithCredit.Credits {
					fmt.Printf("- %s as %s\n", credit.Name, credit.Character)
				}
				fmt.Println("------------------------------------")
			}
			return
		}

		result := services.SelectColumn(titleWithCredits, selects)
		for _, record := range result {
			for key, value := range record {
				fmt.Printf("%s:%s ", key, value)
			}
			fmt.Println("\n")
		}

	},
}

func init() {
	titleCmd.AddCommand(TitlesWithCreditsCmd)
	TitlesWithCreditsCmd.Flags().IntVar(&skip, "skip", 0, "Skip the first N records")
	TitlesWithCreditsCmd.Flags().IntVar(&limit, "limit", -1, "Limit the number of records to M")
	TitlesWithCreditsCmd.Flags().StringVar(&selects, "selects", "", "Print only specified columns")
	TitlesWithCreditsCmd.Flags().StringVar(&order, "order", "", "Order records by ASC | DSC")
	TitlesWithCreditsCmd.Flags().StringVar(&orderBy, "order-by", "", "Define the column on which order is applied")
}
