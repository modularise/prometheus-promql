module github.com/modularise/prometheus-promql

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/edsrzf/mmap-go v1.0.0
	github.com/go-kit/kit v0.10.0
	github.com/modularise/prometheus-tsdb v0.0.0-20200821121532-af1a82ea8acd
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang v1.7.1
	github.com/prometheus/common v0.13.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/goleak v1.0.0
)

replace k8s.io/klog => github.com/simonpasquier/klog-gokit v0.1.0

exclude (
	// Exclude grpc v1.30.0 because of breaking changes. See #7621.
	google.golang.org/grpc v1.30.0

	// Exclude pre-go-mod kubernetes tags, as they are older
	// than v0.x releases but are picked when we update the dependencies.
	k8s.io/client-go v1.4.0
	k8s.io/client-go v1.4.0+incompatible
	k8s.io/client-go v1.5.0
	k8s.io/client-go v1.5.0+incompatible
	k8s.io/client-go v1.5.1
	k8s.io/client-go v1.5.1+incompatible
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/client-go v2.0.0+incompatible
	k8s.io/client-go v2.0.0-alpha.1+incompatible
	k8s.io/client-go v3.0.0+incompatible
	k8s.io/client-go v3.0.0-beta.0+incompatible
	k8s.io/client-go v4.0.0+incompatible
	k8s.io/client-go v4.0.0-beta.0+incompatible
	k8s.io/client-go v5.0.0+incompatible
	k8s.io/client-go v5.0.1+incompatible
	k8s.io/client-go v6.0.0+incompatible
	k8s.io/client-go v7.0.0+incompatible
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/client-go v9.0.0+incompatible
	k8s.io/client-go v9.0.0-invalid+incompatible
)
