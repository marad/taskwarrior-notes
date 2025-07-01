package cmd

import (
	"fmt"
	"taskwarrior-notes/config"
	"taskwarrior-notes/tw"
	"taskwarrior-notes/util"

	"github.com/spf13/cobra"
)

var pathCmd = &cobra.Command{
	Use:   "path [filter] [flags]",
	Short: "Shows paths to task notes given a filter.",
	Long:  `The filter is a regular taskwarrior filter, e.g. "status:pending" or "project:myproject"`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := util.GetTaskFromFlagOrSearchTw(cmd, args)
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}

		paths, err := tw.GetTaskPaths(tasks, config.ReadNotesRoot())
		if err != nil {
			fmt.Println("Error getting task note paths:", err)
			return
		}
		for _, path := range paths {
			fmt.Println(path)
		}
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)
}
