package tw

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestTaskFromString(t *testing.T) {
	validJSON := `{"id":5,"description":"Do thing","entry":"20250101T120000Z","modified":"20250101T120500Z","end":"","status":"pending","uuid":"11111111-2222-3333-4444-555555555555","urgency":3.7,"extra":"ignored"}`
	truncated := `{"id":5`
	wrongType := `{"id":"notNumber"}`
	empty := ``
	whitespace := `   `

	tests := []struct {
		name      string
		input     string
		expectErr bool
		check     func(*Task) error
	}{
		{
			name:  "valid",
			input: validJSON,
			check: func(task *Task) error {
				if task.ID != 5 || task.Description != "Do thing" || task.Entry != "20250101T120000Z" || task.Modified != "20250101T120500Z" || task.End != "" || task.Status != "pending" || task.UUID != "11111111-2222-3333-4444-555555555555" || task.Urgency != 3.7 {
					return &reflect.ValueError{}
				}
				return nil
			},
		},
		{name: "truncated", input: truncated, expectErr: true},
		{name: "wrong_type", input: wrongType, expectErr: true},
		{name: "empty", input: empty, expectErr: true},
		{name: "whitespace", input: whitespace, expectErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := TaskFromString(tc.input)
			if tc.expectErr && err == nil {
				 t.Fatalf("expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				 t.Fatalf("unexpected error: %v", err)
			}
			if !tc.expectErr && tc.check != nil {
				if cerr := tc.check(res); cerr != nil {
					 t.Fatalf("field check failed")
				}
			}
			if tc.expectErr && res != nil {
				 t.Fatalf("expected nil task on error")
			}
		})
	}
}

func TestTaskTimes(t *testing.T) {
	entry := "20250101T120000Z"
	modified := "20250101T121000Z"
	end := "20250101T130000Z"
	bad := "2025BAD"

	baseTask := Task{Entry: entry, Modified: modified, End: end}

	et, err := baseTask.EntryTime()
	if err != nil {
		 t.Fatalf("EntryTime error: %v", err)
	}
	mt, err := baseTask.ModifiedTime()
	if err != nil {
		 t.Fatalf("ModifiedTime error: %v", err)
	}
	endt, err := baseTask.EndTime()
	if err != nil {
		 t.Fatalf("EndTime error: %v", err)
	}
	if !et.Equal(time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)) {
		 t.Fatalf("EntryTime mismatch: %v", et)
	}
	if !mt.Equal(time.Date(2025, 1, 1, 12, 10, 0, 0, time.UTC)) {
		 t.Fatalf("ModifiedTime mismatch: %v", mt)
	}
	if !endt.Equal(time.Date(2025, 1, 1, 13, 0, 0, 0, time.UTC)) {
		 t.Fatalf("EndTime mismatch: %v", endt)
	}

	emptyEndTask := Task{End: ""}
	z, err := emptyEndTask.EndTime()
	if err != nil || !z.IsZero() {
		 t.Fatalf("expected zero time and nil error for empty end")
	}

	badTask := Task{Entry: bad}
	_, err = badTask.EntryTime()
	if err == nil {
		 t.Fatalf("expected error for bad entry time format")
	}
}

func TestTaskNotePath(t *testing.T) {
	root := "/notesroot"

	tests := []struct {
		name        string
		entry       string
		desc        string
		expectErr   bool
		filenameChk func(string) bool
	}{
		{
			name:  "simple",
			entry: "20250101T120000Z",
			desc:  "Do task",
			filenameChk: func(fn string) bool { return fn == "Do task.md" },
		},
		{
			name:  "forbidden_chars",
			entry: "20250102T000000Z",
			desc:  "a/b:c*d?e<f>g|h",
			filenameChk: func(fn string) bool { return fn == "a_b_c_d_e_f_g_h.md" },
		},
		{
			name:  "only_forbidden",
			entry: "20250103T000000Z",
			desc:  "/:*?<>|",
			filenameChk: func(fn string) bool { return fn == "_______.md" }, // 7 forbidden chars -> 7 underscores
		},
		{
			name:      "malformed_entry",
			entry:     "BAD",
			desc:      "X",
			expectErr: true,
		},
		{
			name:  "leap_day",
			entry: "20240229T010203Z",
			desc:  "Leap",
			filenameChk: func(fn string) bool { return fn == "Leap.md" },
		},
	}

	for _, tc := range tests {
		 t.Run(tc.name, func(t *testing.T) {
			task := Task{Entry: tc.entry, Description: tc.desc}
			p, err := TaskNotePath(&task, root)
			if tc.expectErr && err == nil {
				 t.Fatalf("expected error, got nil")
			}
			if !tc.expectErr && err != nil {
				 t.Fatalf("unexpected error: %v", err)
			}
			if tc.expectErr {
				return
			}
			parts := strings.Split(p, "/")
			filename := parts[len(parts)-1]
			if tc.filenameChk != nil && !tc.filenameChk(filename) {
				 t.Fatalf("filename check failed: %s", filename)
			}
		})
	}
}
