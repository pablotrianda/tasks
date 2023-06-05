package tasks

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	File string
	Directory string
}

func NewTask(principalDirectory string) Task {
	os.Chdir(principalDirectory)
	newDirectory := createDirectoryName(principalDirectory)
	createDirectory(newDirectory)
	newDayInput := createTodayInput()
	return Task{
		File: newDayInput,
		Directory: newDirectory,
	}
}

func formatMonth(month int) string {
	stringMonth := strconv.Itoa(month)

	if month < 10 {
		return "0" + stringMonth
	}

	return stringMonth
}

// Return the current date on format: yyyymmdd
func getDateInfo(completeWithDay bool) string {
	currentTime := time.Now()

	year := strconv.Itoa(currentTime.Year())
	month := formatMonth(int(currentTime.Month()))
	day := strconv.Itoa(currentTime.Day())

	// newDirectory name: tasks_yyyymmdd
	if completeWithDay{
		return year + month + day
	}else{
		return year + month
	}
}

// Create TIL directory if is necessary
func createDirectoryName(baseDir string) string {
	// newDirectory name: tasks_yyyymmdd
	dateInfo := getDateInfo(false)
	newDirectory := fmt.Sprintf("tasks_%s", dateInfo)

	return newDirectory
}

func createDirectory(newDirectory string) {
	// If newDirectory not exist, create a new createDirectory
	if _, err := os.Stat(newDirectory); os.IsNotExist(err) {
		err := os.Mkdir(newDirectory, 0750)
		if err != nil {
			panic(fmt.Errorf("Error to create a new directory"))
		}
	}
}

func createTodayInput() string {
	return getDateInfo(true)+"_notes.md"
}
