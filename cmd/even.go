package cmd

import (
	"github.com/ashikov/brain-games/src/games"
	"github.com/spf13/cobra"
)

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "You need to answer whether a given number is even or not",
	Long: `You need to answer whether a given number is even or not`,
	Run: func(cmd *cobra.Command, args []string) {
		games.RunEven()
	},
}

func init() {
	runCmd.AddCommand(evenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
