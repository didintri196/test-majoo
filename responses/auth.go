package responses

import (
	"majoo-test/forms"
)

type AuthOK struct {
	Message string     `json:"message"`
	Status  int        `json:"status"`
	Expired int64      `json:"expired"`
	Token   string     `json:"token"`
	User    forms.User `json:"user"`
}

type AuthError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
