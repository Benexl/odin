package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"

	"github.com/Benexl/odin/app"
	"github.com/Benexl/odin/color"
	"github.com/Benexl/odin/filesystem"
	"github.com/Benexl/odin/icon"
	"github.com/Benexl/odin/util"
	"github.com/Benexl/odin/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

type clearTarget struct {
	name  string
	clear func() error
}

// Specify what can be cleared
var clearTargets = []clearTarget{
	{"cache", func() error {
		return filesystem.Api().RemoveAll(where.Cache())
	}},
	{"logs", func() error {
		return filesystem.Api().RemoveAll(where.Logs())
	}},
}

func init() {
	rootCmd.AddCommand(clearCmd)
	for _, n := range clearTargets {
		clearCmd.Flags().BoolP(n.name, string(n.name[0]), false, "clear "+n.name)
	}
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears sidelined files produced by the " + app.Name,
	Run: func(cmd *cobra.Command, args []string) {
		successStyle := lipgloss.NewStyle().Foreground(color.Green).Render
		var didSomething bool
		for _, n := range clearTargets {
			if lo.Must(cmd.Flags().GetBool(n.name)) {
				handleErr(n.clear())
				fmt.Printf("%s %s cleared\n", successStyle(icon.Check), util.Capitalize(n.name))
				didSomething = true
			}
		}

		if !didSomething {
			_ = cmd.Help()
		}
	},
}
