package tw

import (
	"encoding/json"
	"os/exec"
	"strings"
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

func GetTaskPaths(filter []string, notesRoot string) ([]string, error) {
	tasks, err := GetTasks(filter)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, task := range tasks {
		path, err := GetTaskPath(&task, notesRoot)
		if err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	return paths, nil
}

func GetTaskPath(task *Task, notesRoot string) (string, error) {

	// Try to find existing note by UUID first (in case it's name was changed)
	path, err := FindNoteByUUID(task.UUID, notesRoot)
	if err != nil {
		return "", err
	} else if path != "" {
		return path, nil
	}

	// If no note found by UUID, generate path based on task details
	path, err = TaskNotePath(task, notesRoot)
	if err != nil {
		return "", err
	}

	return path, nil
}

func FindNoteByUUID(uuid string, notesRoot string) (string, error) {
	cmd := exec.Command("grep", "-l", "-r", "--include", "*.md", "uuid: "+uuid, notesRoot)
	out, err := cmd.Output()
	if err != nil {
		// If grep doesn't find anything, it returns an error. We don't want to return an error in that case.
		if strings.Contains(err.Error(), "exit status 1") {
			return "", nil
		}
		return "", err
	}

	if len(out) == 0 {
		return "", nil
	}

	return strings.Trim(string(out), "\n\t "), nil
}
