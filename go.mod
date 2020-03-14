module github.com/modularise/prometheus-promql

go 1.13

require (
	github.com/edsrzf/mmap-go v1.0.0
	github.com/go-kit/kit v0.10.0
	github.com/modularise/prometheus-tsdb v0.0.0-20200314120653-b6cc223e4631
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.0
	github.com/prometheus/common v0.9.1
)

replace k8s.io/klog => github.com/simonpasquier/klog-gokit v0.1.0
