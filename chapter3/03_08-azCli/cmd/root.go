package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	subscriptionId    string
	resourceGroupName string
	name              string
	location          string
	all               bool
)

var template = `
_____  ________
\__  \ \___   /
 / __ \_/    /    A re-imagined Azure CLI
(____  /_____ \
     \/      \/

{{ .CommandPath }}: {{ .Long }}
{{ if .HasAvailableSubCommands}}
Usage: {{ .CommandPath }} [COMMAND] [FLAGS]

Commands: 
{{ range .Commands }}
  {{ .NameAndAliases }} -  {{ .Short }}
{{ end }}
{{ else }}
Usage: {{ .CommandPath }} [FLAGS]
{{ end }}
{{ if .HasAvailableFlags}}
Flags:
{{.Flags.FlagUsages | trimTrailingWhitespaces}}
{{ end }}

{{ if .HasSubCommands}}
Use "{{ .CommandPath }} [COMMAND] --help" for more information about a command.
{{ end }}
`

var rootCmd = &cobra.Command{
	Use:                        "az",
	Short:                      "A re-imagined Azure CLI",
	Long:                       "Azure CLI written in Go language using the Cobra package.",
	SuggestionsMinimumDistance: 1,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddGroup(&cobra.Group{ID: "read", Title: "Read Commands"})
	rootCmd.AddGroup(&cobra.Group{ID: "create", Title: "Create Commands"})
	rootCmd.AddGroup(&cobra.Group{ID: "delete", Title: "Delete Commands"})

	rootCmd.PersistentFlags().StringVarP(&subscriptionId, "subscriptionId", "s", os.Getenv("AZURE_SUBSCRIPTION_ID"), "ID of the Azure subscription where the Azure resources are provisioned.")

	rootCmd.SetHelpTemplate(template)
	rootCmd.SetUsageTemplate(template)
}
