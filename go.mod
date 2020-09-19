module github.com/layer5io/meshery-kuma

go 1.13

replace (
	github.com/kudobuilder/kuttl => github.com/layer5io/kuttl v0.4.1-0.20200806180306-b7e46afd657f

	github.com/mgfeller/common-adapter-library => ../common-adapter-library
)

require (
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/layer5io/gokit v0.0.0-20200903172320-9df3e7bc469c
	github.com/layer5io/learn-layer5/smi-conformance v0.0.0-20200825194222-14309c02bff2 // indirect
	github.com/mgfeller/common-adapter-library v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.7.0
	go.opentelemetry.io/otel v0.10.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.10.0
	go.opentelemetry.io/otel/sdk v0.10.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	google.golang.org/genproto v0.0.0-20200731012542-8145dea6a485 // indirect
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/client-go v0.18.8
	rsc.io/letsencrypt v0.0.3 // indirect
)
