package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/afero"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// NewRootCmd generate root cmd
func NewRootCmd(fs afero.Fs) (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:           "github_starred",
		Short:         "github_starred",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	if err := registerSubCommands(fs, cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

func registerSubCommands(fs afero.Fs, cmd *cobra.Command) error {
	var subCmds []*cobra.Command
	for _, cmdGen := range cmdGenerators {
		subCmd, err := cmdGen(fs)
		if err != nil {
			return err
		}
		subCmds = append(subCmds, subCmd)
	}
	cmd.AddCommand(subCmds...)
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd, err := NewRootCmd(afero.NewOsFs())
	if err != nil {
		panic(err)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
