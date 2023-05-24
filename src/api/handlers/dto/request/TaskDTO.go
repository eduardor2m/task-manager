package request

type TaskDTO struct {
	Title       string `json:"title" example:"Task 1"`
	Description string `json:"description" example:"Task 1 description"`
	Completed   bool   `json:"completed" example:"false"`
}
