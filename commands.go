package main

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/spf13/cobra"
)

func mainCommands() {
	rootCmd := &cobra.Command{
		Use: "openlandings",
		Run: func(cmd *cobra.Command, args []string) {
			beego.Run()
		},
	}

	cmdWeb := &cobra.Command{
		Use:   "web",
		Short: "Start the web server",
		Run: func(cmd *cobra.Command, args []string) {
			beego.Run()
		},
	}

	cmdSiteCollector := &cobra.Command{
		Use:   "collect",
		Short: "Collect the site templates from the host",
		Run: func(cmd *cobra.Command, args []string) {
			// collector.Run()
		},
	}

	rootCmd.AddCommand(cmdWeb, cmdSiteCollector)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
