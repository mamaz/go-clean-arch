package todo

import "go-clean-arch/domain/user"

type Todo struct {
	ID   int64      `json:"id"`
	Text string     `json:"text"`
	Done bool       `json:"done"`
	User *user.User `json:"user"`
}
