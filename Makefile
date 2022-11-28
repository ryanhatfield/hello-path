
docker-%: DOCKER_IMAGE_NAME?=ryanhatfield/hello-path
docker-%: DOCKER_IMAGE_TAG?=latest
docker-%: DOCKER_TAG?=$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)
docker-%: DOCKER_FILE?=docker/Dockerfile.server
docker-%: DOCKER_ENV?=DOCKER_DEFAULT_PLATFORM=linux/amd64

docker-build:
	$(DOCKER_ENV) docker build --tag=$(DOCKER_TAG) .

docker-run:
	$(DOCKER_ENV) docker run --rm -p=8081:8081 --interactive --tty $(DOCKER_TAG) /hello-path

docker-push: docker-build
	$(DOCKER_ENV) docker --config ~/.ryanhatfield-docker push $(DOCKER_TAG)

hello-path: server.go
	go build -o=hello-path .

build: hello-path

run:
	go run . /hello-path

run-help:
	go run . --help