package cmd

import (
	"fmt"
	"os"

	"github.com/PatrickJSli/clickcounter/app"
	"github.com/PatrickJSli/clickcounter/bubble"
	"github.com/PatrickJSli/clickcounter/xinput"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	return &cobra.Command{
		Use:     "clickcounter",
		Short:   "clickcounter is a CLI for counting clicks",
		Version: app.Version,
		Run: func(cmd *cobra.Command, args []string) {

			if !xinput.XinputInstalled() {
				fmt.Println("xinput is not installed. Please install xinput and try again.")
				os.Exit(1)
			}

			bubble.Run()

		},
	}
}
