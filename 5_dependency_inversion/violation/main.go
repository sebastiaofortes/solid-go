package main

import (

	// some dependencies

	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/violation/controller"
	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/violation/repository"
	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/violation/service"
)

func main() {

	repository := repository.NewUserRepository()
	service := service.NewEmailService(repository)

	_ = controller.NewEmailController(service)

}
