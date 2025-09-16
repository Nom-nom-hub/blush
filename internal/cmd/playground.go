package cmd

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/nom-nom-hub/blush/internal/tui/components/playground"
	"github.com/spf13/cobra"
)

var playgroundCmd = &cobra.Command{
	Use:   "playground",
	Short: "Launch the interactive code playground",
	Long: `Launch the interactive code playground for experimenting with code snippets.
The playground supports multiple languages including Go, JavaScript, and Python.`,
	Example: `
# Launch the playground
blush playground
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create a new playground model
		model := playground.New("chat")

		// Create a new program with the playground model
		p := tea.NewProgram(model, tea.WithAltScreen())

		// Run the program
		_, err := p.Run()
		if err != nil {
			return err
		}

		// The program will exit when the user presses ESC or Ctrl+C
		// ESC will send ExitMsg and Ctrl+C will quit the program
		return nil
	},
}

func init() {
	rootCmd.AddCommand(playgroundCmd)
}