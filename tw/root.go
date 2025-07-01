package tw

import (
	"encoding/json"
	"os/exec"
)

func GetTasks(filter string) ([]Task, error) {
	cmd := exec.Command("task", "export", filter)
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
