test: unit-test integration-test
	@echo 'test:ok'

unit-test:
	go vet -v ./...
	go test -v ./...

integration-test: start-local
	@echo 'not implemented'