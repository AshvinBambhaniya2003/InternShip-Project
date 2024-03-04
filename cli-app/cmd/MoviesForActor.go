package cmd

import (
	"fmt"
	"log"
	"netflix/config"
	"netflix/models"
	"netflix/services"

	"github.com/spf13/cobra"
)

// printMoviesForActorCmd represents the printMoviesForActor command
var MoviesForActorCmd = &cobra.Command{
	Use:   "movies-by-actor",
	Short: "print a list of movies for a given actor along with the characters they played.",
	Long: `printmoviesforactor is a command-line application that accepts the name of an actor as input and prints a list of movies in which the actor worked, along with the characters they played.
	It reads the credits data from the provided CSV file, searches for movies involving the specified actor, and outputs the movie titles and corresponding character names.
	This command is useful for obtaining a comprehensive list of movies a particular actor participated in and the roles they portrayed.`,
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

		if len(args) != 1 {
			log.Fatal("Only one argument allow")
		}

		actorMovies := services.ListMoviesForActor(titles, credits, args[0])
		fmt.Printf("Movies for actor %s:\n", args[0])

		actorMovies, err = services.Paginate(actorMovies, skip, limit, orderBy, order)
		if err != nil {
			log.Fatal(err.Error())
		}

		if selects == "" {
			for _, value := range actorMovies {
				fmt.Printf("TitleName:%s CharacterName:%s\n", value.Title, value.CharacterName)
			}
			return
		}

		result := services.SelectColumn(actorMovies, selects)
		for _, record := range result {
			for key, value := range record {
				fmt.Printf("%s:%s ", key, value)
			}
			fmt.Println("")
		}

	},
}

func init() {
	titleCmd.AddCommand(MoviesForActorCmd)
	MoviesForActorCmd.Flags().IntVar(&skip, "skip", 0, "Skip the first N records")
	MoviesForActorCmd.Flags().IntVar(&limit, "limit", -1, "Limit the number of records to M")
	MoviesForActorCmd.Flags().StringVar(&selects, "selects", "", "Print only specified columns")
	MoviesForActorCmd.Flags().StringVar(&order, "order", "", "Order records by ASC | DSC")
	MoviesForActorCmd.Flags().StringVar(&orderBy, "order-by", "", "Define the column on which order is applied")
}
