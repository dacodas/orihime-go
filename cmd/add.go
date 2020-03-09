package cmd

import (
	"os"
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"orihime/internal"

	"github.com/spf13/cobra"
)

var (
	Source string
	Word string
	ChildWord string
	Definition string
	Text string
	Input string
)

// https://stackoverflow.com/questions/22744443/check-if-there-is-something-to-read-on-stdin-in-golang
func pipedToStdin() bool {
	state, _ := os.Stdin.Stat()

	if ( state.Mode() & os.ModeCharDevice ) == 0 {
		fmt.Println("There is data being piped in to stdin")
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
		fmt.Println("Error: Input was provided neither from stdin nor as an argument")
		os.Exit(1)
	case 1:
		fmt.Println("Using argument to add orihime object...")
		Input = args[1]
	case 2:
		fmt.Println("Using stdin to add orihime object...")
		buf, err := ioutil.ReadFile("/dev/stdin")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		Input = string(buf)
	case 3:
		fmt.Println("Error: Input was piped to stdin and an argument was passed to add")
		os.Exit(1)
	}

	fmt.Println("Input status is ", inputStatus)
}

func ensureDefinition() { }
func ensureSource() { }
func ensureParentTextHash() { }

func determineWhatToAdd(args []string) {
	switch args[0] {
	case "source":
		internal.AddSource(args[1])
	case "word":
		ensureDefinition()
		ensureSource()
		internal.AddWord(Input, Definition, Source)
	case "childword":
		ensureDefinition()
		ensureSource()
		ensureParentTextHash()
	case "text":
		ensureSource()
	default:
		fmt.Println("Unknown orihime object to add: " + args[0])
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Insert orihime objects into persistent datastore",
	Long: "",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called with " + strings.Join(args, " "))
		checkStdinOrAnotherArgument(args)
		determineWhatToAdd(args)
	},
}

// Allow to pipe in from stdin
func init() {
	addCmd.Flags().StringVarP(&Source, "source", "s", "", "the source to add or the source associated with the object being added")
	addCmd.Flags().StringVarP(&Word, "word", "w", "", "the word to add")
	addCmd.Flags().StringVarP(&ChildWord, "childword", "c", "", "the child word to add")
	addCmd.Flags().StringVarP(&Definition, "definition", "d", "", "the definition to add")
	addCmd.Flags().StringVarP(&Text, "text", "t", "", "the text to add")

	rootCmd.AddCommand(addCmd)
}
