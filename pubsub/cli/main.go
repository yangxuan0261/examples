package main

import (
	"fmt"
	"time"

	"context"

	proto "github.com/micro/examples/pubsub/srv/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/grpc"
	"github.com/pborman/uuid"
	"github.com/micro/go-micro/metadata"
)

// send events using the publisher
func sendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)

	for _ = range t.C {
		// create new event
		ev := &proto.Event{
			Id:        uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("Messaging you all day on %s", topic),
		}

		log.Logf("publishing %+v\n", ev)

		// 加上一个 metadata
		md := map[string]string{
			"aaa": "111",
			"bbb": "222",
		}
		ctx := metadata.NewContext(context.Background(), md)

		// publish an event
		if err := p.Publish(ctx, ev); err != nil {
			log.Logf("error publishing: %v", err)
		}
	}
}

func main() {
	grpcBroker := grpc.NewBroker()

	// create a service
	service := micro.NewService(
		micro.Name("go.micro.cli.pubsub"),
		micro.Broker(grpcBroker), // 使用 grpcBroker
	)
	// parse command line
	service.Init()

	// create publisher
	pub1 := micro.NewPublisher("example.topic.pubsub.1", service.Client())
	pub2 := micro.NewPublisher("example.topic.pubsub.2", service.Client())

	// pub to topic 1
	go sendEv("example.topic.pubsub.1", pub1)
	// pub to topic 2
	go sendEv("example.topic.pubsub.2", pub2)

	// block forever
	select {}
}
