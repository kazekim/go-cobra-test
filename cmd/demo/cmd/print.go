/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var echoTimes int

func AddPrintCommand(parentCommand *cobra.Command) {

	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	cmdPrint.PersistentFlags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	parentCommand.AddCommand(cmdPrint)
}
