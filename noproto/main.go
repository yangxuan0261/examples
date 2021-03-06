package main

import (
	"context"

	"github.com/micro/go-micro"
)

type Greeter struct{} // 必须 首字母大写, 对外部包可以访问

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	*msg = "Hello " + *name
	return nil
}

func main() {
	// create new service
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// initialise command line
	service.Init()

	// set the handler
	micro.RegisterHandler(service.Server(), new(Greeter))

	// run service
	service.Run()
}
