package endpoints

import "github.com/leonardogregoriocs/task_management_system/internal/user"

type Handler struct {
	UserService user.UserService
}
