/*
Copyright © 2025 Marcin Radoszewski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os/exec"
	"taskwarrior-notes/config"
	"taskwarrior-notes/tw"
	"taskwarrior-notes/util"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronizes a task metadata with the note",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := util.GetTaskFromFlagOrSearchTw(cmd, args)
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}

		for _, task := range tasks {
			syncTaskWithNote(&task, config.ReadNotesRoot())
		}
		fmt.Println("Synchronization complete.")
	},
}

func syncTaskWithNote(task *tw.Task, notesRoot string) error {
	path, err := tw.GetTaskPath(task, notesRoot)
	if err != nil {
		return fmt.Errorf("error getting task path: %w", err)
	}

	// Ustaw status w frontmatter
	args := []string{"set", "status=" + task.Status}

	// Ustaw lub usuń doneDate
	if task.End != "" {
		// Formatowanie daty na YYYY-MM-DD
		endTime, err := task.EndTime() // zakładamy, że task.End jest w formacie ISO8601
		if err != nil {
			return fmt.Errorf("error parsing end time: %w", err)
		}
		doneDate := endTime.Format("2006-01-02")
		args = append(args, "doneDate="+doneDate)
		args = append(args, path)
		cmd := exec.Command("frontmatter", args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("frontmatter set error: %v, output: %s", err, string(output))
		}
	} else {
		args = append(args, path)
		cmd := exec.Command("frontmatter", args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("frontmatter set error: %v, output: %s", err, string(output))
		}

		cmd = exec.Command("frontmatter", "delete", "doneDate", path)
		output, err = cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("frontmatter delete error: %v, output: %s", err, string(output))
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
