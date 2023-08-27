package dependencyinversion

import (
	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/no-violation/controller"
	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/no-violation/repository"
	"github.com/sebastiaofortes/solid-go/5_dependency_inversion/no-violation/service"
)

func main() {
	r := repository.NewEmailRepository()

	s := service.NewEmailService(r)

	_ = controller.NewEmailController(s)

}
