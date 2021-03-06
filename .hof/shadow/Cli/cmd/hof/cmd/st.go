package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/cmd/st"

	"github.com/hofstadter-io/hof/cmd/hof/ga"
)

var stLong = `Structural diff, merge, mask, pick, and query helpers for Cue

Commands generally have the form: <method> <op1> <op2> [...entrypoints]

Where <op> can be:
- Cue: expr: as: string
- @filename.cue: Cue: expr: as: string

If entrypoints are supplied, then an <op> without an @filename.cue will lookup from the entrypoints.
Otherwise, the <op> is interpreted as a complete Cue value.`

var StCmd = &cobra.Command{

	Use: "st",

	Aliases: []string{
		"structural",
	},

	Short: "recursive diff, merge, mask, pick, and query helpers for Cue",

	Long: stLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendCommandPath(cmd.CommandPath())

	},
}

func init() {

	help := StCmd.HelpFunc()
	usage := StCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		ga.SendCommandPath(cmd.CommandPath() + " help")
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		ga.SendCommandPath(cmd.CommandPath() + " usage")
		return usage(cmd)
	}
	StCmd.SetHelpFunc(thelp)
	StCmd.SetUsageFunc(tusage)

	StCmd.AddCommand(cmdst.DiffCmd)
	StCmd.AddCommand(cmdst.MergeCmd)
	StCmd.AddCommand(cmdst.PickCmd)
	StCmd.AddCommand(cmdst.MaskCmd)
	StCmd.AddCommand(cmdst.QueryCmd)

}
