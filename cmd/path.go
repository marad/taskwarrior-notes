package cmd

import (
	"fmt"
	"taskwarrior-notes/tw"

	"github.com/spf13/cobra"
)

var pathCmd = &cobra.Command{
	Use:   "path [filter] [flags]",
	Short: "Shows paths to task notes given a filter.",
	Long:  `The filter is a regular taskwarrior filter, e.g. "status:pending" or "project:myproject"`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tw.GetTasks(args)
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}
		for _, task := range tasks {
			path, err := tw.TaskNotePath(&task, "notes")
			if err != nil {
				fmt.Println("Error getting task note path:", err)
				return
			}
			fmt.Println(path)
		}
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)
}
