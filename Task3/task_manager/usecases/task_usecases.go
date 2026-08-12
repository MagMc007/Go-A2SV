package usecases

/*
	business logic, normal services, validation
*/
import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"task_manager/domain"
)

type TaskUsecase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUsecase(taskRepository domain.TaskRepository) domain.TaskUsecase {
	return &TaskUsecase{
		taskRepository: taskRepository,
	}
}

func (tu *TaskUsecase) AddTask(task domain.Task) (domain.Task, error) {
	return tu.taskRepository.AddTask(task)
}

func (tu *TaskUsecase) GetAllTasks() ([]domain.Task, error) {
	return tu.taskRepository.GetAllTasks()
}

func (tu *TaskUsecase) GetTaskDetails(id primitive.ObjectID) (domain.Task, error) {
	return tu.taskRepository.GetTaskDetails(id)
}

func (tu *TaskUsecase) UpdateTask(id primitive.ObjectID, task domain.Task) (domain.Task, error) {
	return tu.taskRepository.UpdateTask(id, task)
}

func (tu *TaskUsecase) DeleteTask(id primitive.ObjectID) error {
	return tu.taskRepository.DeleteTask(id)
}