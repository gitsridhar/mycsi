NAME=my-csi-driver

.PHONY: all $NAME

REPO_NAME=svenkatdock
IMAGE_VERSION=1.0.0
IMAGE_NAME=$(NAME)

all: $NAME

$NAME:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o  _output/$(NAME) ./src/cmd/main.go
	cp _output/$(NAME) ./src

docker-login:
ifndef $(and DOCKER_USERNAME, DOCKER_PASSWORD)
        $(error DOCKER_USERNAME and DOCKER_PASSWORD must be defined, required for goal (docker-login))
endif
	@docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD) $(DOCKER_SERVER)

build-image: $NAME
	docker build --network=host -t $(REPO_NAME)/$(IMAGE_NAME):$(IMAGE_VERSION) ./src

save-image: build-image
	docker save $(REPO_NAME)/$(IMAGE_NAME):$(IMAGE_VERSION) -o _output/$(IMAGE_NAME)_$(IMAGE_VERSION).tar

push-image: save-image
	docker push $(REPO_NAME)/$(IMAGE_NAME):$(IMAGE_VERSION)
