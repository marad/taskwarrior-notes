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
	"os"
	"os/exec"
	"taskwarrior-notes/config"
	"taskwarrior-notes/tw"
	"taskwarrior-notes/util"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Opens the editor to modify a task note",
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
		if len(tasks) == 0 {
			fmt.Println("No tasks found to edit.")
			return
		}

		if len(tasks) > 1 {
			fmt.Printf("Multiple tasks found (%d). Do you want to edit all of them? (y/n): ", len(tasks))
			var response string
			fmt.Scanln(&response)
			if response != "y" {
				return
			}
		}

		for _, task := range tasks {
			err := openEditor(task)
			if err != nil {
				fmt.Println("Error opening editor:", err)
			}
		}
	},
}

func openEditor(task tw.Task) error {
	path, err := tw.GetTaskPath(&task, config.ReadNotesRoot())
	if err != nil {
		return fmt.Errorf("error getting task path: %w", err)
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	fmt.Printf("Opening editor '%s' for task '%s'...\n", editor, task.Description)
	cmd := exec.Command(editor, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
