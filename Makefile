.PHONY: assets
assets:
	go-assets-builder assets -o assets.go

.PHONY: clean
clean:
	rm -rf test-app

.PHONY: build
build:
	go build .

.PHONY: run
run:
	./go-boiler -n test-app -p github.com/YusukeKishino -g 1.15.2 --node-version 14.5.0

.PHONY: test-run
test-run: assets clean build run
