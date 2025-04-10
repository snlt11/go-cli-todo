package todo

import (
	"cli-todo/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const taskFile = "tasks.json"

func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

func LoadTasks() ([]models.Task, error) {
	var tasks []models.Task

	if _, err := os.Stat(taskFile); errors.Is(err, os.ErrNotExist) {
		return tasks, nil 
	}

	data, err := os.ReadFile(taskFile)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return tasks, nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("could not parse tasks.json: %w", err)
	}
	return tasks, nil
}
