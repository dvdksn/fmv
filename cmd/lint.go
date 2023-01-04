package cmd

import (
	"fmt"

	"github.com/dvdksn/fmv/markdown"
	"github.com/dvdksn/fmv/schema"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lint called")
		parser := markdown.New()
		validator := schema.Init(viper.GetString("schema"))
		for _, arg := range args {
			metadata, err := parser.GetMetadata(arg)
			if err != nil {
				panic(err)
			}
			valid := validator.Validate(metadata)

			fmt.Println(valid)
		}
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)
}
