package main

import (
	"os"
	"os/exec"

	t "github.com/pablotrianda/tasks/src/pkg/tasks"
	c "github.com/pablotrianda/tasks/src/pkg/config"
)


func main(){
	config := c.LoadConfig()
	task := t.NewTask(config.TaskDir)
	runEditor(config.Editor, task)
}


// Run the editor with the name of the new file
func runEditor(editor string, task t.Task)  {
	os.Chdir(task.Directory)
	cmd := exec.Command(editor, task.File)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd.Wait()
}

