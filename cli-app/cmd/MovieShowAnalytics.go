package cmd

import (
	"fmt"
	"netflix/config"
	"netflix/models"

	"github.com/spf13/cobra"
)

var moviesCountByReleaseYear, moviesCountByAgeCertificate, movieCountByRuntime, titleAnalysisByGenres, titleAnalysisByCountry, titleCountBySeason, movieShowCountWithPercentage, mostWorkingActor, titleCountByImdbScore, titleCountByRuntime bool

// printMoviesForActorCmd represents the printMoviesForActor command
var MovieShowAnalyticsCmd = &cobra.Command{
	Use:   "movie-show-analytics",
	Short: "Perform analytics on a dataset of movies and shows.",
	Long: `The movie-show-analytics command allows users to analyze a dataset of movies and shows.
	It offers various analytical tasks such as counting movies by release year, age certification, runtime, genres, country, IMDB score, and more.
	Users can specify which analyses they want to perform using command-line options. The results of the analysis are displayed to the user, providing insights into the characteristics and distribution of movies and shows in the dataset.`,
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

		if moviesCountByReleaseYear {
			countRecords := models.ListMoviesCountByReleaseYear(titles)

			fmt.Println("Movies count with the release year:")
			for year, count := range countRecords {
				fmt.Printf("%d: %d\n", year, count)
			}
		}

		if moviesCountByAgeCertificate {
			countRecords := models.ListMoviesCountByAgeCertificate(titles)

			fmt.Println("Movies count with Age certificates:")
			for year, count := range countRecords {
				fmt.Printf("%s: %d\n", year, count)
			}
		}

		if movieCountByRuntime {
			titlesCountByRuntimeMap := models.ListMovieCountByRuntime(titles)

			for threshold, count := range titlesCountByRuntimeMap {
				fmt.Printf("%d movies runtimes are less than %d minutes.\n", count, threshold)
			}
		}

		if titleAnalysisByGenres {
			genreCount, totalTitles := models.ListTitlesCountPercentageByGenres(titles)
			for genre, count := range genreCount {
				percentage := float64(count) / float64(totalTitles) * 100
				fmt.Printf("%s: %d titles (%.2f%%)\n", genre, count, percentage)
			}
		}

		if titleAnalysisByCountry {
			countryCount, totalTitles := models.ListTitlesCountPercentageByCountry(titles)
			for country, count := range countryCount {
				percentage := float64(count) / float64(totalTitles) * 100
				fmt.Printf("%s: %d titles (%.2f%%)\n", country, count, percentage)
			}
		}

		if titleCountBySeason {
			seasonsCounts := models.ListTitleCountBySeasons(titles)
			for season, count := range seasonsCounts {
				fmt.Printf("%d titles have %d sessions\n", count, season)
			}
		}

		if titleCountByImdbScore {
			titlesCountByImdbMap := models.ListTitlesCountByIMDbScore(titles)

			for thresholdArray, count := range titlesCountByImdbMap {
				fmt.Printf("%d titles IMDB rating is greater or equal to %d and less than %d \n", count, thresholdArray[0], thresholdArray[1])
			}
		}

		if titleCountByRuntime {
			titlesCountByRuntimeMap := models.ListTitlesCountByRuntime(titles)

			for threshold, count := range titlesCountByRuntimeMap {
				fmt.Printf("%d titles have greater runtime than %d minutes.\n", count, threshold)
			}
		}

		if movieShowCountWithPercentage {
			movieCount, moviePercentage, showCount, showPercentage := models.GetTitleTypeCountsAndPercentages(titles)
			fmt.Printf("Movie count: %d (%.2f%%)\n", movieCount, moviePercentage)
			fmt.Printf("Show count: %d (%.2f%%)\n", showCount, showPercentage)
		}

		if mostWorkingActor {
			mostWorkingActor := models.GetMostWorkingActor(credits)

			fmt.Println("Most working actor:", mostWorkingActor)
		}
	},
}

func init() {
	titleCmd.AddCommand(MovieShowAnalyticsCmd)
	MovieShowAnalyticsCmd.Flags().BoolVar(&moviesCountByReleaseYear, "movies-count-by-release-year", false, "Count movies with release year")
	MovieShowAnalyticsCmd.Flags().BoolVar(&moviesCountByAgeCertificate, "movies-count-by-age-certificates", false, "Count movies by age certificate")
	MovieShowAnalyticsCmd.Flags().BoolVar(&movieCountByRuntime, "movie-count-by-runtime", false, "Count of movie by runtime")
	MovieShowAnalyticsCmd.Flags().BoolVar(&titleAnalysisByGenres, "title-analysis-by-genres", false, "Perform genres-wise count and percentage analysis")
	MovieShowAnalyticsCmd.Flags().BoolVar(&titleAnalysisByCountry, "title-analysis-by-country", false, "Perform country-wise count and percentage analysis")
	MovieShowAnalyticsCmd.Flags().BoolVar(&titleCountBySeason, "title-count-by-season", false, "Perform session count analysis")
	MovieShowAnalyticsCmd.Flags().BoolVar(&titleCountByImdbScore, "title-count-by-imdb-score", false, "Perform movie and show count analysis")
	MovieShowAnalyticsCmd.Flags().BoolVar(&titleCountByRuntime, "title-count-by-runtime", false, "Threshold for Title runtime")
	MovieShowAnalyticsCmd.Flags().BoolVar(&movieShowCountWithPercentage, "movie-show-count-with-percentage", false, "Perform movie and show count analysis")
	MovieShowAnalyticsCmd.Flags().BoolVar(&mostWorkingActor, "most-working-actor", false, "Find the most working actor")

}
