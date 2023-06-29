package services

import (
	"github.com/eduardor2m/task-manager/src/core/domain/task"
	"github.com/eduardor2m/task-manager/src/core/interfaces/repository"
	"github.com/google/uuid"
)

type TaskServices struct {
	taskRepository repository.TaskLoader
}

func (instance TaskServices) CreateTask(task task.Task) (*uuid.UUID, error) {
	return instance.taskRepository.CreateTask(task)
}

func (instance TaskServices) GetTask(id uuid.UUID) (*task.Task, error) {
	return instance.taskRepository.GetTask(id)
}

func (instance TaskServices) GetTasks() ([]*task.Task, error) {
	return instance.taskRepository.GetTasks()
}

func (instance TaskServices) UpdateTask(task task.Task) (*task.Task, error) {
	return instance.taskRepository.UpdateTask(task)
}

func (instance TaskServices) UpdateTaskStatus(id uuid.UUID) (*task.Task, error) {
	return instance.taskRepository.UpdateTaskStatus(id)
}

func (instance TaskServices) DeleteTask(id uuid.UUID) error {
	return instance.taskRepository.DeleteTask(id)
}

func (instance TaskServices) DeleteTasks() error {
	return instance.taskRepository.DeleteTasks()
}

func NewTaskServices(taskRepository repository.TaskLoader) TaskServices {
	return TaskServices{taskRepository: taskRepository}
}
