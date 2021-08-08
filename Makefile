clean:
	rm -rf vendor

vendor:
	go mod vendor

build: clean vendor
	docker image build . -t registry:hello

test:
	docker image build . -t test --target=test

publish:
	docker image push registry:hello

release: build test publish
