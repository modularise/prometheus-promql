module github.com/modularise/prometheus-promql

go 1.13

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/edsrzf/mmap-go v1.0.0
	github.com/go-kit/kit v0.10.0
	github.com/modularise/prometheus-tsdb v0.0.0-20200625120835-909378dad7da
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang v1.6.0
	github.com/prometheus/common v0.10.0
	github.com/uber/jaeger-client-go v2.23.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/atomic v1.6.0 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200603131246-cc40288be839 // indirect
)

replace k8s.io/klog => github.com/simonpasquier/klog-gokit v0.1.0
