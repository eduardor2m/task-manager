package request

type TaskDTO struct {
	Title       string `json:"title" example:"Task 1"`
	Description string `json:"description" example:"Task 1 description"`
	Status      bool   `json:"status" example:"false"`
	Category    string `json:"category" example:"Work"`
	Date        string `json:"date" example:"2021-01-01"`
}
