module github.com/micro/examples

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

exclude sourcegraph.com/sourcegraph/go-diff v0.5.1

replace github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d => github.com/nats-io/nats.go v1.8.1

require (
	github.com/99designs/gqlgen v0.9.1
	github.com/emicklei/go-restful v2.8.1+incompatible
	github.com/gin-contrib/sse v0.0.0-20190125020943-a7658810eb74 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.2
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711
	github.com/lucas-clemente/quic-go v0.11.2 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.7.1-0.20190711204633-5157241c88e0
	github.com/micro/go-plugins v1.1.2-0.20190710094942-bf407858372c
	github.com/micro/micro v1.7.1-0.20190711215914-2cddc2c877c5
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/pborman/uuid v1.2.0
	github.com/ugorji/go v1.1.5-pre // indirect
	github.com/vektah/gqlparser v1.1.2
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7
	google.golang.org/genproto v0.0.0-20190708153700-3bdd9d9f5532
	google.golang.org/grpc v1.22.0
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)
