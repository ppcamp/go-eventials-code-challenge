#Challenge Makefile

start:
#TODO: commands necessary to start the API
	cd ./api && go build -o ./bin/app ./cmd/web/*go
	cd ./api && ./bin/app

check:
#TODO: include command to test the code and show the results
	cd ./api && go test ./...

setup:
#if needed to setup the enviroment before starting it
# starting a detached docker database
	cd ./docker && docker-compose up -d
# installing the soda cli
	cd ./api && go get github.com/gobuffalo/pop/...
# running soda migrations
	cd ./api && soda migrate
# installing all package deps
	cd ./api && go mod tidy