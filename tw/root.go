package tw

import (
	"encoding/json"
	"os/exec"
)

func GetTasks(filter []string) ([]Task, error) {
	args := append(filter, "export")
	cmd := exec.Command("task", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(out, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
