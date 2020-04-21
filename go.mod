module github.com/modularise/prometheus-promql

go 1.13

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/edsrzf/mmap-go v1.0.0
	github.com/go-kit/kit v0.10.0
	github.com/modularise/prometheus-tsdb v0.0.0-20200421120443-c343a51a2de4
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/common v0.9.1
)

replace k8s.io/klog => github.com/simonpasquier/klog-gokit v0.1.0
