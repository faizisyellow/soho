package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Option int

const (
	Repository Option = iota
	Service
	Handler
	Resource
)

func (opt Option) String() string {

	return [...]string{"repository", "service", "handler", "resource"}[opt]
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate {option} {name}",
	Example: "soho generate {repository | service | handler | resource} users",
	Aliases: []string{"g"},
	Args:    cobra.ExactArgs(2),
	Short:   "generate generates files to project",
	RunE: func(cmd *cobra.Command, args []string) error {

		option := args[0]

		switch option {
		case Repository.String():

			fmt.Println("you choose should be repo", option)

		case Service.String():
			fmt.Println("you choose should be service", option)

		case Handler.String():
			fmt.Println("you choose should be hanlder", option)

		case Resource.String():
			fmt.Println("you choose should be resource", option)

		default:
			return fmt.Errorf("invalid option, see examples to set the options")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generateCmd.Flags().BoolP("test", "t", false, "include test")
}
