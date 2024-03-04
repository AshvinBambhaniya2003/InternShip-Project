package cmd

import (
	"fmt"
	"log"
	"netflix/config"
	"netflix/models"
	"netflix/services"

	"github.com/spf13/cobra"
)

var uniqueActorsCmd = &cobra.Command{
	Use:   "unique-actors",
	Short: "Extract and display a list of unique actors from credits data.",
	Long: `unique-actors command extracts and displays a list of unique actors from the credits data stored in the provided CSV file.
	It reads the credits data, iterates over the actors, and stores them in a map to ensure uniqueness.
	Then, it extracts the unique actor names and prints them out.
	This command is useful for obtaining a comprehensive list of actors involved in the production of titles.`,
	Run: func(cmd *cobra.Command, args []string) {

		credits, err := models.ReadCredits(config.CreditFilePath)
		if err != nil {
			fmt.Println("Error reading credits:", err)
			return
		}

		actors := services.ListUniqueActors(credits)

		paginateActor, err := services.Paginate(actors, skip, limit, orderBy, order)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("List of Unique Actors:")
		for _, actor := range paginateActor {
			fmt.Println(actor.Name)
		}
	},
}

func init() {
	titleCmd.AddCommand(uniqueActorsCmd)
	uniqueActorsCmd.Flags().IntVar(&skip, "skip", 0, "Skip the first N records")
	uniqueActorsCmd.Flags().IntVar(&limit, "limit", -1, "Limit the number of records to M")
	uniqueActorsCmd.Flags().StringVar(&order, "order", "", "Order records by ASC | DSC")
	uniqueActorsCmd.Flags().StringVar(&orderBy, "order-by", "", "Define the column on which order is applied")
}
