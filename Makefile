.PHONY: build
build:
	go build -o ./_output/dashboard

.PHONY: run
run:
	./_output/dashboard https://www.google.com

.PHONY: docker-build
docker-build:
	docker build . -f Dockerfile.builder -t builder
	docker run --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp builder make build
