DOCKER=$(shell which docker)
GO=$(shell which go)
SWAG=$(shell which swag)


run:
	${GO} run cmd/api/main.go

mysql_create:
	${DOCKER} run --name vuttr --network host -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=main -d mysql

mysql_delete:
	${DOCKER} rm -f vuttr

doc:
	${SWAG} init --dir cmd/api/ --parseDependency