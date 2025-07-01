package tw

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

type Task struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Entry       string  `json:"entry"`
	Modified    string  `json:"modified"`
	Status      string  `json:"status"`
	UUID        string  `json:"uuid"`
	Urgency     float64 `json:"urgency"`
}

func TaskFromString(s string) (*Task, error) {
	var t Task
	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

const TW_TIME_FORMAT = "20060102T150405Z"

func (t *Task) EntryTime() (time.Time, error) {
	return time.Parse(TW_TIME_FORMAT, t.Entry)
}

func (t *Task) ModifiedTime() (time.Time, error) {
	return time.Parse(TW_TIME_FORMAT, t.Modified)
}

func TaskNotePath(t *Task, taskNoteRoot string) (string, error) {
	entryTime, err := t.EntryTime()
	if err != nil {
		return "", err
	}
	year := entryTime.Format("2006")
	month := entryTime.Format("01")
	day := entryTime.Format("02")

	// Sanitize description for filename (very basic, replace spaces with underscores)
	desc := t.Description
	forbidden := []rune{'/', '\\', ':', '*', '?', '"', '<', '>', '|'}
	for _, f := range forbidden {
		desc = strings.ReplaceAll(desc, string(f), "_")
	}
	filename := fmt.Sprintf("%s.md", desc)
	monthDay := fmt.Sprintf("%s-%s", month, day)
	path := filepath.Join(taskNoteRoot, year, monthDay, filename)
	return path, nil
}
