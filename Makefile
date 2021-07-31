.PHONY : prepare build run

prepare:
	@if [ ! -d "cmd/$(service)" ]; then  echo "ERROR: service '$(service)' undefined"; exit 1; fi
	@ln -sf cmd/$(service)/main.go main_service.go

build: prepare
	go build -o bin

run: build
	./bin

clear:
	rm main_service.go bin backend-service
	
PACKAGES = $(shell go list ./... | grep -v -e . -e mocks | tr '\n' ',')

# unit test & calculate code coverage
test:
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test ./... -cover -coverprofile=coverage.out -covermode=count -coverpkg=$(PACKAGES)
	@go tool cover -func=coverage.out

# make generate swagger from swagger.yml
swagger:
	@if [ -f /docs/swagger/docs.json ]; then rm /docs/swagger/docs.json; fi;
	@echo ">> running generate swagger docs.json from swagger.yml"
	@swagger generate spec -i ./docs/swagger/swagger.yml -o ./docs/swagger/docs.json

# make mock
mock:
	@mockery --all --recursive=true --inpackage --case snake