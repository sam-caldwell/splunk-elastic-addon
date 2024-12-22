.PHONY: build
build:
	@mkdir -p build/splunk-elastic-addon
	@rsync -av src/addon/* build/splunk-elastic-addon/
	@go build -o build/splunk-elastic-addon/default/bin/query-elastic src/cmd/query-elastic/main.go
	@go build -o build/splunk-elastic-addon/default/bin/setup src/cmd/setup/main.go
