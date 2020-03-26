package container

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

var resetLong = `Reset a Studios container.`

var ResetCmd = &cobra.Command{

	Use: "reset <name or id>",

	Short: "Reset a Studios container.",

	Long: resetLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'Ident'")
			cmd.Usage()
			os.Exit(1)
		}

		var ident string

		if 0 < len(args) {

			ident = args[0]

		}

		// Default body

		fmt.Println("hof studios container reset", ident)

	},
}
