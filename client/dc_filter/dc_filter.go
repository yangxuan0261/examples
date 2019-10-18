package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	example "github.com/micro/examples/server/proto/example"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// A Wrapper that creates a Datacenter Selector Option
type dcWrapper struct {
	client.Client
}

func (dc *dcWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("--- md:%+v\n", md)

	filter := func(services []*registry.Service) []*registry.Service {
		for _, service := range services {
			var nodes []*registry.Node
			for _, node := range service.Nodes {
				fmt.Printf("--- node:%+v\n", node)
				if node.Metadata["datacenter"] == md["datacenter"] {
					nodes = append(nodes, node)
				}
			}
			service.Nodes = nodes
		}
		return services
	}

	callOptions := append(opts, client.WithSelectOption(
		selector.WithFilter(filter),
	))

	fmt.Printf("[DC Wrapper] filtering for datacenter %s\n", md["datacenter"])
	return dc.Client.Call(ctx, req, rsp, callOptions...)
}

func NewDCWrapper(c client.Client) client.Client {
	return &dcWrapper{c}
}

func call(i int) {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest("go.micro.srv.example", "Example.Call", &example.Request{
		Name: "John",
	})

	// create context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"datacenter": "local",
		"aaa":        "111",
	})

	rsp := &example.Response{}

	// Call service
	if err := client.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Call:", i, "rsp:", rsp.Msg)
}

func main() {
	cmd.Init()

	registerDrive := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})
	_ = registerDrive

	client.DefaultClient = client.NewClient(
		client.Wrap(NewDCWrapper),
		client.Registry(registerDrive),
	)

	fmt.Println("\n--- Call example ---")
	for i := 0; i < 10; i++ {
		call(i)
	}
}
