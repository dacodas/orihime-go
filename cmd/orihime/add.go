package main

import (
	"os"
	"log"
	"strings"
	"encoding/hex"
	"io/ioutil"
	// "orihime/internal/database"
	"orihime/internal/client"

	"github.com/spf13/cobra"
)

var (
	Source string
	Word string
	ChildWord string
	Definition string
	ParentTextHashHexString string
	ParentTextHash []uint8
	Input string
	User string
)

// https://stackoverflow.com/questions/22744443/check-if-there-is-something-to-read-on-stdin-in-golang
func pipedToStdin() bool {
	state, _ := os.Stdin.Stat()

	if ( state.Mode() & os.ModeCharDevice ) == 0 {
		log.Printf("There is data being piped in to stdin")
		return true
	} else {
		return false
	}
}

func checkStdinOrAnotherArgument(args []string) {
	var inputStatus uint8 = 0

	if len(args) == 2 {
		inputStatus |= 1
	}

	if pipedToStdin() {
		inputStatus |= ( 1 << 1 )
	}

	switch inputStatus {
	case 0:
		log.Printf("Error: Input was provided neither from stdin nor as an argument")
		os.Exit(1)
	case 1:
		log.Printf("Using argument to add orihime object...")
		Input = args[1]
	case 2:
		log.Printf("Using stdin to add orihime object...")
		buf, err := ioutil.ReadFile("/dev/stdin")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		Input = string(buf)
	case 3:
		log.Printf("Error: Input was piped to stdin and an argument was passed to add")
		os.Exit(1)
	}

	log.Printf("Input status is ", inputStatus)
}

func exitNotSpecified(unspecifiedObject string) {
	log.Printf("Error: %v not specified", unspecifiedObject)
	os.Exit(1)
}

func ensureUser() {
	if User == "" {
		exitNotSpecified("user")
	}
}

func ensureDefinition() {
	if Definition == "" {
		exitNotSpecified("definition")
	}
}

func ensureSource() {
	if Source == "" {
		exitNotSpecified("source")
	}
}

func ensureParentTextHash() {
	if ParentTextHashHexString == "" {
		exitNotSpecified("parent-text")
	}

	var hashErr error
	ParentTextHash, hashErr = hex.DecodeString(ParentTextHashHexString)
	if hashErr != nil {
		log.Fatal(hashErr)
		os.Exit(1)
	}

	log.Printf("Using hash %v, %v", ParentTextHashHexString, ParentTextHash)
}

func determineWhatToAdd(args []string) {
	switch args[0] {
	case "source":
		client.AddSource(Input)
	case "word":
		ensureDefinition()
		ensureSource()
		client.AddWord(Input, Definition, Source)
	case "child-word":
		ensureDefinition()
		ensureSource()
		ensureParentTextHash()
		ensureUser()
		client.AddChildWord(Input, Definition, Source, User, ParentTextHash)
	case "text":
		ensureSource()
		client.AddText(Input, Source)
	default:
		log.Printf("Unknown orihime object to add: %v", args[0])
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Insert orihime objects into persistent datastore",
	Long: "",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("add called with %v", strings.Join(args, " "))
		checkStdinOrAnotherArgument(args)
		determineWhatToAdd(args)
	},
}

// Allow to pipe in from stdin
func init() {
	addCmd.Flags().StringVarP(&Source, "source", "s", "", "the source to add or the source associated with the object being added")
	addCmd.Flags().StringVarP(&Word, "word", "w", "", "the word to add")
	addCmd.Flags().StringVarP(&Definition, "definition", "d", "", "the definition to add")
	addCmd.Flags().StringVarP(&ParentTextHashHexString, "parent-text", "p", "", "the parent text")
	addCmd.Flags().StringVarP(&User, "user", "u", "", "the user to use for this transaction")

	rootCmd.AddCommand(addCmd)
}
