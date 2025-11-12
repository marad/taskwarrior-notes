package tw

import (
	"encoding/json"
	"os/exec"
	"strings"
)

// NoteFinder abstracts locating an existing note by task UUID.
type NoteFinder func(uuid, notesRoot string) (string, error)

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

func GetTaskPathsWithFinder(tasks []Task, notesRoot string, finder NoteFinder) ([]string, error) {
	var paths []string
	for _, task := range tasks {
		p, err := GetTaskPathWithFinder(&task, notesRoot, finder)
		if err != nil {
			return nil, err
		}
		paths = append(paths, p)
	}
	return paths, nil
}

func GetTaskPaths(tasks []Task, notesRoot string) ([]string, error) {
	return GetTaskPathsWithFinder(tasks, notesRoot, FindNoteByUUID)
}

func GetTaskPathWithFinder(task *Task, notesRoot string, finder NoteFinder) (string, error) {
	path, err := finder(task.UUID, notesRoot)
	if err != nil {
		return "", err
	} else if path != "" {
		return path, nil
	}
	return TaskNotePath(task, notesRoot)
}

func GetTaskPath(task *Task, notesRoot string) (string, error) {
	return GetTaskPathWithFinder(task, notesRoot, FindNoteByUUID)
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

