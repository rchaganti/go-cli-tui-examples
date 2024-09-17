package cmd

import (
	"gla/pkg/gai"
	"gla/ui"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apikey         string
	model          string
	messageHistory []gai.MessageContent
)

var rootCmd = &cobra.Command{
	Use:   "gla",
	Short: "A Learning Assistant that helps you chat with a generative AI model to learn, explore, and experiment",
	Run: func(cmd *cobra.Command, args []string) {
		apikey := viper.GetString("api-key")
		model := viper.GetString("model")

		messageHistory = append(messageHistory, gai.MessageContent{
			Message: string(gai.WelcomeMessage),
			Role:    "model",
		})

		if apikey == "" {
			cmd.Help()
			os.Exit(1)
		}

		p := tea.NewProgram(
			ui.InitialChatModel(apikey, model, messageHistory),
			tea.WithAltScreen(),
		)

		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&apikey, "api-key", "a", "", "API key to interact with the Gemini AI.")
	viper.BindPFlag("api-key", rootCmd.Flags().Lookup("api-key"))
	viper.BindEnv("api-key", "GEMINI_API_KEY")

	rootCmd.Flags().StringVarP(&model, "model", "m", "gemini-1.5-flash", "A generative AI model to use for this session. Default is gemini-1.5-flash.")
	viper.BindPFlag("model", rootCmd.Flags().Lookup("model"))
	viper.BindEnv("model", "GEMINI_MODEL")
}

func Execute() error {
	return rootCmd.Execute()
}
