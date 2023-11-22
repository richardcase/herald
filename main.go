package main

import (
	"fmt"
	slog "log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"

	"github.com/richardcase/herald/internal/ui"
)

type options struct {
	Debug      bool
	ConfigFile string
}

var (
	opts    = options{}
	rootCmd = &cobra.Command{
		Use: "herald",
		//Version: = getVersion(),
		Short: "herald is a command line utility for receiving GitHub notifications",
		Run: func(cmd *cobra.Command, args []string) {
			if err := run(); err != nil {
				fmt.Fprintf(os.Stderr, "There was an error executing herald: %w", err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&opts.ConfigFile, "config", "c", "", "the config file to use")
	rootCmd.Flags().BoolVarP(&opts.Debug, "debug", "d", false, "If true this will write debugging information to debug.log")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There has been an error executing the commands: '%s'", err)
		os.Exit(1)
	}
}

func run() error {
	lipgloss.SetHasDarkBackground(termenv.HasDarkBackground())
	slog.Println("Starting")
	//markdown.InitializeMarkdownStyle(termenv.HasDarkBackground())

	//TODO: add test colours command???

	// Initialize the cfg
	// cfg, err := config.New(colors)
	// if err != nil {
	// 	return err
	// }
	// defer cfg.Close()

	// Create the notifier
	//notifier := notifier.New(cfg)

	// Start the program
	//p := tea.NewProgram(notifier)

	model, logger := createModel(opts.ConfigFile, opts.Debug)
	if logger != nil {
		defer logger.Close()
	}

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}

func createModel(configPath string, debug bool) (ui.Model, *os.File) {
	var logFile *os.File

	if debug {
		newConfigFile, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(newConfigFile)
			log.SetTimeFormat(time.Kitchen)
			log.SetReportCaller(true)
			log.SetLevel(log.DebugLevel)
		} else {
			logFile, _ = tea.LogToFile("debug.log", "debug")
			slog.Print("Failed setting up logging", err)
		}
	}

	return ui.NewModule(configPath), logFile
}
