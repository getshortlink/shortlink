package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// EnvironmentVariablePrefix is the prefix for environment variables
	// that are used to configure the application.
	EnvironmentVariablePrefix = "shortlink"
)

var rootCmd = &cobra.Command{
	Use:  "shortlink <command> <subcommand> [flags]",
	Long: "Shortlink is a URL shortening service.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := initializeConfig(cmd); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//
	// Add commands to the root command.
	//
	rootCmd.AddCommand(NewServerCommand().Command)

	// Add common flags to the root command.
	// To access these flags from a child command, you can use the Flags directly or
	// use the command.Config struct. See command/command.go for an example.
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --number
	// binds to an environment variable SHORTLINK_NUMBER. This helps
	// avoid conflicts.
	v.SetEnvPrefix(EnvironmentVariablePrefix)

	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --number to SHORTLINK_NUMBER
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Bind to environment variables.
	v.AutomaticEnv()

	// Bind the current command's flags to viper
	bindFlags(cmd, v)

	return nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))

		}
	})
}
