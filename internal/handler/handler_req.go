package handler

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FetchTaskRequest struct {
	Title       string `query:"title"`
	Description string `query:"description"`
	Page        int64  `query:"page"`
	PerPage     int64  `query:"perPage"`
}

type UpateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ID          string `json:"_id"`
	Status      string `json:"status"`
}
