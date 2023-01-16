package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/richardcase/herald/internal/colorscheme"
	"github.com/richardcase/herald/internal/config"
	"github.com/richardcase/herald/internal/model/notifier"
	"github.com/spf13/cobra"
)

type options struct {
	colorSchemePath string
}

var (
	opts    = options{}
	rootCmd = &cobra.Command{
		Use:   "herald",
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
	rootCmd.Flags().StringVarP(&opts.colorSchemePath, "colorscheme_path", "s", "", "The path to the color scheme file")
}

func run() error {
	colors := colorscheme.New(opts.colorSchemePath)

	//TODO: add test colours command???

	// Initialize the cfg
	cfg, err := config.New(colors)
	if err != nil {
		return err
	}
	defer cfg.Close()

	// Create the notifier
	notifier := notifier.New(cfg)

	// Start the program
	p := tea.NewProgram(notifier)
	if _, err = p.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There has been an error executing the commands: '%s'", err)
		os.Exit(1)
	}
}
