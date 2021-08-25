package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "playing_cobra_go",
	Short: "cobra setup",
	Long:  `Please use the 'playing_cobra_go' command to start the service with the appropriate flags.`,
}

var hellocmd = &cobra.Command{
	Use:   "hellocmd",
	Short: "says hello",
	Long:  `Please use the hellocmd command to start a greet service with the appropriate flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello ", name)
	},
}

var name = "default"

func init() {
	fmt.Println("In Init method ")
	RootCmd.AddCommand(hellocmd)

	hellocmd.Flags().StringVarP(&name, "name", "n", "default", "Source directory to read from")
}
