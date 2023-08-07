.PHONY: docker-build
docker-build:
	docker build . -t flightplanner

.PHONY: docker-demo
docker-demo:
	docker run --rm flightplanner

.PHONY: test
test:
	go test -v ./...
