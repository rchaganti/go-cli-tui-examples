package helper

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

const banner = `          
_____  ________
\__  \ \___   /
 / __ \_/    /    A re-imagined Azure CLI
(____  /_____ \
     \/      \/
`

func printBanner() {
	fmt.Printf("%s\n", banner)
}

func printHelpAndUsage(cmd *cobra.Command) {
	printBanner()
	hasSubCommands := false
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

	cmdPath := cmd.CommandPath()
	fmt.Fprintf(writer, "%s\t: %s\n", cmdPath, cmd.Long)

	if cmd.HasAvailableSubCommands() {
		hasSubCommands = true
		fmt.Fprintf(writer, "Usage\t: %s [COMMAND] [flags]\n", cmdPath)
		fmt.Fprintf(writer, "\nCommands:\n")
		for _, c := range cmd.Commands() {
			fmt.Fprintf(writer, "  %s\t%s\n", c.NameAndAliases(), c.Short)
		}
	} else {
		fmt.Fprintf(writer, "Usage\t: %s [flags]\n", cmdPath)
	}

	if cmd.Flags().HasAvailableFlags() {
		fmt.Fprintf(writer, "\nFlags\t:\n%s\n", cmd.Flags().FlagUsages())
	}

	if hasSubCommands {
		fmt.Fprintf(writer, "Use \"%s [COMMAND] --help\" for more information about a command.\n", cmdPath)
	}
	writer.Flush()
}
func UsageFunc(cmd *cobra.Command) error {
	printHelpAndUsage(cmd)
	return nil
}

func HelpFunc(cmd *cobra.Command, args []string) {
	printHelpAndUsage(cmd)
}
