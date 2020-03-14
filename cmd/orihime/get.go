package main

import (
	"encoding/hex"
	"log"
	"os"
	"github.com/spf13/cobra"

	"orihime/internal/client"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Query orihime objects",
	Long: "",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkStdinOrAnotherArgument(args)
		ensureUser()

		var hashErr error
		ParentTextHash, hashErr = hex.DecodeString(Input)
		if hashErr != nil {
			log.Fatal(hashErr)
			os.Exit(1)
		}

		client.GetTextTree(ParentTextHash, User)
	},
}

// Allow to pipe in from stdin
func init() {
	getCmd.Flags().StringVarP(&User, "user", "u", "", "the user to use for this transaction")

	rootCmd.AddCommand(getCmd)
}
