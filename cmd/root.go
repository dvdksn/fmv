package cmd

import (
	"fmt"
	"os"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "fmv",
		Short: "fmv - front matter validator",
		Long:  `fmv is a CLI tool for validating the YAML front matter metadata in markdown files.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) {
		// },
	}
)

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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "c", "config file (default is $HOME/.fmv.yaml)")

	// Get the Cue schema to use
	rootCmd.PersistentFlags().StringP("schema", "s", "frontmatter.cue", "schema file to use")
	viper.BindPFlag("schema", rootCmd.PersistentFlags().Lookup("schema"))
	// If set, use FMV_SCHEMA environment variable as default
	viper.SetEnvPrefix("fmv")
	viper.BindEnv("schema")
	viper.SetDefault("schema", viper.Get("schema"))

	// generate autocompletion
	carapace.Gen(rootCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".fmv" (without extension).
		viper.SetConfigType("yaml")
		viper.SetConfigName(".fmv")
		viper.AddConfigPath(home)
		viper.AddConfigPath(".") // optionally look for config in the working directory
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
