package util

import (
	"taskwarrior-notes/tw"

	"github.com/spf13/cobra"
)

func GetTaskFromFlagOrSearchTw(cmd *cobra.Command, filter []string) ([]tw.Task, error) {
	taskJson, _ := cmd.Flags().GetString("task")
	tasks := []tw.Task{}

	if taskJson != "" {
		task, err := tw.TaskFromString(taskJson)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, *task)
	} else {
		var err error
		tasks, err = tw.GetTasks(filter)
		if err != nil {
			return nil, err
		}
	}

	return tasks, nil
}
