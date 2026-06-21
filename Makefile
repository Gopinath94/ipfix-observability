.PHONY: build-collector run-collector-build clean-collector run-collector

build-collector:
	go build -o ./bin/ipfix-collector ./cmd/collector/

run-collector-build: build-collector
	./bin/ipfix-collector

clean-collector:
	rm -rf ./bin/ipfix-collector

run-collector:
	go run ./cmd/collector/
