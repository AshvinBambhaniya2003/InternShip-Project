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

		if len(args) > 0 {
			actorMap := services.ListMoviesForActor(titles, credits, args[0])
			fmt.Printf("Movies for actor %s:\n", args[0])
			for title, character := range actorMap {
				fmt.Printf("- %s as %s\n", title, character)
			}

		} else {
			log.Fatal("Only one argument allow")
		}

	},
}

func init() {
	titleCmd.AddCommand(MoviesForActorCmd)
}
