module github.com/zorneth/osg-proxy

go 1.27.0

require github.com/zorneth/osg-core v0.1.0-alpha.1

require (
	github.com/kr/text v0.2.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/zorneth/osg-core => ../osg-core
