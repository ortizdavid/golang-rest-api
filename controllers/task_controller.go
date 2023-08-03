package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ortizdavid/golang-rest-api/entities"
	"github.com/ortizdavid/golang-rest-api/models"
)

type TaskController struct {
}

func (task TaskController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/tasks", task.getAllTasks).Methods("GET")
	router.HandleFunc("/api/tasks", task.createTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", task.getTask).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", task.updateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", task.deleteTask).Methods("DELETE")
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

func (TaskController) createTask(w http.ResponseWriter, r  *http.Request) {
	var task entities.Task
	var taskModel models.TaskModel
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := taskModel.Create(task).Error; err != nil {
		WriteError(w, http.StatusInternalServerError, "Error while creating data")
		return
	}
	WriteJSON(w, "Task Created!", http.StatusCreated, task, 1)
}


func (TaskController) getTask(w http.ResponseWriter, r  *http.Request) {
	var taskModel models.TaskModel
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if !taskModel.Exists(taskId) {
		WriteError(w, http.StatusNotFound, "Task not found")
		return
	}
	task := models.TaskModel{}.FindById(taskId)
	WriteJSON(w, "Tasks found", http.StatusOK, task, 1)
}

func (TaskController) updateTask(w http.ResponseWriter, r  *http.Request) {
	var taskModel models.TaskModel
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if !taskModel.Exists(taskId) {
		WriteError(w, http.StatusNotFound, "Task not found")
		return
	}
	task := models.TaskModel{}.FindById(taskId)
	taskModel.Update(task)
	WriteJSON(w, "Task Updated!", http.StatusOK, task, 1)
}

func (TaskController) deleteTask(w http.ResponseWriter, r  *http.Request) {
	WriteJSON(w, "Task Deleted!", http.StatusOK, nil, 1)
}