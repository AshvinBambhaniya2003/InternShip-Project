package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var titleCmd = &cobra.Command{
	Use:   "title",
	Short: "Manages titles in the database or CSV file.",
	Long: `Allows users to manage titles stored in the database or CSV file.
	It provides functionalities for listing all titles, 
	Users can also perform various operations on titles, such as sorting, filtering, and pagination, to effectively manage and analyze the title data.
	Additionally, the command may include options for displaying titles with associated person counts.`,
}


func Execute() {
	err := titleCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


