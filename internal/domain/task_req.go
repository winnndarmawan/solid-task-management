package task

import (
	"errors"
	"fmt"
)

type CreateReq struct {
	Title       string
	Description string
}

func (s *CreateReq) Validate() error {
	fmt.Printf("Validating CreateReq: %+v\n", s)
	if s.Title == "" {
		return errors.New("title is empty")
	}
	if s.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

type FetchTasksReq struct {
	Title       string
	Description string
	Page        int64
	PerPage     int64
}
