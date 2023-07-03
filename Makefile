#DEV
DOCKER_IMAGE_NAME = videochat-system

build-dev:
	docker build -t $(DOCKER_IMAGE_NAME) -f containers/images/Dockerfile . && docker build -t turn -f containers/images/turn/Dockerfile.turn .

clean-dev:
	docker-compose -f containers/composes/dc.dev.yml down -v

run-dev:
	docker-compose -f containers/composes/dc.dev.yml up -d