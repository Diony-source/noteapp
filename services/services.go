package services

import (
	"encoding/json"
	"os"
	"noteapp/entities"
)

const filePath = "note.json"

func LoadNotes() ([]entities.Note, error) {
	var notes []entities.Note

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return notes, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return notes, err
	}

	err = json.Unmarshal(data, &notes)
	return notes, err
}

func SaveNotes(notes []entities.Note) error {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}