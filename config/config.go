package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ReadNotesRoot() string {
	home, err := os.UserHomeDir()
	if err != nil {
		os.Stderr.WriteString("Error: cannot determine home directory.\n")
		os.Exit(1)
	}
	file, err := os.Open(filepath.Join(home, ".taskrc"))
	if err != nil {
		os.Stderr.WriteString("Error: cannot open ~/.taskrc.\n")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "notes.root=") {
			return strings.TrimSpace(strings.TrimPrefix(line, "notes.root="))
		}
	}
	if err := scanner.Err(); err != nil {
		os.Stderr.WriteString("Error: reading ~/.taskrc failed.\n")
		os.Exit(1)
	}
	os.Stderr.WriteString("Error: notes.root not found in ~/.taskrc.\n")
	os.Exit(1)
	return "" // unreachable
}
