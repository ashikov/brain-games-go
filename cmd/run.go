package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts the game",
	Long:  "Starts the game passed as an argument",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}

			os.Exit(0)
        }
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
