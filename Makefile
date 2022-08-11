deps:
	go mod

run-tests:
	ginkgo

setup-tests:
	export PATH=$PATH:$(go env GOPATH)/bin
	ginkgo bootstrap

start:
	go run main.go

verify-code:
	go fmt && go vet
