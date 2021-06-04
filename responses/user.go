package responses

import (
	"majoo-test/forms"
)

type UserDefault struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UserData struct {
	Message string     `json:"message"`
	Status  int        `json:"status"`
	Data    forms.User `json:"data"`
}

type ListUser struct {
	Total       int          `json:"total"`
	PerPage     int          `json:"per_page"`
	CurrentPage int          `json:"current_page"`
	LastPage    int          `json:"last_page"`
	NextPage    int          `json:"next_page"`
	PrevPage    int          `json:"prev_page"`
	Data        []forms.User `json:"data"`
	Status      int          `json:"status"`
}
