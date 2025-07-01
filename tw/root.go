package tw

import (
	"encoding/json"
	"os/exec"
	"strings"
	"taskwarrior-notes/config"
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

func GetTaskPaths(filter []string) ([]string, error) {
	tasks, err := GetTasks(filter)
	if err != nil {
		return nil, err
	}

	notesRoot := config.ReadNotesRoot()
	var paths []string
	for _, task := range tasks {
		path, err := FindNoteByUUID(task.UUID)
		if err != nil {
			return nil, err
		} else if path != "" {
			paths = append(paths, path)
			continue
		}

		path, err = TaskNotePath(&task, notesRoot)
		if err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	return paths, nil
}

func FindNoteByUUID(uuid string) (string, error) {
	notesRoot := config.ReadNotesRoot()
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
