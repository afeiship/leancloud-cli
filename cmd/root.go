package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lcc/misc"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lcc",
	Short: "Leancloud cli for options curd.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
		// get optionKey
		optionKey, _ := cmd.Flags().GetString("optionKey")
		// Get value by optionKey
		if optionKey != "" {
			fmt.Println("optionKey:", optionKey)
			res := misc.LcGet(optionKey)
			fmt.Println("res:", res)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lcc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// get value by optionKey
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("optionKey", "k", "", "Get value by option key")
}
