package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ortizdavid/golang-rest-api/entities"
	"github.com/ortizdavid/golang-rest-api/models"
)

type TaskController struct {
}

func (task TaskController) RegisterRoutes() {
	http.HandleFunc("/api/tasks", task.getAllTasks)
	http.HandleFunc("/api/tasks/create", task.createTask)
	http.HandleFunc("/api/tasks/get", task.getTask)
	http.HandleFunc("/api/tasks/update", task.updateTask)
	http.HandleFunc("/api/tasks/delete", task.deleteTask)
}

func (TaskController) getAllTasks(w http.ResponseWriter, r  *http.Request) {
	tasks := models.TaskModel{}.FindAll()
	count := len(tasks)
	if count == 0 {
		WriteJSON(w, "Tasks not found", http.StatusNotFound, tasks, count)
	} else {
		WriteJSON(w, "Tasks found successfully", http.StatusOK, tasks, count)
	}
}

func (TaskController) getTask(w http.ResponseWriter, r  *http.Request) {
	var task entities.Task
	WriteJSON(w, "Tasks not found", http.StatusOK, task, nil)
}

func (TaskController) createTask(w http.ResponseWriter, r  *http.Request) {
	var task entities.Task
	var taskModel models.TaskModel
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	if err := taskModel.Create(task).Error; err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while creating data")
		return
	}
	WriteJSON(w, "Task Created!", http.StatusOK, task, 1)
}

func (TaskController) updateTask(w http.ResponseWriter, r  *http.Request) {
	WriteJSON(w, "Task Updated!", http.StatusOK, nil, 1)
}

func (TaskController) deleteTask(w http.ResponseWriter, r  *http.Request) {
	WriteJSON(w, "Task Deleted!", http.StatusOK, nil, nil)
}