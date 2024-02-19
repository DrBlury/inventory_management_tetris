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

gen-sql:
	@echo "Generating the sql"
	npx -p @dbml/cli dbml2sql ./database/schema.dbml -o ./database/schema.sql

gen-sqlc:
	@echo "Generating the sqlc"
	docker run --rm -v $(PWD)/database:/src -w /src sqlc/sqlc generate
	rm -rf ./src/pkg/repo/generated
	mv ./database/repo ./src/pkg/

lint-api:
	@echo "Linting the api-spec"
	docker run --rm -v ${PWD}/api-spec/:/spec redocly/cli lint api.yml

bundle-api:
	@echo "Bundling the api-spec using redocly/redoc"
	docker run --rm -v ${PWD}/api-spec/:/spec redocly/cli bundle api.yml -o bundle.yml

generate-api-echo:
	@echo "Generating the api for echo server using oapi-codegen"
	oapi-codegen -generate types,server,spec -package server ./api-spec/bundle.yml > ./src/pkg/server/api.gen.go

generate-api-chi:
	@echo "Generating the api for chi server using oapi-codegen"
	oapi-codegen --config ./api-spec/server.cfg.yml ./api-spec/bundle.yml > ./src/pkg/server/generated/api.gen.go
	mv ./server.gen.go ./src/pkg/server/generated/api.gen.go
install-local-tools:
	@echo "Installing the dependencies"
	@echo "Installing oapi-codegen"
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
	@echo "Installing sqlc"
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
