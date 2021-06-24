#Challenge Makefile

start:
#TODO: commands necessary to start the API
	cd ./api && go build -o ./bin/app ./cmd/web/*go
	cd ./api && ./bin/app

check:
#TODO: include command to test the code and show the results
	cd ./api && go test ./... -v
	cd ./api && go test ./... -cover profile=coverage.out && go tool cover -html=coverage.out

setup:
#if needed to setup the enviroment before starting it
	cd ./docker && docker-compose up
	cd ./api && soda migrate