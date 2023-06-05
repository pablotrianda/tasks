package config

import (
	"fmt"
	"os"

	"github.com/pablotrianda/config"
)

type Config struct{
	TaskDir string
	Editor string
}

func LoadConfig() Config{
	// Load the configuration file
	homeDir, err := os.UserHomeDir()
	cfg, err := config.Load(homeDir + "/.config/tasks/conf.yaml")
	if err != nil {
		panic(fmt.Errorf("Config file not found, put on ~/.config/tasks/conf.yaml"))
	}
	tasksDir, err := cfg.Get("tasks")
	if err != nil {
		panic(fmt.Errorf("tasks config is missing"))
	}
	editor, err := cfg.Get("editor")
	if err != nil {
		panic(fmt.Errorf("editor config is missing"))
	}

	return Config{
		TaskDir: tasksDir,
		Editor: editor,
	}

}
