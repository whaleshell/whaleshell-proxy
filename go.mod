module github.com/zorneth/osg-proxy

go 1.27.0

require (
	github.com/glaciforge/slogx v0.0.0
	github.com/zorneth/osg-core v0.1.0-alpha.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lkmavi/saferefl v0.4.0 // indirect
	go.opentelemetry.io/otel v1.46.0 // indirect
	go.opentelemetry.io/otel/sdk v1.46.0 // indirect
	go.opentelemetry.io/otel/trace v1.46.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
)

replace (
	github.com/glaciforge/slogx => ../slogx
	github.com/zorneth/osg-core => ../osg-core
)
