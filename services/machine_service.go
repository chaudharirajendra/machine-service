package services

import (
	"encoding/json"
	"errors"
	"machine-service/models"
	"os"
)

// GetMachines fetches all machines from the JSON file
func GetMachines() ([]models.Machine, error) {
	file, err := os.Open("services/machines.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var machines []models.Machine
	if err := json.NewDecoder(file).Decode(&machines); err != nil {
		return nil, err
	}
	return machines, nil
}

// GetMachineByID fetches a single machine by its ID
func GetMachineByID(machineID string) (*models.Machine, error) {
	machines, err := GetMachines()
	if err != nil {
		return nil, err
	}

	for _, machine := range machines {
		if machine.ID == machineID {
			return &machine, nil
		}
	}

	return nil, errors.New("not found")
}
