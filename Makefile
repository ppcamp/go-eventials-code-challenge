#Challenge Makefile

start:
#TODO: commands necessary to start the API
	cd ./api && go build -o ./bin/app ./cmd/web/*go
	cd ./api && ./bin/app

check:
#TODO: include command to test the code and show the results
	cd ./api && go test ./...

run:
# run another instance of make file with this flag to makes the database update
	cd ./scripts && go build -o ./bin/tool ./populateDB.go
# running the script with update
	cd ./script && ./bin/tool -request ../docker/q2_clientData.csv PUT http://localhost:3000/company/website "name;zip;website"

setup:
#if needed to setup the enviroment before starting it
# extracting the datafile.tgz
	cd ./docker && tar -xf ./dataIntegrationChallenge.tgz
# starting a detached docker database
	cd ./docker && docker-compose up -d
# installing the soda cli
	cd ./api && go get github.com/gobuffalo/pop/...
# running soda migrations that will update the table
	cd ./api && soda migrate
# installing all package deps
	cd ./api && go mod tidy