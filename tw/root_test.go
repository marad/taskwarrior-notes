package tw

import "testing"

func TestGetTaskPathWithFinder(t *testing.T) {
	root := "/root"
	genTask := Task{Entry: "20250101T000000Z", Description: "Test desc", UUID: "abc"}
	// Finder returns existing path
	existing := func(uuid, notesRoot string) (string, error) { return notesRoot + "/existing.md", nil }
	p, err := GetTaskPathWithFinder(&genTask, root, existing)
	if err != nil {
		 t.Fatalf("unexpected error: %v", err)
	}
	if p != root+"/existing.md" {
		 t.Fatalf("expected existing path, got %s", p)
	}
	// Finder returns empty -> fallback
	emptyFinder := func(uuid, notesRoot string) (string, error) { return "", nil }
	p2, err := GetTaskPathWithFinder(&genTask, root, emptyFinder)
	if err != nil {
		 t.Fatalf("unexpected error: %v", err)
	}
	if p2 == root+"/existing.md" {
		 t.Fatalf("fallback path should differ from existing")
	}
	// Finder error propagated
	errFinder := func(uuid, notesRoot string) (string, error) { return "", assertErr{} }
	_, err = GetTaskPathWithFinder(&genTask, root, errFinder)
	if err == nil {
		 t.Fatalf("expected error from finder")
	}
}

type assertErr struct{}

func (assertErr) Error() string { return "finder error" }

func TestGetTaskPathsWithFinder(t *testing.T) {
	root := "/root"
	t1 := Task{Entry: "20250101T000000Z", Description: "A", UUID: "u1"}
	t2 := Task{Entry: "20250102T000000Z", Description: "B", UUID: "u2"}
	finder := func(uuid, notesRoot string) (string, error) {
		if uuid == "u1" {
			return notesRoot + "/existing1.md", nil
		}
		return "", nil
	}
	paths, err := GetTaskPathsWithFinder([]Task{t1, t2}, root, finder)
	if err != nil {
		 t.Fatalf("unexpected error: %v", err)
	}
	if len(paths) != 2 {
		 t.Fatalf("expected 2 paths, got %d", len(paths))
	}
	if paths[0] != root+"/existing1.md" {
		 t.Fatalf("first path should be existing: %s", paths[0])
	}
	if paths[1] == root+"/existing1.md" {
		 t.Fatalf("second path should be generated, got existing")
	}
	// Empty slice
	paths, err = GetTaskPathsWithFinder([]Task{}, root, finder)
	if err != nil || len(paths) != 0 {
		 t.Fatalf("expected empty result for empty input")
	}
}
