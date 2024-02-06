run-local-build-docker:
	@echo "Starting the linuxcode/inventory_managerlication with docker."
	docker-compose up --build

lint:
	@echo "linting the project"	
	docker run --rm -v "${PWD}/src:/src"  golangci/golangci-lint:latest /bin/sh -c "cd /src && golangci-lint run"

build-push:
	@echo "Building the image and pushing to repo"

	linuxcode/inventory_manager.IMAGE=example
	linuxcode/inventory_manager.TAG=latest

	docker login ${DOCKER.URL} -u ${DOCKER.USER} -p ${DOCKER.PASS}
	docker build . -t ${DOCKER.URL}/${linuxcode/inventory_manager.IMAGE}:${linuxcode/inventory_manager.TAG} 
	docker push -t ${DOCKER.URL}/${linuxcode/inventory_manager.IMAGE}:${linuxcode/inventory_manager.TAG}
	docker logout

gen-proto:
	docker run -v ./protobuf:/defs namely/protoc-all -f *.proto -l go 
	mv ./protobuf/gen/pb-go/linuxcode/domain/protobuf ./src/protobuf