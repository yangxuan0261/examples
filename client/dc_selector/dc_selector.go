package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	example "github.com/micro/examples/server/proto/example"
)

// Built in random hashed node selector
type dcSelector struct {
	opts selector.Options
}

var (
	datacenter = "local"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func (n *dcSelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&n.opts)
	}
	return nil
}

func (n *dcSelector) Options() selector.Options {
	return n.opts
}

var i int

func (n *dcSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	fmt.Printf("--- Select, service:%s\n", service)
	services, err := n.opts.Registry.GetService(service)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	var nodes []*registry.Node

	// Filter the nodes for datacenter
	for _, service := range services {
		for _, node := range service.Nodes { // 每次给到的 Nodes 是乱序的, 可以优化到排序后本地缓存, 长度不一致时重新缓存最新的即可
			if node.Metadata["datacenter"] == datacenter {
				// fmt.Println("--- node.Metadata:", node.Metadata)
				fmt.Printf("--- node:%+v\n", node)
				// fmt.Printf("--- node, Address:%s, Id:%s\n", node.Address, node.Id)
				nodes = append(nodes, node)
			}
		}
	}

	if len(nodes) == 0 {
		return nil, selector.ErrNotFound
	}

	var mtx sync.Mutex

	fmt.Printf("--- Select, len(nodes):%d\n", len(nodes))

	return func() (*registry.Node, error) {
		mtx.Lock()
		defer mtx.Unlock()
		i++
		fmt.Printf("--- Select, call node:%d\n", i%len(nodes))
		return nodes[i%len(nodes)], nil
	}, nil
}

func (n *dcSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (n *dcSelector) Reset(service string) {
	return
}

func (n *dcSelector) Close() error {
	return nil
}

func (n *dcSelector) String() string {
	return "dc"
}

// Return a new first node selector
func DCSelector(opts ...selector.Option) selector.Selector {
	var sopts selector.Options
	for _, opt := range opts {
		opt(&sopts)
	}
	if sopts.Registry == nil {
		fmt.Println("--- sopts.Registry == nil")
		registerDrive := etcdv3.NewRegistry(func(op *registry.Options) {
			op.Addrs = []string{
				"http://127.0.0.1:2379",
			}
		})
		_ = registerDrive
		sopts.Registry = registerDrive // 指定 etcdv3 为服务发现
		// sopts.Registry = registry.DefaultRegistry
	}
	return &dcSelector{sopts}
}

func call(i int) {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest("go.micro.srv.example", "Example.Call", &example.Request{
		Name: "John",
	})

	rsp := &example.Response{}

	// Call service
	if err := client.Call(context.Background(), req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Call:", i, "rsp:", rsp.Msg)
}

func main() {
	cmd.Init()

	client.DefaultClient = client.NewClient(
		client.Selector(DCSelector()),
	)

	fmt.Println("\n--- Call example ---")
	for i := 0; i < 10; i++ {
		call(i)
		time.Sleep(time.Second * 1)
	}
}
