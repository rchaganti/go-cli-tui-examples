package cmd

import (
	"fmt"
	helper "gla/internal"
	"gla/pkg/gai"
	"gla/ui"
	"log"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	storePath     = ".gla"
	storeFileName = "sessions.json"
)

var (
	apikey      string
	model       string
	sessionid   string
	listSession bool
	session     helper.Session
	se          []helper.Session
	err         error
)

var rootCmd = &cobra.Command{
	Use:   "gla",
	Short: "A Learning Assistant that helps you chat with a generative AI model to learn, explore, and experiment",
	Run: func(cmd *cobra.Command, args []string) {
		ss := helper.SessionStore{
			Path:     storePath,
			FileName: storeFileName,
		}

		se, err = ss.GetSessionEntries()
		if err != nil {
			log.Fatal(err)
		}

		// List existing saved sessions can be executed without an API key
		listSession = viper.GetBool("list-session")
		if listSession {
			if len(se) > 0 {
				for _, s := range se {
					fmt.Printf("%s\n", s.Id)
				}
			} else {
				fmt.Println("No saved sessions found.")
			}

			os.Exit(0)
		}

		// handle flags
		apikey := viper.GetString("api-key")
		if apikey == "" {
			cmd.Help()
			os.Exit(1)
		}

		model = viper.GetString("model")

		sessionid = viper.GetString("session-id")
		if sessionid == "" {
			sessionid = strings.Split(uuid.New().String(), "-")[0]
		}

		session, err = ss.GetSessionEntry(sessionid)
		if session.Id == "" {
			session = helper.Session{
				Id:         sessionid,
				Model:      model,
				Timestamp:  time.Now().Format("2006-01-02 15:04"),
				IsSelected: true,
				Messages: []gai.MessageContent{
					{
						Role:    "model",
						Message: gai.WelcomeMessage,
					},
				},
				TokenCount: 0,
			}

			_, err = ss.AddSession(session)
			if err != nil {
				log.Fatal(err)
			}
		}

		session, err = ss.SelectSession(sessionid)
		if err != nil {
			log.Fatal(err)
		}

		p := tea.NewProgram(
			ui.InitialChatModel(apikey, session, ss),
			tea.WithAltScreen(),
		)
		if _, err = p.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	// Initialize the store
	_, err := helper.InitializeSessionStore(storePath, storeFileName)
	if err != nil {
		log.Fatal(err)
	}

	// Define the flags
	rootCmd.Flags().StringVarP(&apikey, "api-key", "a", "", "API key to interact with the Gemini AI. Alternatively, you can set GEMINI_API_KEY envionment variable.")
	viper.BindPFlag("api-key", rootCmd.Flags().Lookup("api-key"))
	viper.BindEnv("api-key", "GEMINI_API_KEY")

	rootCmd.Flags().StringVarP(&model, "model", "m", "gemini-1.5-flash", "A generative AI model to use for this session. Alternatively, you can set GEMINI_MODEL environment variable.")
	viper.BindPFlag("model", rootCmd.Flags().Lookup("model"))
	viper.BindEnv("model", "GEMINI_MODEL")

	rootCmd.Flags().StringVarP(&sessionid, "session-id", "s", "", "Creates new session or loads a saved session from history. A random ID is generated when this flag is not specified.")
	viper.BindPFlag("session-id", rootCmd.Flags().Lookup("session-id"))

	rootCmd.Flags().BoolVarP(&listSession, "list-session", "l", false, "List existing saved sessions.")
	viper.BindPFlag("list-session", rootCmd.Flags().Lookup("list-session"))

	//_ = rootCmd.MarkFlagRequired("session")
	rootCmd.MarkFlagsMutuallyExclusive("model", "list-session")
	rootCmd.MarkFlagsMutuallyExclusive("session-id", "list-session")
}

func Execute() error {
	return rootCmd.Execute()
}
