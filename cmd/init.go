package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"

	"github.com/spf13/cobra"
)

type Config struct {
	Handler    string
	Router     string
	Service    string
	Repository string
}

const FILENAME = ".soho.toml"

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Args:  cobra.NoArgs,
	Short: "init will create initialization for Soho CLI to generate files.",
	RunE: func(cmd *cobra.Command, args []string) error {

		root, err := os.Getwd()
		if err != nil {
			return err
		}

		cfg := Config{
			Handler:    "/cmd/api",
			Router:     "/cmd/api",
			Service:    "/internal/service",
			Repository: "/internal/repository",
		}

		tomByte, err := toml.Marshal(cfg)
		if err != nil {
			return err
		}

		f, err := os.Create(filepath.Join(root, FILENAME))
		if err != nil {
			return err
		}

		_, err = f.Write(tomByte)
		if err != nil {
			return err
		}

		fmt.Println("initialize Soho CLI at ", root)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
