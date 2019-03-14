NAME		:= itemizer
CN := $(NAME)-container
OUT_DIR 	:= artifacts
OUTPUT 		:= $(OUT_DIR)/go-service
BUILD_DIR 	:= build
export GO111MODULE=on

default: help

artifacts: clean
	go mod download
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o $(OUTPUT) main.go

clean:
	rm -rf $(OUT_DIR)

docker: artifacts
	docker build -t $(NAME) .

### DOCKER
up:
	docker run --name $(CN) -p 8080:8080 -d $(NAME)

inspect:
	docker inspect $(CN)

shell:
	docker exec -it $(CN) bash

logs:
	docker logs -f $(CN)

down:
	docker stop $(CN)

rm: down
	docker rm $(CN)

rmi: rm
	docker rmi $(NAME)

push:
	docker tag $(NAME) henkin/$(NAME)
	docker push henkin/$(NAME)
