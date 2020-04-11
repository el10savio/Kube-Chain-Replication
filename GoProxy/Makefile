
run:
	@echo "Running GoProxy"
	go run main.go

build:
	@echo "Building GoProxy"
	cd $$GOPATH/src/k8s.io/klog; git checkout v0.4.0
	go build -v -o bin/goproxy main.go

all: build