package entities

import "time"

type Task struct {
	TaskId      int64  `json:"taskId"`
	TaskName    string `json:"taskName"`
	StartDate   time.Time `json:"startDate"`
	EndDate   	time.Time `json:"endDate"`
	Description string `json:"description"`
	Status      string `json:"status"`
}