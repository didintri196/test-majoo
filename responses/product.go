package responses

import (
	"majoo-test/forms"
)

type ProductDefault struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type ProductData struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Data    forms.Product `json:"data"`
}

type ListProduct struct {
	Total       int             `json:"total"`
	PerPage     int             `json:"per_page"`
	CurrentPage int             `json:"current_page"`
	LastPage    int             `json:"last_page"`
	NextPage    int             `json:"next_page"`
	PrevPage    int             `json:"prev_page"`
	Data        []forms.Product `json:"data"`
	Status      int             `json:"status"`
}
