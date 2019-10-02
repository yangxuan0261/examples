module github.com/micro/examples

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

exclude sourcegraph.com/sourcegraph/go-diff v0.5.1

replace github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d => github.com/nats-io/nats.go v1.8.1

replace k8s.io/api => k8s.io/api v0.0.0-20190708174958-539a33f6e817

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d

replace k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190708180123-608cd7da68f7

replace k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible

replace k8s.io/component-base => k8s.io/component-base v0.0.0-20190708175518-244289f83105

require (
	github.com/99designs/gqlgen v0.9.1
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/emicklei/go-restful v2.8.1+incompatible
	github.com/gin-contrib/sse v0.0.0-20190125020943-a7658810eb74 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/websocket v1.4.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.5
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.11.1-0.20191001180929-e8a53610f17b
	github.com/micro/go-plugins v1.3.0
	github.com/micro/micro v1.11.1-0.20191001181547-221ee138ab7c
	github.com/pborman/uuid v1.2.0
	github.com/ugorji/go v1.1.5-pre // indirect
	github.com/vektah/gqlparser v1.1.2
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	google.golang.org/genproto v0.0.0-20190801165951-fa694d86fc64
	google.golang.org/grpc v1.22.1
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)

go 1.13
