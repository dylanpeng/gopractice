package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "the short description shown in the 'help' output.",
	Long: `root description detail
			Long is the long message shown in the 'help <this-command>' output.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Printf("root verbose: %v\n", Verbose)
		fmt.Printf("root config: %s\n", ConfigPath)
		fmt.Printf("root string: %s\n", SomeString)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
