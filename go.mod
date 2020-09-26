module github.com/layer5io/meshery-kuma

go 1.13

replace (
	github.com/kudobuilder/kuttl => github.com/layer5io/kuttl v0.4.1-0.20200806180306-b7e46afd657f

	github.com/mgfeller/common-adapter-library => ../common-adapter-library
)

require (
	github.com/layer5io/gokit v0.1.12
	github.com/mgfeller/common-adapter-library v0.0.0-00010101000000-000000000000
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/viper v1.7.0
)
