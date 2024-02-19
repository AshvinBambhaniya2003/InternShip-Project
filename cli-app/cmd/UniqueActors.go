package cmd

import (
	"fmt"
	"netflix/models"
	"netflix/services"

	"github.com/spf13/cobra"
)

var creditfilepath = "CSV/credits.csv"

var uniqueActorsCmd = &cobra.Command{
	Use:   "unique-actors",
	Short: "Extract and display a list of unique actors from credits data.",
	Long: `unique-actors command extracts and displays a list of unique actors from the credits data stored in the provided CSV file.
	It reads the credits data, iterates over the actors, and stores them in a map to ensure uniqueness.
	Then, it extracts the unique actor names and prints them out.
	This command is useful for obtaining a comprehensive list of actors involved in the production of titles.`,
	Run: func(cmd *cobra.Command, args []string) {

		credits, err := models.ReadCredits(creditfilepath)
		if err != nil {
			fmt.Println("Error reading credits:", err)
			return
		}

		actors := services.ListUniqueActors(credits)

		fmt.Println("List of Unique Actors:")
		for _, actor := range actors {
			fmt.Println(actor)
		}
	},
}

func init() {
	titleCmd.AddCommand(uniqueActorsCmd)
}