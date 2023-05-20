package handler

import "std/pkg/service"

type handler struct {
	service.UserService
}

func NewHandler(userService service.UserService) *handler {
	return &handler{
		UserService: userService,
	}
}
