package models

import "time"

type Employee struct {
	Name               string    `json:"name"`
	BirthDate          time.Time `json:"birthDate"`
	Salary             float64   `json:"salary"`
	ExperienceTime     int8      `json:"experienceTime"`
	physicalDisability bool      `json:"physicalDisability"`
}
